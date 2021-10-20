package rest

import "encoding/json"

// ** Octy REST Request Models **

// ---

type CreateOctyProfileReq struct {
	Profiles []ProfileReq `json:"profiles"`
}

type ProfileReq struct {
	CustomerID   string      `json:"customer_id"`
	HasCharged   bool        `json:"has_charged"`
	ProfileData  interface{} `json:"profile_data"`
	PlatformInfo interface{} `json:"platform_info"`
}

func (r *CreateOctyProfileReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type UpdateOctyProfileReq struct {
	Profiles []UProfileReq `json:"profiles"`
}

type UProfileReq struct {
	ProfileID    string      `json:"profile_id"`
	CustomerID   string      `json:"customer_id"`
	ProfileData  interface{} `json:"profile_data"`
	PlatformInfo interface{} `json:"platform_info"`
	Status       string      `json:"status"`
	HasCharged   bool        `json:"has_charged"`
}

func (r *UpdateOctyProfileReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type CreateOctyEventReq struct {
	EventType       string                 `json:"event_type"`
	EventProperties map[string]interface{} `json:"event_properties"`
	ProfileID       string                 `json:"profile_id"`
}

func (r *CreateOctyEventReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type CreateOctyItemsReq struct {
	Items []CItemReq `json:"items"`
}

type CItemReq struct {
	ItemID          string `json:"item_id"`
	ItemCategory    string `json:"item_category"`
	ItemName        string `json:"item_name"`
	ItemDescription string `json:"item_description"`
	ItemPrice       int64  `json:"item_price"`
}

func (r *CreateOctyItemsReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type UpdateOctyItemsReq struct {
	Items []UItemReq `json:"items"`
}

type UItemReq struct {
	ItemID          string `json:"item_id"`
	ItemCategory    string `json:"item_category"`
	ItemName        string `json:"item_name"`
	ItemDescription string `json:"item_description"`
	ItemPrice       int64  `json:"item_price"`
	Status          string `json:"status"`
}

func (r *UpdateOctyItemsReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type GenOctyContentReq struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	TemplateID         string                   `json:"template_id"`
	ItemRecommendation bool                     `json:"item_recommendation"`
	Data               []map[string]interface{} `json:"data"`
}

func (r *GenOctyContentReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

type GetOctyRecReq struct {
	ProfileIDS []string `json:"profile_ids"`
}

func (r *GetOctyRecReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// ---

// ** Octy REST Response Models **

// ---
type CreateOctyProfileResp struct {
	RequestMeta    RequestMeta    `json:"request_meta"`
	Profiles       []CProfileResp `json:"profiles"`
	FailedToCreate []interface{}  `json:"failed_to_create"`
}

type CProfileResp struct {
	ProfileID    string      `json:"profile_id"`
	CustomerID   string      `json:"customer_id"`
	ProfileData  interface{} `json:"profile_data"`
	PlatformInfo interface{} `json:"platform_info"`
	HasCharged   bool        `json:"has_charged"`
	Status       string      `json:"status"`
	CreatedAt    string      `json:"created_at"`
}

type RequestMeta struct {
	RequestStatus string `json:"request_status"`
	Message       string `json:"message"`
	Count         int64  `json:"count"`
}

func UnmarshalCreateOctyProfileResp(data []byte) (CreateOctyProfileResp, error) {
	var r CreateOctyProfileResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type UpdateOctyProfileResp struct {
	RequestMeta    RequestMeta    `json:"request_meta"`
	Profiles       []UProfileResp `json:"profiles"`
	FailedToUpdate []interface{}  `json:"failed_to_update"`
}

type UProfileResp struct {
	ProfileID    string      `json:"profile_id"`
	CustomerID   string      `json:"customer_id"`
	ProfileData  interface{} `json:"profile_data"`
	PlatformInfo interface{} `json:"platform_info"`
	HasCharged   bool        `json:"has_charged"`
	Status       string      `json:"status"`
	CreatedAt    string      `json:"created_at"`
	LastUpdated  string      `json:"last_updated"`
}

func UnmarshalUpdateOctyProfileResp(data []byte) (UpdateOctyProfileResp, error) {
	var r UpdateOctyProfileResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type GetOctyProfileResp struct {
	RequestMeta RequestMeta `json:"request_meta"`
	Profiles    []GProfile  `json:"profiles"`
}

type GProfile struct {
	ProfileID        string        `json:"profile_id"`
	CustomerID       string        `json:"customer_id"`
	ProfileData      interface{}   `json:"profile_data"`
	PlatformInfo     interface{}   `json:"platform_info"`
	SegmentTags      []interface{} `json:"segment_tags"`
	Rfm              interface{}   `json:"rfm_score"`
	ChurnProbability string        `json:"churn_probability"`
	HasCharged       bool          `json:"has_charged"`
	Status           string        `json:"status"`
	CreatedAt        string        `json:"created_at"`
	UpdatedAt        string        `json:"updated_at"`
}

func UnmarshalGetOctyProfileResp(data []byte) (GetOctyProfileResp, error) {
	var r GetOctyProfileResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type OctyProfileMetaResp struct {
	RequestMeta  RequestMeta    `json:"request_meta"`
	ProfilesMeta []ProfilesMeta `json:"profiles_meta"`
}

type ProfilesMeta struct {
	ProvidedIdentifier string     `json:"provided_identifier"`
	Profile            MProfile   `json:"profile"`
	MergedInfo         MergedInfo `json:"merged_info"`
}

type MergedInfo struct {
	WasMerged            bool                 `json:"was_merged"`
	MergedAt             *string              `json:"merged_at"`
	AuthenticatedIDKey   *string              `json:"authenticated_id_key"`
	AuthenticatedIDValue *string              `json:"authenticated_id_value"`
	ParentOrChild        *string              `json:"parent_or_child"`
	ParentProfile        ParentProfile        `json:"parent_profile"`
	MergedChildProfiles  []MergedChildProfile `json:"merged_child_profiles"`
}

type MergedChildProfile struct {
	ProfileID  string `json:"profile_id"`
	CustomerID string `json:"customer_id"`
}

type ParentProfile struct {
	ParentProfileID  *string `json:"parent_profile_id"`
	ParentCustomerID *string `json:"parent_customer_id"`
}

type MProfile struct {
	ProfileExists bool    `json:"profile_exists"`
	ProfileID     *string `json:"profile_id"`
	CustomerID    *string `json:"customer_id"`
	CreatedAt     *string `json:"created_at"`
	UpdatedAt     *string `json:"updated_at"`
}

type IdentifyOctyProfileResp struct {
	Identifiers []Identifier `json:"identifiers"`
}

type Identifier struct {
	Identifier           string  `json:"identifier"`
	ParentProfileID      *string `json:"parent_profile_id"`
	ParentCustomerID     *string `json:"parent_customer_id"`
	AuthenticatedIDKey   *string `json:"authenticated_id_key"`
	AuthenticatedIDValue *string `json:"authenticated_id_value"`
}

func UnmarshalOctyProfileMetaResp(data []byte) (OctyProfileMetaResp, error) {
	var r OctyProfileMetaResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type CreateOctyEventResp struct {
	RequestMeta     RequestMeta `json:"request_meta"`
	EventID         string      `json:"event_id"`
	EventType       string      `json:"event_type"`
	EventProperties interface{} `json:"event_properties"`
	ProfileID       string      `json:"profile_id"`
	CreatedAt       string      `json:"created_at"`
}

func UnmarshalCreateOctyEventResp(data []byte) (CreateOctyEventResp, error) {
	var r CreateOctyEventResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type GetOctyTemplatesResp struct {
	RequestMeta RequestMeta `json:"request_meta"`
	Templates   []Template  `json:"templates"`
}

type Template struct {
	FriendlyName  string      `json:"friendly_name"`
	TemplateType  string      `json:"template_type"`
	Title         string      `json:"title"`
	Content       string      `json:"content"`
	RequiredData  []string    `json:"required_data"`
	DefaultValues interface{} `json:"default_values"`
	Metadata      interface{} `json:"metadata"`
	Status        string      `json:"status"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	TemplateID    string      `json:"template_id"`
}

func UnmarshalGetOctyTemplatesResp(data []byte) (GetOctyTemplatesResp, error) {
	var r GetOctyTemplatesResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type GenOctyContentResp struct {
	RequestMeta       RequestMeta               `json:"request_meta"`
	GeneratedMessages []GeneratedMessage        `json:"generated_messages"`
	FailedMessages    []FailedTemplatesMessages `json:"failed_messages"`
	FailedTemplates   []FailedTemplatesMessages `json:"failed templates"`
}

type GeneratedMessage struct {
	TemplateID   string `json:"template_id"`
	FriendlyName string `json:"friendly_name"`
	Title        string `json:"title"`
	Content      string `json:"content"`
}

type FailedTemplatesMessages struct {
	TemplateID string `json:"template_id"`
}

func UnmarshalGenOctyContentResp(data []byte) (GenOctyContentResp, error) {
	var r GenOctyContentResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---

type GetOctyRecResp struct {
	RequestMeta     RequestMeta                    `json:"request_meta"`
	Recommendations []GetOctyRecRespRecommendation `json:"recommendations"`
	ModelMetaData   ModelMetaData                  `json:"model_meta_data"`
}

type ModelMetaData struct {
	TrainingJobID        string  `json:"training_job_id"`
	ModelAccuracyScore   float64 `json:"model_accuracy_score"`
	RecommenderEventType string  `json:"recommender_event_type"`
	ModelCreatedAt       string  `json:"model_created_at"`
}

type GetOctyRecRespRecommendation struct {
	ProfileID       string                         `json:"profile_id"`
	Recommendations []RecommendationRecommendation `json:"recommendations"`
	Error           interface{}                    `json:"error"`
}

type RecommendationRecommendation struct {
	ItemID string  `json:"item_id"`
	Score  float64 `json:"score"`
}

func UnmarshalGetOctyRecResp(data []byte) (GetOctyRecResp, error) {
	var r GetOctyRecResp
	err := json.Unmarshal(data, &r)
	return r, err
}

// ---
