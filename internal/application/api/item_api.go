package api

// Create octy item with the given parameters.
func (api Application) CreateItem(itemID string, itemCategory string,
	itemName string, itemDescription string, itemPrice int64) error {

	err := api.rest.CreateOctyItem(itemID, itemCategory, itemName, itemDescription, itemPrice)
	if err != nil {
		return err
	}
	return nil
}

// Update octy item with the given parameters.
func (api Application) UpdateItem(itemID string, itemCategory string,
	itemName string, itemDescription string, itemPrice int64, status string) error {

	err := api.rest.UpdateOctyItem(itemID, itemCategory, itemName, itemDescription, itemPrice, status)
	if err != nil {
		return err
	}
	return nil
}
