package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/fatih/structs"
)

func getContentController(resta Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var errResp ErrorResponse

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not get content"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}
		c, err := UnmarshalGetContentReq(b)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not get content"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		var contentSections []map[string]interface{}
		for _, cs := range c.Sections {
			contentSections = append(contentSections, structs.Map(cs))
		}
		sections, err := resta.api.GetContent(c.OctyCustomerID, &contentSections)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not get content"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		dto := struct {
			Status   int                 `json:"status"`
			Message  string              `json:"message"`
			Sections []map[string]string `json:"sections"`
		}{
			Status:   200,
			Message:  "OK",
			Sections: sections,
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&dto)
	}
}
