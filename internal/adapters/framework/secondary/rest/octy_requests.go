package rest

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/fatih/structs"
)

// ** Profiles **

func (ha Adapter) CreateOctyProfile(customerID string, hasCharged bool,
	profileData map[string]interface{}, platformInfo map[string]interface{}) (string, error) {

	profileReq := CreateOctyProfileReq{
		Profiles: []ProfileReq{
			{
				CustomerID:   customerID,
				HasCharged:   hasCharged,
				ProfileData:  profileData,
				PlatformInfo: platformInfo,
			},
		},
	}

	requestBody, err := profileReq.Marshal()
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", ha.config.OctyAPIURIs.CreateProfile, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(
		ha.config.OctyCreds.PublicKey+":"+ha.config.OctyCreds.SecretKey)))
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Println(string(body))

	if resp.StatusCode > 202 {
		return "", errors.New("Invalid HTTP status code: " + strconv.Itoa(resp.StatusCode))
	}

	profileResp, err := UnmarshalCreateOctyProfileResp(body)
	if err != nil {
		return "", err
	}
	return profileResp.Profiles[0].ProfileID, nil
}

func (ha Adapter) UpdateOctyProfile(profileID string, customerID string, hasCharged bool,
	profileData map[string]interface{}, platformInfo map[string]interface{}, status string) (string, error) {

	profileReq := UpdateOctyProfileReq{
		Profiles: []UProfileReq{
			{
				ProfileID:    profileID,
				CustomerID:   customerID,
				HasCharged:   hasCharged,
				ProfileData:  profileData,
				PlatformInfo: platformInfo,
				Status:       status,
			},
		},
	}

	requestBody, err := profileReq.Marshal()
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", ha.config.OctyAPIURIs.UpdateProfile, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(
		ha.config.OctyCreds.PublicKey+":"+ha.config.OctyCreds.SecretKey)))
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 202 {
		return "", errors.New("Invalid HTTP status code: " + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	profileResp, err := UnmarshalUpdateOctyProfileResp(body)
	if err != nil {
		return "", err
	}
	return profileResp.Profiles[0].ProfileID, nil
}

func (ha Adapter) GetOctyProfile(customerID string) (map[string]interface{}, error) {

	if customerID == "" {
		return nil, errors.New("no customerID supplied")
	}
	url := ha.config.OctyAPIURIs.GetProfiles + "?id=" + customerID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(
		ha.config.OctyCreds.PublicKey+":"+ha.config.OctyCreds.SecretKey)))
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 202 && resp.StatusCode <= 400 {
		return nil, errors.New("no profiles")
	} else if resp.StatusCode > 400 {
		return nil, errors.New("server error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	profileResp, err := UnmarshalGetOctyProfileResp(body)
	if err != nil {
		return nil, err
	}

	return structs.Map(profileResp.Profiles[0]), nil
}

func (ha Adapter) IdentifyOctyProfile(customerID string) (map[string]interface{}, error) {
	identifyProfileReq := IdentifyOctyProfileReq{
		Identifiers: []string{
			customerID,
		},
	}

	requestBody, err := identifyProfileReq.Marshal()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", ha.config.OctyAPIURIs.IdentifyProfiles, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(
		ha.config.OctyCreds.PublicKey+":"+ha.config.OctyCreds.SecretKey)))
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 200 {
		return nil, errors.New("Invalid HTTP status code: " + strconv.Itoa(resp.StatusCode))
	}

	identifyProfileResp, err := UnmarshalIdentifyOctyProfileResp(body)
	if err != nil {
		return nil, err
	}

	profileIden := map[string]interface{}{
		"Identifier":           identifyProfileResp.Identifiers[0].Identifier,
		"ParentProfileID":      *identifyProfileResp.Identifiers[0].ParentProfileID,
		"ParentCustomerID":     *identifyProfileResp.Identifiers[0].ParentCustomerID,
		"AuthenticatedIDValue": *identifyProfileResp.Identifiers[0].AuthenticatedIDValue,
	}

	return profileIden, nil
}

// ** Events **

