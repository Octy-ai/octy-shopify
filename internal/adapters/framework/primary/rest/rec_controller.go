package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getRecController(resta Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var errResp ErrorResponse

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not get recommendations"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}
		rec, err := UnmarshalGetRecReq(b)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not get recommendations"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}
		recommendations, err := resta.api.GetRecommendations(rec.OctyCustomerID)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not get recommendations"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		dto := struct {
			Status  int                      `json:"status"`
			Message string                   `json:"message"`
			Items   []map[string]interface{} `json:"items"`
		}{
			Status:  200,
			Message: "OK",
			Items:   recommendations,
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&dto)
	}
}
