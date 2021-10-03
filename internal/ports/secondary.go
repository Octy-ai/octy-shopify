package ports

// Driving ports

type RestPort interface {

	// ** OCTY **

	// Profiles

	CreateOctyProfile(customerID string, hasCharged bool,
		profileData map[string]interface{}, platformInfo map[string]interface{}) (string, error)

	UpdateOctyProfile(profileID string, customerID string, hasCharged bool,
		profileData map[string]interface{}, platformInfo map[string]interface{}, status string) (string, error)

	GetOctyProfile(customerID string) (map[string]interface{}, error)

	IdentifyOctyProfile(customerID string) (map[string]interface{}, error)

	// Events

	CreateOctyEvent(eventType string, eventProperties map[string]interface{}, profileID string) (string, error)

	// Items

	CreateOctyItem(itemID string, itemCategory string, itemName string,
		itemDescription string, itemPrice int64) error

	UpdateOctyItem(itemID string, itemCategory string, itemName string,
		itemDescription string, itemPrice int64, status string) error

	// Content

	GetTemplates() ([]map[string]interface{}, error)

	GenerateContent(messages []map[string]string) ([]map[string]interface{}, []map[string]interface{}, []map[string]interface{}, error)

	// Recommendations

	GetRecommendations(profileID string) ([]map[string]interface{}, error)

	// ** SHOPIFY **

	// Customers

	GetShopifyCustomer(customerID string) (string, error)

	GetShopifyProducts(itemIDs string) ([]map[string]interface{}, error)
}

type DatabasePort interface {
	Connect()

	CreateCustomer(octyCustomerID string, octyProfileID string, shopifyCustomerID string) error

	GetCustomer(octyCustomerID string) (map[string]string, error)

	UpdateCustomer(octyCustomerID string, octyProfileID string, shopifyCustomerID string) error

	DeleteCustomer(octyCustomerID string) error
}
