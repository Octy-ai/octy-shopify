package rest

import "encoding/json"

// ** Octy-Shopify app REST Request Models **

// ---

type CustomerReq struct {
	OctyCustomerID    string                 `json:"octy_customer_id"`
	ShopifyCustomerID string                 `json:"shopify_customer_id"`
	GeneratedNewID    bool                   `json:"generated_new_id"`
	HasCharged        bool                   `json:"has_charged"`
	ProfileData       map[string]interface{} `json:"profile_data"`
	PlatformInfo      map[string]interface{} `json:"platform_info"`
}

func UnmarshalCreateCustomerReq(data []byte) (CustomerReq, error) {
	var r CustomerReq
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type CreateEventReq struct {
	EventType       string                 `json:"event_type"`
	EventProperties map[string]interface{} `json:"event_properties"`
	OctyCustomerID  string                 `json:"octy_customer_id"`
}

func UnmarshalCreateEventReq(data []byte) (CreateEventReq, error) {
	var r CreateEventReq
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type CreateItemReq struct {
	ItemID          string `json:"item_id"`
	ItemCategory    string `json:"item_category"`
	ItemName        string `json:"item_name"`
	ItemDescription string `json:"item_description"`
	ItemPrice       int64  `json:"item_price"`
}

func UnmarshalCreateItemReq(data []byte) (CreateItemReq, error) {
	var r CreateItemReq
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type UpdateItemReq struct {
	ItemID          string `json:"item_id"`
	ItemCategory    string `json:"item_category"`
	ItemName        string `json:"item_name"`
	ItemDescription string `json:"item_description"`
	ItemPrice       int64  `json:"item_price"`
	Status          string `json:"status"`
}

func UnmarshalUpdateItemReq(data []byte) (UpdateItemReq, error) {
	var r UpdateItemReq
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type DeleteItemsReq struct {
	Items []string `json:"items"`
}

func UnmarshalDeleteItemReq(data []byte) (DeleteItemsReq, error) {
	var r DeleteItemsReq
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type ContentReq struct {
	OctyCustomerID string    `json:"octy_customer_id"`
	Sections       []Section `json:"sections"`
}

type Section struct {
	SectionID    string `json:"section_id"`
	DefaultValue string `json:"default_value"`
}

func UnmarshalGetContentReq(data []byte) (ContentReq, error) {
	var r ContentReq
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type RecReq struct {
	OctyCustomerID string `json:"octy_customer_id"`
}

func UnmarshalGetRecReq(data []byte) (RecReq, error) {
	var r RecReq
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

// ** Octy-Shopify app REST Response Models **

// ---
type ErrorResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func (r *ErrorResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

// ** Shopify Webhook Request Models **

// ---

type CreateShopifyProductsWHReq struct {
	ID                float64     `json:"id"`
	Title             string      `json:"title"`
	BodyHTML          interface{} `json:"body_html"`
	Vendor            string      `json:"vendor"`
	ProductType       string      `json:"product_type"`
	CreatedAt         interface{} `json:"created_at"`
	Handle            string      `json:"handle"`
	UpdatedAt         string      `json:"updated_at"`
	PublishedAt       string      `json:"published_at"`
	TemplateSuffix    interface{} `json:"template_suffix"`
	Status            string      `json:"status"`
	PublishedScope    string      `json:"published_scope"`
	Tags              string      `json:"tags"`
	AdminGraphqlAPIID string      `json:"admin_graphql_api_id"`
	Variants          []Variant   `json:"variants"`
	Options           []Option    `json:"options"`
	Images            []Image     `json:"images"`
	Image             interface{} `json:"image"`
}

type Image struct {
	ID                float64       `json:"id"`
	ProductID         float64       `json:"product_id"`
	Position          int64         `json:"position"`
	CreatedAt         interface{}   `json:"created_at"`
	UpdatedAt         interface{}   `json:"updated_at"`
	Alt               interface{}   `json:"alt"`
	Width             int64         `json:"width"`
	Height            int64         `json:"height"`
	Src               string        `json:"src"`
	VariantIDS        []interface{} `json:"variant_ids"`
	AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
}

type Option struct {
	ID        float64  `json:"id"`
	ProductID float64  `json:"product_id"`
	Name      string   `json:"name"`
	Position  int64    `json:"position"`
	Values    []string `json:"values"`
}

type Variant struct {
	ID                   float64     `json:"id"`
	ProductID            float64     `json:"product_id"`
	Title                string      `json:"title"`
	Price                string      `json:"price"`
	Sku                  string      `json:"sku"`
	Position             int64       `json:"position"`
	InventoryPolicy      string      `json:"inventory_policy"`
	CompareAtPrice       string      `json:"compare_at_price"`
	FulfillmentService   string      `json:"fulfillment_service"`
	InventoryManagement  string      `json:"inventory_management"`
	Option1              string      `json:"option1"`
	Option2              interface{} `json:"option2"`
	Option3              interface{} `json:"option3"`
	CreatedAt            interface{} `json:"created_at"`
	UpdatedAt            interface{} `json:"updated_at"`
	Taxable              bool        `json:"taxable"`
	Barcode              interface{} `json:"barcode"`
	Grams                int64       `json:"grams"`
	ImageID              interface{} `json:"image_id"`
	Weight               float64     `json:"weight"`
	WeightUnit           string      `json:"weight_unit"`
	InventoryItemID      interface{} `json:"inventory_item_id"`
	InventoryQuantity    int64       `json:"inventory_quantity"`
	OldInventoryQuantity int64       `json:"old_inventory_quantity"`
	RequiresShipping     bool        `json:"requires_shipping"`
	AdminGraphqlAPIID    string      `json:"admin_graphql_api_id"`
}

func UnmarshalCreateShopifyProductsWHReq(data []byte) (CreateShopifyProductsWHReq, error) {
	var r CreateShopifyProductsWHReq
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type ShopifyOrderPaymentWHReq struct {
	ID                     float64               `json:"id"`
	Email                  string                `json:"email"`
	ClosedAt               interface{}           `json:"closed_at"`
	CreatedAt              string                `json:"created_at"`
	UpdatedAt              string                `json:"updated_at"`
	Number                 int64                 `json:"number"`
	Note                   interface{}           `json:"note"`
	Token                  string                `json:"token"`
	Gateway                interface{}           `json:"gateway"`
	Test                   bool                  `json:"test"`
	TotalPrice             string                `json:"total_price"`
	SubtotalPrice          string                `json:"subtotal_price"`
	TotalWeight            int64                 `json:"total_weight"`
	TotalTax               string                `json:"total_tax"`
	TaxesIncluded          bool                  `json:"taxes_included"`
	Currency               Currency              `json:"currency"`
	FinancialStatus        string                `json:"financial_status"`
	Confirmed              bool                  `json:"confirmed"`
	TotalDiscounts         string                `json:"total_discounts"`
	TotalLineItemsPrice    string                `json:"total_line_items_price"`
	CartToken              interface{}           `json:"cart_token"`
	BuyerAcceptsMarketing  bool                  `json:"buyer_accepts_marketing"`
	Name                   string                `json:"name"`
	ReferringSite          interface{}           `json:"referring_site"`
	LandingSite            interface{}           `json:"landing_site"`
	CancelledAt            string                `json:"cancelled_at"`
	CancelReason           string                `json:"cancel_reason"`
	TotalPriceUsd          interface{}           `json:"total_price_usd"`
	CheckoutToken          interface{}           `json:"checkout_token"`
	Reference              interface{}           `json:"reference"`
	UserID                 interface{}           `json:"user_id"`
	LocationID             interface{}           `json:"location_id"`
	SourceIdentifier       interface{}           `json:"source_identifier"`
	SourceURL              interface{}           `json:"source_url"`
	ProcessedAt            interface{}           `json:"processed_at"`
	DeviceID               interface{}           `json:"device_id"`
	Phone                  interface{}           `json:"phone"`
	CustomerLocale         string                `json:"customer_locale"`
	AppID                  interface{}           `json:"app_id"`
	BrowserIP              interface{}           `json:"browser_ip"`
	LandingSiteRef         interface{}           `json:"landing_site_ref"`
	OrderNumber            int64                 `json:"order_number"`
	DiscountApplications   []DiscountApplication `json:"discount_applications"`
	DiscountCodes          []interface{}         `json:"discount_codes"`
	NoteAttributes         []interface{}         `json:"note_attributes"`
	PaymentGatewayNames    []string              `json:"payment_gateway_names"`
	ProcessingMethod       string                `json:"processing_method"`
	CheckoutID             interface{}           `json:"checkout_id"`
	SourceName             string                `json:"source_name"`
	FulfillmentStatus      string                `json:"fulfillment_status"`
	TaxLines               []interface{}         `json:"tax_lines"`
	Tags                   string                `json:"tags"`
	ContactEmail           string                `json:"contact_email"`
	OrderStatusURL         string                `json:"order_status_url"`
	PresentmentCurrency    Currency              `json:"presentment_currency"`
	TotalLineItemsPriceSet Set                   `json:"total_line_items_price_set"`
	TotalDiscountsSet      Set                   `json:"total_discounts_set"`
	TotalShippingPriceSet  Set                   `json:"total_shipping_price_set"`
	SubtotalPriceSet       Set                   `json:"subtotal_price_set"`
	TotalPriceSet          Set                   `json:"total_price_set"`
	TotalTaxSet            Set                   `json:"total_tax_set"`
	LineItems              []LineItem            `json:"line_items"`
	Fulfillments           []interface{}         `json:"fulfillments"`
	Refunds                []interface{}         `json:"refunds"`
	TotalTipReceived       string                `json:"total_tip_received"`
	OriginalTotalDutiesSet interface{}           `json:"original_total_duties_set"`
	CurrentTotalDutiesSet  interface{}           `json:"current_total_duties_set"`
	AdminGraphqlAPIID      string                `json:"admin_graphql_api_id"`
	ShippingLines          []ShippingLine        `json:"shipping_lines"`
	BillingAddress         Address               `json:"billing_address"`
	ShippingAddress        Address               `json:"shipping_address"`
	Customer               Customer              `json:"customer"`
}

type Address struct {
	FirstName    *string     `json:"first_name"`
	Address1     string      `json:"address1"`
	Phone        string      `json:"phone"`
	City         string      `json:"city"`
	Zip          string      `json:"zip"`
	Province     string      `json:"province"`
	Country      string      `json:"country"`
	LastName     *string     `json:"last_name"`
	Address2     interface{} `json:"address2"`
	Company      *string     `json:"company"`
	Latitude     interface{} `json:"latitude"`
	Longitude    interface{} `json:"longitude"`
	Name         string      `json:"name"`
	CountryCode  string      `json:"country_code"`
	ProvinceCode string      `json:"province_code"`
	ID           *float64    `json:"id,omitempty"`
	CustomerID   *float64    `json:"customer_id,omitempty"`
	CountryName  *string     `json:"country_name,omitempty"`
	Default      *bool       `json:"default,omitempty"`
}

type Customer struct {
	ID                        int64       `json:"id"`
	Email                     string      `json:"email"`
	AcceptsMarketing          bool        `json:"accepts_marketing"`
	CreatedAt                 interface{} `json:"created_at"`
	UpdatedAt                 interface{} `json:"updated_at"`
	FirstName                 string      `json:"first_name"`
	LastName                  string      `json:"last_name"`
	OrdersCount               int64       `json:"orders_count"`
	State                     string      `json:"state"`
	TotalSpent                string      `json:"total_spent"`
	LastOrderID               interface{} `json:"last_order_id"`
	Note                      interface{} `json:"note"`
	VerifiedEmail             bool        `json:"verified_email"`
	MultipassIdentifier       interface{} `json:"multipass_identifier"`
	TaxExempt                 bool        `json:"tax_exempt"`
	Phone                     interface{} `json:"phone"`
	Tags                      string      `json:"tags"`
	LastOrderName             interface{} `json:"last_order_name"`
	Currency                  Currency    `json:"currency"`
	AcceptsMarketingUpdatedAt interface{} `json:"accepts_marketing_updated_at"`
	MarketingOptInLevel       interface{} `json:"marketing_opt_in_level"`
	AdminGraphqlAPIID         string      `json:"admin_graphql_api_id"`
	DefaultAddress            Address     `json:"default_address"`
}

type DiscountApplication struct {
	Type             string `json:"type"`
	Value            string `json:"value"`
	ValueType        string `json:"value_type"`
	AllocationMethod string `json:"allocation_method"`
	TargetSelection  string `json:"target_selection"`
	TargetType       string `json:"target_type"`
	Description      string `json:"description"`
	Title            string `json:"title"`
}

type LineItem struct {
	ID                         float64              `json:"id"`
	VariantID                  int64                `json:"variant_id"`
	Title                      string               `json:"title"`
	Quantity                   int64                `json:"quantity"`
	Sku                        string               `json:"sku"`
	VariantTitle               interface{}          `json:"variant_title"`
	Vendor                     interface{}          `json:"vendor"`
	FulfillmentService         string               `json:"fulfillment_service"`
	ProductID                  int64                `json:"product_id"`
	RequiresShipping           bool                 `json:"requires_shipping"`
	Taxable                    bool                 `json:"taxable"`
	GiftCard                   bool                 `json:"gift_card"`
	Name                       string               `json:"name"`
	VariantInventoryManagement string               `json:"variant_inventory_management"`
	Properties                 []interface{}        `json:"properties"`
	ProductExists              bool                 `json:"product_exists"`
	FulfillableQuantity        int64                `json:"fulfillable_quantity"`
	Grams                      int64                `json:"grams"`
	Price                      string               `json:"price"`
	TotalDiscount              string               `json:"total_discount"`
	FulfillmentStatus          interface{}          `json:"fulfillment_status"`
	PriceSet                   Set                  `json:"price_set"`
	TotalDiscountSet           Set                  `json:"total_discount_set"`
	DiscountAllocations        []DiscountAllocation `json:"discount_allocations"`
	Duties                     []interface{}        `json:"duties"`
	AdminGraphqlAPIID          string               `json:"admin_graphql_api_id"`
	TaxLines                   []interface{}        `json:"tax_lines"`
}

type DiscountAllocation struct {
	Amount                   string `json:"amount"`
	DiscountApplicationIndex int64  `json:"discount_application_index"`
	AmountSet                Set    `json:"amount_set"`
}

type Set struct {
	ShopMoney        Money `json:"shop_money"`
	PresentmentMoney Money `json:"presentment_money"`
}

type Money struct {
	Amount       string   `json:"amount"`
	CurrencyCode Currency `json:"currency_code"`
}

type ShippingLine struct {
	ID                            float64       `json:"id"`
	Title                         string        `json:"title"`
	Price                         string        `json:"price"`
	Code                          interface{}   `json:"code"`
	Source                        string        `json:"source"`
	Phone                         interface{}   `json:"phone"`
	RequestedFulfillmentServiceID interface{}   `json:"requested_fulfillment_service_id"`
	DeliveryCategory              interface{}   `json:"delivery_category"`
	CarrierIdentifier             interface{}   `json:"carrier_identifier"`
	DiscountedPrice               string        `json:"discounted_price"`
	PriceSet                      Set           `json:"price_set"`
	DiscountedPriceSet            Set           `json:"discounted_price_set"`
	DiscountAllocations           []interface{} `json:"discount_allocations"`
	TaxLines                      []interface{} `json:"tax_lines"`
}

type Currency string

const (
	Gbp Currency = "GBP"
)

func UnmarshalShopifyOrderPaymentWHReq(data []byte) (ShopifyOrderPaymentWHReq, error) {
	var r ShopifyOrderPaymentWHReq
	err := json.Unmarshal(data, &r)
	return r, err
}
