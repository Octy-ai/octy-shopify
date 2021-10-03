package ports

// Driven ports

// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {

	// Customers

	CreateUpdateCustomer(octyCustomerID string, shopifyCustomerID string, generatedNewID bool, hasCharged bool,
		profileData map[string]interface{}, platformInfo map[string]interface{}) (string, string, error)

	// Events

	CreateEvent(eventType string, eventProperties map[string]interface{},
		customerID string) (string, string, error)

	// Items

	CreateItem(itemID string, itemCategory string,
		itemName string, itemDescription string, itemPrice int64) error

	UpdateItem(itemID string, itemCategory string,
		itemName string, itemDescription string, itemPrice int64, status string) error

	// Content

	GetContent(customerID string, sections *[]map[string]interface{}) ([]map[string]string, error)

	// Recommendations

	GetRecommendations(customerID string) ([]map[string]interface{}, error)
}
