package api

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Octy-ai/octy-shopify/pkg/utils"
)

// Returns product recommendations for a given customer identifier.
func (api Application) GetRecommendations(customerID string) ([]map[string]interface{}, error) {

	customer, err := api.db.GetCustomer(customerID)
	if err != nil {
		return nil, err
	}

	if customer["octyProfileID"] == "" {
		return nil, errors.New("no profile ID found for customer")
	}
	profileID := customer["octyProfileID"]

	recs, err := api.rest.GetRecommendations(profileID)
	if err != nil {
		return nil, err
	}

	if len(recs) == 0 {
		return []map[string]interface{}{}, nil
	}

	var recItemIDs string
	var ids []string
	for _, r := range recs {
		s := strings.Split(r["itemID"].(string), "-")
		if !utils.Contains(ids, s[0]) {
			recItemIDs = recItemIDs + "," + s[0]
			ids = append(ids, s[0])
		}

	}

	items, err := api.rest.GetShopifyProducts(recItemIDs)
	if err != nil {
		return nil, err
	}

	// map shopify items to rec items and create response object
	var itemsResp []map[string]interface{}
	for _, r := range recs {
		s := strings.Split(r["itemID"].(string), "-")
		for _, item := range items {
			if s[0] == strconv.Itoa(int(item["ID"].(int64))) {
				itemsResp = append(itemsResp, map[string]interface{}{
					"item_name":        item["Title"].(string),
					"item_category":    item["ProductType"].(string),
					"item_description": item["BodyHTML"].(string),
					"item_price":       item["Variants"].([]interface{})[0].(map[string]interface{})["Price"].(string),
					"item_image_url":   item["Image"].(map[string]interface{})["Src"],
					"item_link":        api.config.Shopify.StoreRootURL + "products/" + item["Handle"].(string),
					"id":               strconv.Itoa(int(item["ID"].(int64))),
					"octy_rec_score":   r["score"].(float64),
				})
			}
		}
	}

	return itemsResp, nil
}
