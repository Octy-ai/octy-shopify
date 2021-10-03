package api

import (
	"github.com/google/uuid"
)

// Create or Update octy profile with the given parameters.
func (api Application) CreateUpdateCustomer(octyCustomerID string, shopifyCustomerID string, generatedNewID bool, hasCharged bool,
	profileData map[string]interface{}, platformInfo map[string]interface{}) (string, string, error) {

	if generatedNewID {
		// new octy-customer-id was generated client side
		// In this case, we want to create a new customer
		octyProfileID, err := create(api, octyCustomerID, shopifyCustomerID, hasCharged, profileData, platformInfo)
		if err != nil {
			return "", "", err
		}

		return octyProfileID, octyCustomerID, nil

	} else {

		// octy-customer-id was found in client local storage.
		// In this case, we need to determine if the profile exists
		// in both app.db and on Octy's servers.

		existingProfile, err := api.rest.GetOctyProfile(octyCustomerID)
		if err != nil {

			if err.Error() == "no profiles" {
				// is NOT exisitng profile with customer_id : octy-customer-id
				// In this case, we need to determine if the profile has been
				// deleted or merged.

				// identify profile
				profileIden, err := api.rest.IdentifyOctyProfile(octyCustomerID)
				if err != nil {
					return "", "", err
				}

				if profileIden["ParentProfileID"] != nil {
					// is parent profile
					// This tells us that the profile has been merged. In this case, we
					// need to update the customer in app.db and pass the new customerID to the
					// client so the value for octy-customer-id can be updated in local storage.

					parentProfileID := profileIden["ParentProfileID"].(string)
					parentCustomerID := profileIden["ParentCustomerID"].(string)
					authenticatedIDValue := profileIden["AuthenticatedID"].(string)

					if parentCustomerID != octyCustomerID {
						// if the parent customer ID does not match provided octy-customer-id, then
						// delete existing customer from app.db with provided octy-customer-id.
						err := api.db.DeleteCustomer(octyCustomerID)
						if err != nil {
							return "", "", err
						}
					}

					profileID, customerID, err := update(api, parentCustomerID, parentProfileID, authenticatedIDValue,
						hasCharged, profileData, platformInfo)
					if err != nil {
						return "", "", err
					}
					return profileID, customerID, nil

				} else {
					// is NOT parent profile
					// This tells us that the profile has been deleted. In this case, we
					// need to delete the customer in app.db, create a new customer pass the new cty-customer-id to the
					// client so the value for octy-customer-id can be updated in local storage.
					id := uuid.New()
					newOctyCustomerID := "octy-customer-id-" + id.String()
					err := api.db.DeleteCustomer(octyCustomerID)
					if err != nil {
						return "", "", err
					}
					octyProfileID, err := create(api, newOctyCustomerID, shopifyCustomerID, hasCharged, profileData, platformInfo)
					if err != nil {
						return "", "", err
					}

					return octyProfileID, newOctyCustomerID, nil

				}

			} else {
				// error occurred
				return "", "", err
			}

		} else {
			// is exisitng profile with customer_id : octy-customer-id
			// In this case, we simply update the customer in app.db with
			// the provided parameters.
			exProfileID := existingProfile["ProfileID"].(string)
			profileID, customerID, err := update(api, octyCustomerID, exProfileID, shopifyCustomerID,
				hasCharged, profileData, platformInfo)
			if err != nil {
				return "", "", err
			}
			return profileID, customerID, nil
		}

	}

}

// <Private Functions>

func create(api Application, octyCustomerID string, shopifyCustomerID string, hasCharged bool,
	profileData map[string]interface{}, platformInfo map[string]interface{}) (string, error) {

	octyProfileID, err := api.rest.CreateOctyProfile(
		octyCustomerID,
		hasCharged,
		profileData,
		platformInfo,
	)
	if err != nil {
		return "", err
	}

	err = api.db.CreateCustomer(octyCustomerID, octyProfileID, shopifyCustomerID)
	if err != nil {
		return "", err
	}
	return octyProfileID, nil
}

func update(api Application, customerID string, octyprofileID string, shopifyCustomerID string, hasCharged bool,
	profileData map[string]interface{}, platformInfo map[string]interface{}) (string, string, error) {

	if shopifyCustomerID != "" {
		// customer is authenticated
		// API call update existing profile>profile_data>shopify_customer_id param
		status := "active"
		_, err := api.rest.UpdateOctyProfile(
			octyprofileID,
			customerID,
			hasCharged,
			profileData,
			platformInfo,
			status,
		)
		if err != nil {
			return "", "", err
		}
	}

	// create or update in DB (with provided shopifyCustomerID)
	_, err := api.db.GetCustomer(customerID)
	if err != nil {
		if err.Error() == "empty result" {
			// If no customers, create new customer
			err = api.db.CreateCustomer(customerID, octyprofileID, shopifyCustomerID)
			if err != nil {
				return "", "", err
			}
		} else {
			// actual DB error occurred
			return "", "", err
		}
	} else {
		// customer already exists, perform update
		err = api.db.UpdateCustomer(customerID, octyprofileID, shopifyCustomerID)
		if err != nil {
			return "", "", err
		}
	}

	return octyprofileID, customerID, nil
}
