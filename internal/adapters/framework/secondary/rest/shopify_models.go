package rest

import "encoding/json"

// ** Shopify REST Request Models **

// ---

type GetShopifyCustomerResp struct {
	Customer Customer `json:"customer"`
}

type Customer struct {
	ID                        int64         `json:"id"`
	Email                     string        `json:"email"`
	AcceptsMarketing          bool          `json:"accepts_marketing"`
	CreatedAt                 string        `json:"created_at"`
	UpdatedAt                 string        `json:"updated_at"`
	FirstName                 string        `json:"first_name"`
	LastName                  string        `json:"last_name"`
	OrdersCount               int64         `json:"orders_count"`
	State                     string        `json:"state"`
	TotalSpent                string        `json:"total_spent"`
	LastOrderID               int64         `json:"last_order_id"`
	Note                      interface{}   `json:"note"`
	VerifiedEmail             bool          `json:"verified_email"`
	MultipassIdentifier       interface{}   `json:"multipass_identifier"`
	TaxExempt                 bool          `json:"tax_exempt"`
	Phone                     string        `json:"phone"`
	Tags                      string        `json:"tags"`
	LastOrderName             string        `json:"last_order_name"`
	Currency                  string        `json:"currency"`
	Addresses                 []Address     `json:"addresses"`
	AcceptsMarketingUpdatedAt string        `json:"accepts_marketing_updated_at"`
	MarketingOptInLevel       interface{}   `json:"marketing_opt_in_level"`
	TaxExemptions             []interface{} `json:"tax_exemptions"`
	AdminGraphqlAPIID         string        `json:"admin_graphql_api_id"`
	DefaultAddress            Address       `json:"default_address"`
}

type Address struct {
	ID           int64       `json:"id"`
	CustomerID   int64       `json:"customer_id"`
	FirstName    interface{} `json:"first_name"`
	LastName     interface{} `json:"last_name"`
	Company      interface{} `json:"company"`
	Address1     string      `json:"address1"`
	Address2     string      `json:"address2"`
	City         string      `json:"city"`
	Province     string      `json:"province"`
	Country      string      `json:"country"`
	Zip          string      `json:"zip"`
	Phone        string      `json:"phone"`
	Name         string      `json:"name"`
	ProvinceCode string      `json:"province_code"`
	CountryCode  string      `json:"country_code"`
	CountryName  string      `json:"country_name"`
	Default      bool        `json:"default"`
}

func UnmarshalGetShopifyCustomerResp(data []byte) (GetShopifyCustomerResp, error) {
	var r GetShopifyCustomerResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type GetShopifyProductsResp struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID                int64       `json:"id"`
	Title             string      `json:"title"`
	BodyHTML          string      `json:"body_html"`
	Vendor            string      `json:"vendor"`
	ProductType       string      `json:"product_type"`
	CreatedAt         string      `json:"created_at"`
	Handle            string      `json:"handle"`
	UpdatedAt         string      `json:"updated_at"`
	PublishedAt       string      `json:"published_at"`
	TemplateSuffix    interface{} `json:"template_suffix"`
	PublishedScope    string      `json:"published_scope"`
	Tags              string      `json:"tags"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
	Variants          []Variant   `json:"variants"`
	Options           []Option    `json:"options"`
	Images            []Image     `json:"images"`
	Image             *Image      `json:"image"`
}

type Image struct {
	ID                int64       `json:"id"`
	ProductID         int64       `json:"product_id"`
	Position          int64       `json:"position"`
	CreatedAt         string      `json:"created_at"`
	UpdatedAt         string      `json:"updated_at"`
	Alt               interface{} `json:"alt"`
	Width             int64       `json:"width"`
	Height            int64       `json:"height"`
	Src               string      `json:"src"`
	VariantIDS        []int64     `json:"variant_ids"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
}

type Option struct {
	ID        int64    `json:"id"`
	ProductID int64    `json:"product_id"`
	Name      string   `json:"name"`
	Position  int64    `json:"position"`
	Values    []string `json:"values"`
}

type Variant struct {
	ID                   int64              `json:"id"`
	ProductID            int64              `json:"product_id"`
	Title                string             `json:"title"`
	Price                string             `json:"price"`
	Sku                  string             `json:"sku"`
	Position             int64              `json:"position"`
	InventoryPolicy      string             `json:"inventory_policy"`
	CompareAtPrice       interface{}        `json:"compare_at_price"`
	FulfillmentService   string             `json:"fulfillment_service"`
	InventoryManagement  string             `json:"inventory_management"`
	Option1              string             `json:"option1"`
	Option2              interface{}        `json:"option2"`
	Option3              interface{}        `json:"option3"`
	CreatedAt            string             `json:"created_at"`
	UpdatedAt            string             `json:"updated_at"`
	Taxable              bool               `json:"taxable"`
	Barcode              string             `json:"barcode"`
	Grams                int64              `json:"grams"`
	ImageID              *int64             `json:"image_id"`
	Weight               float64            `json:"weight"`
	WeightUnit           string             `json:"weight_unit"`
	InventoryItemID      int64              `json:"inventory_item_id"`
	InventoryQuantity    int64              `json:"inventory_quantity"`
	OldInventoryQuantity int64              `json:"old_inventory_quantity"`
	PresentmentPrices    []PresentmentPrice `json:"presentment_prices"`
	RequiresShipping     bool               `json:"requires_shipping"`
	AdminGraphqlAPIID    string             `json:"admin_graphql_api_id"`
}

type PresentmentPrice struct {
	Price          Price       `json:"price"`
	CompareAtPrice interface{} `json:"compare_at_price"`
}

type Price struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}

func UnmarshalGetShopifyProductsResp(data []byte) (GetShopifyProductsResp, error) {
	var r GetShopifyProductsResp
	err := json.Unmarshal(data, &r)
	return r, err
}
