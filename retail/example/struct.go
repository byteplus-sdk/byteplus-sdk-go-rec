package main

type DemoUser struct {
	UserId                string `json:"user_id,omitempty"`
	Gender                string `json:"gender,omitempty"`
	Age                   string `json:"age,omitempty"`
	Tags                  string `json:"tags,omitempty"` // json array
	RegistrationTimestamp int64  `json:"registration_timestamp,omitempty"`
	ActivationChannel     string `json:"activation_channel,omitempty"`
	MembershipLevel       string `json:"membership_level,omitempty"`
	Language              string `json:"language,omitempty"`
	City                  string `json:"city,omitempty"`
	Country               string `json:"country,omitempty"`
	District              string `json:"district,omitempty"`
	Province              string `json:"province,omitempty"`
	CustomField           string `json:"custom_field,omitempty"`
}

type DemoProduct struct {
	ProductId                 string  `json:"product_id,omitempty"`
	IsRecommendable           int32   `json:"is_recommendable,omitempty"`
	CurrentPrice              float32 `json:"current_price,omitempty"`
	Categories                string  `json:"categories,omitempty"` // json array
	Brands                    string  `json:"brands,omitempty"`
	Title                     string  `json:"title,omitempty"`
	OriginalPrice             float32 `json:"original_price,omitempty"`
	Tags                      string  `json:"tags,omitempty"`                         // json array
	DisplayCoverMultimediaUrl string  `json:"display_cover_multimedia_url,omitempty"` // json array
	ProductGroupId            string  `json:"product_group_id,omitempty"`
	UserRating                float32 `json:"user_rating,omitempty"`
	CommentCount              int32   `json:"comment_count,omitempty"`
	SoldCount                 int32   `json:"sold_count,omitempty"`
	Source                    string  `json:"source,omitempty"`
	PublishTimestamp          int64   `json:"publish_timestamp,omitempty"`
	SellerId                  string  `json:"seller_id,omitempty"`
	SellerLevel               string  `json:"seller_level,omitempty"`
	SellerRating              float32 `json:"seller_rating,omitempty"`
	CustomField               string  `json:"custom_field,omitempty"`
}

type DemoUserEvent struct {
	UserId           string  `json:"user_id,omitempty"`
	EventType        string  `json:"event_type,omitempty"`
	EventTimestamp   int64   `json:"event_timestamp,omitempty"`
	SceneName        string  `json:"scene_name,omitempty"`
	PageNumber       int32   `json:"page_number,omitempty"`
	Offset           int32   `json:"offset,omitempty"`
	ProductId        string  `json:"product_id,omitempty"`
	Platform         string  `json:"platform,omitempty"`
	OsType           string  `json:"os_type,omitempty"`
	AppVersion       string  `json:"app_version,omitempty"`
	DeviceModel      string  `json:"device_model,omitempty"`
	OsVersion        string  `json:"os_version,omitempty"`
	Network          string  `json:"network,omitempty"`
	Query            string  `json:"query,omitempty"`
	ParentProductId  string  `json:"parent_product_id,omitempty"`
	AttributionToken string  `json:"attribution_token,omitempty"`
	TrafficSource    string  `json:"traffic_source,omitempty"`
	PurchaseCount    int32   `json:"purchase_count,omitempty"`
	PaidPrice        float32 `json:"paid_price,omitempty"`
	Currency         string  `json:"currency,omitempty"`
	City             string  `json:"city,omitempty"`
	Country          string  `json:"country,omitempty"`
	District         string  `json:"district,omitempty"`
	Province         string  `json:"province,omitempty"`
	CustomField      string  `json:"custom_field,omitempty"`
}