func (ha Adapter) CreateOctyEvent(eventType string,
	eventProperties map[string]interface{}, profileID string) (string, error) {

	eventReq := CreateOctyEventReq{
		EventType:       eventType,
		EventProperties: eventProperties,
		ProfileID:       profileID,
	}

	requestBody, err := eventReq.Marshal()
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", ha.config.OctyAPIURIs.CreateEvent, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(
		ha.config.OctyCreds.PublicKey+":"+ha.config.OctyCreds.SecretKey)))
	req.Header.Add("Content-Type", "application/json")

	log.Printf("Making request to %v", ha.config.OctyAPIURIs.CreateEvent)
	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 202 {
		return "", errors.New("Invalid HTTP status code: " + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	eventResp, err := UnmarshalCreateOctyEventResp(body)
	if err != nil {
		return "", err
	}
	return eventResp.EventID, nil
}

// ** Items **

func (ha Adapter) CreateOctyItem(itemID string, itemCategory string, itemName string,
	itemDescription string, itemPrice int64) error {

	itemReq := CreateOctyItemsReq{
		Items: []CItemReq{
			{
				ItemID:          itemID,
				ItemCategory:    itemCategory,
				ItemName:        itemName,
				ItemDescription: itemDescription,
				ItemPrice:       itemPrice,
			},
		},
	}

	requestBody, err := itemReq.Marshal()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", ha.config.OctyAPIURIs.CreateItem, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(
		ha.config.OctyCreds.PublicKey+":"+ha.config.OctyCreds.SecretKey)))
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 202 {
		return errors.New("Invalid HTTP status code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}

func (ha Adapter) UpdateOctyItem(itemID string, itemCategory string, itemName string,
	itemDescription string, itemPrice int64, status string) error {

	itemReq := UpdateOctyItemsReq{
		Items: []UItemReq{
			{
				ItemID:          itemID,
				ItemCategory:    itemCategory,
				ItemName:        itemName,
				ItemDescription: itemDescription,
				ItemPrice:       itemPrice,
				Status:          status,
			},
		},
	}

	requestBody, err := itemReq.Marshal()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", ha.config.OctyAPIURIs.UpdateItem, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(
		ha.config.OctyCreds.PublicKey+":"+ha.config.OctyCreds.SecretKey)))
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 202 {
		return errors.New("Invalid HTTP status code: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}

// ** Content **

func (ha Adapter) GetTemplates() ([]map[string]interface{}, error) {

	req, err := http.NewRequest("GET", ha.config.OctyAPIURIs.GetTemplates, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(
		ha.config.OctyCreds.PublicKey+":"+ha.config.OctyCreds.SecretKey)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cursor", "0")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 202 && resp.StatusCode <= 400 {
		return nil, errors.New("failed to generate message content")
	} else if resp.StatusCode > 202 && resp.StatusCode > 400 {
		return nil, errors.New("Invalid HTTP status code: " + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	getTemplatesResp, err := UnmarshalGetOctyTemplatesResp(body)
	if err != nil {
		return nil, err
	}

	var templates []map[string]interface{}
	for _, t := range getTemplatesResp.Templates {
		templates = append(templates, structs.Map(t))
	}

	return templates, nil

}

func (ha Adapter) GenerateContent(messages []map[string]string) ([]map[string]interface{}, []map[string]interface{}, []map[string]interface{}, error) {

	reqMessages := []Message{}
	for _, m := range messages {
		reqMessages = append(reqMessages, Message{
			TemplateID:         m["TemplateID"],
			ItemRecommendation: false,
			Data:               []map[string]interface{}{{}}, // pass empty object. We are no populating placeholder values.
		})
	}

	genContentReq := GenOctyContentReq{
		Messages: reqMessages,
	}

	requestBody, err := genContentReq.Marshal()
	if err != nil {
		return nil, nil, nil, err
	}

	req, err := http.NewRequest("POST", ha.config.OctyAPIURIs.GenerateContent,
		bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, nil, nil, err
	}

	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(
		ha.config.OctyCreds.PublicKey+":"+ha.config.OctyCreds.SecretKey)))
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, err
	}

	if resp.StatusCode > 202 && resp.StatusCode <= 400 {
		return nil, nil, nil, errors.New("failed to generate message content")
	} else if resp.StatusCode > 202 && resp.StatusCode > 400 {
		return nil, nil, nil, errors.New("Invalid HTTP status code: " + strconv.Itoa(resp.StatusCode))
	}

	genContentResp, err := UnmarshalGenOctyContentResp(body)
	if err != nil {
		return nil, nil, nil, err
	}

	var generatedMessages []map[string]interface{}
	for _, m := range genContentResp.GeneratedMessages {
		generatedMessages = append(generatedMessages, structs.Map(m))
	}
	var failedMessages []map[string]interface{}
	for _, fm := range genContentResp.FailedMessages {
		//failedMessages = append(failedMessages, fm.(map[string]interface{}))
		failedMessages = append(failedMessages, structs.Map(fm))
	}
	var failedTemplates []map[string]interface{}
	for _, ft := range genContentResp.FailedTemplates {
		//failedTemplates = append(failedTemplates, ft.(map[string]interface{}))
		failedTemplates = append(failedTemplates, structs.Map(ft))
	}

	return generatedMessages, failedMessages, failedTemplates, nil
}

// ** Recommendation **

func (ha Adapter) GetRecommendations(profileID string) ([]map[string]interface{}, error) {

	getRecReq := GetOctyRecReq{
		ProfileIDS: []string{profileID},
	}

	requestBody, err := getRecReq.Marshal()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", ha.config.OctyAPIURIs.PredictRecommendations,
		bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(
		ha.config.OctyCreds.PublicKey+":"+ha.config.OctyCreds.SecretKey)))
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 202 && resp.StatusCode <= 400 {
		return nil, errors.New("unable to get item recommendations for provided profiles")
	} else if resp.StatusCode > 202 && resp.StatusCode > 400 {
		return nil, errors.New("Invalid HTTP status code: " + strconv.Itoa(resp.StatusCode))
	}

	getRecResp, err := UnmarshalGetOctyRecResp(body)
	if err != nil {
		return nil, err
	}

	var items []map[string]interface{}
	for _, r := range getRecResp.Recommendations {
		for _, i := range r.Recommendations {
			items = append(items, map[string]interface{}{
				"itemID": i.ItemID,
				"score":  i.Score,
			})
		}
	}
	return items, nil
}
