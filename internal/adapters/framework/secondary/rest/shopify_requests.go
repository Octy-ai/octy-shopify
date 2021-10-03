package rest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/fatih/structs"
)

// ** Customers **

func (ha Adapter) GetShopifyCustomer(customerID string) (string, error) {

	req, err := http.NewRequest("GET", ha.config.Shopify.GetCustomerURI+customerID+".json", nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("X-Shopify-Access-Token",
		ha.config.Shopify.APISecret)
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
	customerResp, err := UnmarshalGetShopifyCustomerResp(body)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(customerResp.Customer.ID)), nil
}

// ** Products **
func (ha Adapter) GetShopifyProducts(itemIDs string) ([]map[string]interface{}, error) {

	req, err := http.NewRequest("GET", ha.config.Shopify.GetProductsURI+"?ids="+itemIDs, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Shopify-Access-Token",
		ha.config.Shopify.APISecret)
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 202 {
		return nil, errors.New("Invalid HTTP status code: " + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	productResp, err := UnmarshalGetShopifyProductsResp(body)
	if err != nil {
		return nil, err
	}

	var shopifyProducts []map[string]interface{}
	for _, m := range productResp.Products {
		shopifyProducts = append(shopifyProducts, structs.Map(m))
	}
	return shopifyProducts, nil
}
