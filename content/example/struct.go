package main

type DemoUser struct {
	UserId                string `json:"user_id,omitempty"`
	UserIdType            string `json:"user_id_type,omitempty"`
	Gender                string `json:"gender,omitempty"`
	Age                   string `json:"age,omitempty"`
	Tags                  string `json:"tags,omitempty"` // json array
	Language              string `json:"language,omitempty"`
	SubscriberType        string `json:"subscriber_type,omitempty"`
	NetWork               string `json:"net_work,omitempty"`
	Platform              string `json:"platform,omitempty"`
	OsType                string `json:"os_type,omitempty"`
	AppVersion            string `json:"app_version,omitempty"`
	DeviceModel           string `json:"device_model,omitempty"`
	OsVersion             string `json:"os_version,omitempty"`
	MembershipLevel       string `json:"membership_level,omitempty"`
	RegistrationTimestamp int32  `json:"registration_timestamp,omitempty"`
	UpdateTimestamp       int32  `json:"update_timestamp,omitempty"`
	LastLoginTimestamp    int32  `json:"last_login_timestamp,omitempty"`
	Country               string `json:"country,omitempty"`
	Province              string `json:"province,omitempty"`
	City                  string `json:"city,omitempty"`
	District              string `json:"district,omitempty"`
	Area                  string `json:"area,omitempty"`
	CustomField           string `json:"custom_field,omitempty"`
}

type DemoContent struct {
	ContentId               string  `json:"content_id,omitempty"`
	IsRecommendable         int32   `json:"is_recommendable,omitempty"`
	Categories              string  `json:"categories,omitempty"` // json array
	ContentType             string  `json:"content_type,omitempty"`
	VideoDuration           int32   `json:"video_duration,omitempty"`
	ContentTitle            string  `json:"content_title,omitempty"`
	Description             string  `json:"description,omitempty"`
	ContentOwner            string  `json:"content_owner,omitempty"`
	ContentOwnerFollowers   int32   `json:"content_owner_followers,omitempty"`
	ContentOwnerRating      float32 `json:"content_owner_rating,omitempty"`
	ContentOwnerName        string  `json:"content_owner_name,omitempty"`
	CollectionId            string  `json:"collection_id,omitempty"`
	Tags                    string  `json:"tags,omitempty"`       // json array
	TopicTags               string  `json:"topic_tags,omitempty"` // json array
	ImageUrls               string  `json:"image_urls,omitempty"` // json array
	DetailPicNum            int32   `json:"detail_pic_num,omitempty"`
	VideoUrls               string  `json:"video_urls,omitempty"` // json array
	UserRating              float32 `json:"user_rating,omitempty"`
	ViewsCount              int32   `json:"views_count,omitempty"`
	CommentsCount           int32   `json:"comments_count,omitempty"`
	LikesCount              int32   `json:"likes_count,omitempty"`
	SharesCount             int32   `json:"shares_count,omitempty"`
	SaveCount               int32   `json:"save_count,omitempty"`
	CurrentPrice            int32   `json:"current_price,omitempty"`
	OriginalPrice           int32   `json:"original_price,omitempty"`
	AvailableLocation       string  `json:"available_location,omitempty"` // json array
	PublishTimestamp        int32   `json:"publish_timestamp,omitempty"`
	UpdateTimestamp         int32   `json:"update_timestamp,omitempty"`
	CopyrightStartTimestamp int32   `json:"copyright_start_timestamp,omitempty"`
	CopyrightEndTimestamp   int32   `json:"copyright_end_timestamp,omitempty"`
	IsPaidContent           bool    `json:"is_paid_content,omitempty"`
	Language                string  `json:"language,omitempty"`
	RelatedContentIds       string  `json:"related_content_ids,omitempty"` // json array
	SoldCount               int32   `json:"sold_count,omitempty"`
	Source                  string  `json:"source,omitempty"`
	CustomField             string  `json:"custom_field,omitempty"`
}

type DemoUserEvent struct {
	UserId           string `json:"user_id,omitempty"`
	EventType        string `json:"event_type,omitempty"`
	EventTimestamp   int32  `json:"event_timestamp,omitempty"`
	ContentId        string `json:"content_id,omitempty"`
	TrafficSource    string `json:"traffic_source,omitempty"`
	RequestId        string `json:"request_id,omitempty"`
	RecInfo          string `json:"rec_info,omitempty"`
	AttributionToken string `json:"attribution_token,omitempty"`
	SceneName        string `json:"scene_name,omitempty"`
	PageNumber       int32  `json:"page_number,omitempty"`
	Offset           int32  `json:"offset,omitempty"`
	PlayDuration     int32  `json:"play_duration,omitempty"`
	VideoDuration    int32  `json:"video_duration,omitempty"`
	StartTime        int32  `json:"start_time,omitempty"`
	EndTime          int32  `json:"end_time,omitempty"`
	ParentContentId  string `json:"parent_content_id,omitempty"`
	ContentOwnerId   string `json:"content_owner_id,omitempty"`
	DetailStayTime   int32  `json:"detail_stay_time,omitempty"`
	DislikeType      string `json:"dislike_type,omitempty"`
	DislikeValue     string `json:"dislike_value,omitempty"`
	Query            string `json:"query,omitempty"`
	Platform         string `json:"platform,omitempty"`
	OsType           string `json:"os_type,omitempty"`
	AppVersion       string `json:"app_version,omitempty"`
	DeviceModel      string `json:"device_model,omitempty"`
	OsVersion        string `json:"os_version,omitempty"`
	Network          string `json:"network,omitempty"`
	Country          string `json:"country,omitempty"`
	Province         string `json:"province,omitempty"`
	City             string `json:"city,omitempty"`
	District         string `json:"district,omitempty"`
	Area             string `json:"area,omitempty"`
	CustomField      string `json:"custom_field,omitempty"`
}
