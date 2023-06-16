package main

type DemoUser struct {
	UserId                string `json:"user_id,omitempty"`
	Gender                string `json:"gender,omitempty"`
	Age                   string `json:"age,omitempty"`
	Tags                  string `json:"tags,omitempty"` // json array
	Language              string `json:"language,omitempty"`
	SubscriberType        string `json:"subscriber_type,omitempty"`
	MembershipLevel       string `json:"membership_level,omitempty"`
	RegistrationTimestamp int64  `json:"registration_timestamp,omitempty"`
	Country               string `json:"country,omitempty"`
	Province              string `json:"province,omitempty"`
	City                  string `json:"city,omitempty"`
	District              string `json:"district,omitempty"`
	CustomField           string `json:"custom_field,omitempty"`
}

type DemoContent struct {
	ContentId        string  `json:"content_id,omitempty"`
	IsRecommendable  int32   `json:"is_recommendable,omitempty"`
	ContentType      string  `json:"content_type,omitempty"`
	Categories       string  `json:"categories,omitempty"` // json array
	VideoDuration    int32   `json:"video_duration,omitempty"`
	ContentTitle     string  `json:"content_title,omitempty"`
	Description      string  `json:"description,omitempty"`
	ContentOwner     string  `json:"content_owner,omitempty"`
	CollectionId     string  `json:"collection_id,omitempty"`
	Tags             string  `json:"tags,omitempty"`       // json array
	ImageUrls        string  `json:"image_urls,omitempty"` // json array
	VideoUrls        string  `json:"video_urls,omitempty"` // json array
	UserRating       float32 `json:"user_rating,omitempty"`
	CurrentPrice     float32 `json:"current_price,omitempty"`
	OriginalPrice    float32 `json:"original_price,omitempty"`
	PublishTimestamp int64   `json:"publish_timestamp,omitempty"`
	IsPaidContent    bool    `json:"is_paid_content,omitempty"`
	Language         string  `json:"language,omitempty"`
	LinkedProductId  string  `json:"linked_product_id,omitempty"` // json array
	Source           string  `json:"source,omitempty"`
	CustomField      string  `json:"custom_field,omitempty"`
}

type DemoUserEvent struct {
	UserId           string  `json:"user_id,omitempty"`
	EventType        string  `json:"event_type,omitempty"`
	EventTimestamp   int64   `json:"event_timestamp,omitempty"`
	ContentId        string  `json:"content_id,omitempty"`
	TrafficSource    string  `json:"traffic_source,omitempty"`
	AttributionToken string  `json:"attribution_token,omitempty"`
	SceneName        string  `json:"scene_name,omitempty"`
	PageNumber       int32   `json:"page_number,omitempty"`
	Offset           int32   `json:"offset,omitempty"`
	StayDuration     int32   `json:"stay_duration,omitempty"`
	ParentContentId  string  `json:"parent_content_id,omitempty"`
	ContentOwnerId   string  `json:"content_owner_id,omitempty"`
	Query            string  `json:"query,omitempty"`
	Platform         string  `json:"platform,omitempty"`
	OsType           string  `json:"os_type,omitempty"`
	AppVersion       string  `json:"app_version,omitempty"`
	DeviceModel      string  `json:"device_model,omitempty"`
	OsVersion        string  `json:"os_version,omitempty"`
	Network          string  `json:"network,omitempty"`
	Country          string  `json:"country,omitempty"`
	Province         string  `json:"province,omitempty"`
	City             string  `json:"city,omitempty"`
	District         string  `json:"district,omitempty"`
	PurchaseCount    int32   `json:"purchase_count,omitempty"`
	PaidPrice        float32 `json:"paid_price,omitempty"`
	Currency         string  `json:"currency,omitempty"`
	CustomField      string  `json:"custom_field,omitempty"`
}
