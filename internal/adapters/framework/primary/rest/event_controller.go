package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func createEventController(resta Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var errResp ErrorResponse

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not create event"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		e, err := UnmarshalCreateEventReq(b)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not create event"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}
		eventID, profileID, err := resta.api.CreateEvent(e.EventType, e.EventProperties, e.OctyCustomerID)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not create event"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		dto := struct {
			Status    int    `json:"status"`
			Message   string `json:"message"`
			EventID   string `json:"event_id"`
			ProfileID string `json:"profile_id"`
		}{
			Status:    201,
			Message:   "OK",
			ProfileID: profileID,
			EventID:   eventID,
		}
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(&dto)
	}
}

func createChargedEventSHWHController(resta Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var errResp ErrorResponse

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not create event"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		e, err := UnmarshalShopifyOrderPaymentWHReq(b)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not create event"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}
		shopifyCustomerID := strconv.Itoa(int(e.Customer.ID))
		generatedNewID := false
		hasCharged := true
		profileData := map[string]interface{}{
			"shopify_customer_id":     shopifyCustomerID,       // required
			"city":                    e.BillingAddress.City,   // optional
			"buyer_accepts_marketing": e.BuyerAcceptsMarketing, // optional
		} // NOTE: add other profile data attributes
		platformInfo := map[string]interface{}{} // NOTE: add platform info attributes

		// create update customer
		_, octyCustomerID, err := resta.api.CreateUpdateCustomer(
			shopifyCustomerID,
			shopifyCustomerID,
			generatedNewID,
			hasCharged,
			profileData,
			platformInfo)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			w.WriteHeader(400)
			errResp.Message = err.Error()
			json.NewEncoder(w).Encode(errResp)
			return
		}

		for _, item := range e.LineItems {

			eventType := "charged"
			var paymentMethod string
			if e.PaymentGatewayNames[0] == "" {
				paymentMethod = "cash"
			} else {
				paymentMethod = e.PaymentGatewayNames[0]
			}
			eventProperties := map[string]interface{}{
				"payment_method": paymentMethod,
				"item_id":        strconv.Itoa(int(item.VariantID)),
			}

			_, _, err := resta.api.CreateEvent(eventType, eventProperties, octyCustomerID)
			if err != nil {
				log.Println(err)
				errResp.StatusCode = 400
				errResp.Message = "Could not create event"
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(errResp)
				return
			}
		}

		dto := struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{
			Status:  201,
			Message: "OK",
		}
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(&dto)

	}
}
