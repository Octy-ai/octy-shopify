package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func createCustomerController(resta Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var errResp ErrorResponse

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			w.WriteHeader(400)
			errResp.Message = "Could not create customer"
			json.NewEncoder(w).Encode(errResp)
			return
		}
		c, err := UnmarshalCreateCustomerReq(b)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			w.WriteHeader(400)
			errResp.Message = "Could not create customer"
			json.NewEncoder(w).Encode(errResp)
			return
		}
		octyProfileID, octyCustomerID, err := resta.api.CreateUpdateCustomer(
			c.OctyCustomerID,
			c.ShopifyCustomerID,
			c.GeneratedNewID,
			c.HasCharged,
			c.ProfileData,
			c.PlatformInfo)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			w.WriteHeader(400)
			errResp.Message = err.Error()
			json.NewEncoder(w).Encode(errResp)
			return
		}

		dto := struct {
			Status     int    `json:"status"`
			Message    string `json:"message"`
			ProfileID  string `json:"profile_id"`
			CustomerID string `json:"customer_id"`
		}{
			Status:     201,
			Message:    "OK",
			ProfileID:  octyProfileID,
			CustomerID: octyCustomerID,
		}
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(&dto)
	}
}
