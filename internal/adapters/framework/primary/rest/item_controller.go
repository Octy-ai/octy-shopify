package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Octy-ai/octy-shopify/pkg/utils"
)

func createItemController(resta Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var errResp ErrorResponse

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not create item"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		i, err := UnmarshalCreateShopifyProductsWHReq(b)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not create item"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		// NOTE: each vartiant is considred a seperate item
		for _, variant := range i.Variants {
			// convert string representation of price to lowest denomination numerical value
			itemPrice := formatVariantPrice(variant.Price)
			var desc string
			descLen := len(i.BodyHTML.(string))
			switch {
			case descLen < 1:
				desc = variant.Title // set the description to the variant title
			case descLen > 39:
				desc = utils.TruncateString(i.BodyHTML.(string), 40) // trim provided description to 40 characters
			default:
				desc = i.BodyHTML.(string)
			}
			name := i.Title + ">" + variant.Title
			// concatinate the productID and variantID to be used as the Octy itemID
			variantProductID := strconv.Itoa(int(i.ID)) + "-" + strconv.Itoa(int(variant.ID))
			err = resta.api.CreateItem(variantProductID, i.ProductType, name, desc, itemPrice)
			if err != nil {
				log.Println(err)
				errResp.StatusCode = 400
				errResp.Message = "Could not create item"
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

func updateItemController(resta Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var errResp ErrorResponse

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not update item"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}
		i, err := UnmarshalCreateShopifyProductsWHReq(b)
		if err != nil {
			log.Println(err)
			errResp.StatusCode = 400
			errResp.Message = "Could not update item"
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		// NOTE: each vartiant is considred a seperate item
		for _, variant := range i.Variants {
			// convert string representation of price to lowest denomination numerical value
			itemPrice := formatVariantPrice(variant.Price)

			// convert status to allowed octy status
			var status string = "active"
			if i.Status != "active" {
				status = "inactive"
			}

			var desc string
			descLen := len(i.BodyHTML.(string))
			switch {
			case descLen < 1:
				desc = variant.Title
			case descLen > 39:
				desc = utils.TruncateString(i.BodyHTML.(string), 40) // trim description to 40 characters
			default:
				desc = i.BodyHTML.(string)
			}
			name := i.Title + ">" + variant.Title
			// concatinate the productID and variantID to be used as the Octy itemID
			variantProductID := strconv.Itoa(int(i.ID)) + "-" + strconv.Itoa(int(variant.ID))
			err = resta.api.UpdateItem(variantProductID, i.ProductType, name, desc, itemPrice, status)
			if err != nil {
				log.Println(err)
				errResp.StatusCode = 400
				errResp.Message = "Could not update item"
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(errResp)
				return
			}
		}

		dto := struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{
			Status:  200,
			Message: "OK",
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&dto)
	}
}

// < Private Functions >
func formatVariantPrice(price string) int64 {
	t := strings.Replace(price, ".", "", -1)
	itemPrice, _ := strconv.ParseInt(t, 10, 64)
	return itemPrice
}
