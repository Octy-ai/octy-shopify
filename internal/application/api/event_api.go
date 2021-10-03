package api

import "errors"

// Create octy event instance with given parameters.
func (api Application) CreateEvent(eventType string, eventProperties map[string]interface{},
	customerID string) (string, string, error) {

	customer, err := api.db.GetCustomer(customerID)
	if err != nil {
		return "", "", err
	}

	if customer["octyProfileID"] == "" {
		return "", "", errors.New("no profile ID found for customer")
	}
	profileID := customer["octyProfileID"]

	eID, err := api.rest.CreateOctyEvent(eventType, eventProperties, profileID)
	if err != nil {
		return "", "", err
	}
	return eID, profileID, nil
}
