package main

type DemoUser struct {
	UserId                 string `json:"user_id,omitempty"`
	Gender                 string `json:"gender,omitempty"`
	Age                    string `json:"age,omitempty"`
	Tags                   string `json:"tags,omitempty"` // json array
	ActivationChannel      string `json:"activation_channel,omitempty"`
	MembershipLevel        string `json:"membership_level,omitempty"`
	RegistrationTimestamp  int64  `json:"registration_timestamp,omitempty"`
	LocationCity           string `json:"location_city,omitempty"`
	LocationCountry        string `json:"location_country,omitempty"`
	LocationDistrictOrArea string `json:"location_district_or_area,omitempty"`
	LocationPostcode       string `json:"location_postcode,omitempty"`
	CustomField            string `json:"custom_field,omitempty"`
}

type DemoProduct struct {
	ProductId                     string  `json:"product_id,omitempty"`
	Category                      string  `json:"category,omitempty"` // json array
	Brands                        string  `json:"brands,omitempty"`
	IsRecommendable               int32   `json:"is_recommendable,omitempty"`
	Title                         string  `json:"title,omitempty"`
	PriceCurrentPrice             int64   `json:"price_current_price,omitempty"`
	PriceOriginPrice              int64   `json:"price_origin_price,omitempty"`
	QualityScore                  float64 `json:"quality_score,omitempty"`
	Tags                          string  `json:"tags,omitempty"` // json array
	DisplayCoverMultimediaUrl     string  `json:"display_cover_multimedia_url,omitempty"`
	DisplayListingPageDisplayType string  `json:"display_listing_page_display_type,omitempty"`
	DisplayListingPageDisplayTags string  `json:"display_listing_page_display_tags,omitempty"` // json array
	DisplayDetailPageDisplayTags  string  `json:"display_detail_page_display_tags,omitempty"`  // json array
	SellerId                      string  `json:"seller_id,omitempty"`
	SellerSellerLevel             string  `json:"seller_seller_level,omitempty"`
	SellerSellerRating            float64 `json:"seller_seller_rating,omitempty"`
	ProductSpecProductGroupId     string  `json:"product_spec_product_group_id,omitempty"`
	ProductSpecUserRating         float64 `json:"product_spec_user_rating,omitempty"`
	ProductSpecCommentCount       int32   `json:"product_spec_comment_count,omitempty"`
	ProductSpecSource             string  `json:"product_spec_source,omitempty"`
	ProductSpecPublishTimestamp   int64   `json:"product_spec_publish_timestamp,omitempty"`
	CustomField                   string  `json:"custom_field,omitempty"`
}

type DemoUserEvent struct {
	UserId               string `json:"user_id,omitempty"`
	EventType            string `json:"event_type,omitempty"`
	EventTimestamp       int64  `json:"event_timestamp,omitempty"`
	Scene                string `json:"scene,omitempty"`
	ScenePageNumber      int32  `json:"scene_page_number,omitempty"`
	SceneOffset          int32  `json:"scene_offset,omitempty"`
	ProductId            string `json:"product_id,omitempty"`
	DevicePlatform       string `json:"device_platform,omitempty"`
	DeviceOsType         string `json:"device_os_type,omitempty"`
	DeviceAppVersion     string `json:"device_app_version,omitempty"`
	DeviceDeviceModel    string `json:"device_device_model,omitempty"`
	DeviceDeviceBrand    string `json:"device_device_brand,omitempty"`
	DeviceOsVersion      string `json:"device_os_version,omitempty"`
	DeviceBrowserType    string `json:"device_browser_type,omitempty"`
	DeviceUserAgent      string `json:"device_user_agent,omitempty"`
	DeviceNetwork        string `json:"device_network,omitempty"`
	ContextQuery         string `json:"context_query,omitempty"`
	ContextRootProductId string `json:"context_root_product_id,omitempty"`
	AttributionToken     string `json:"attribution_token,omitempty"`
	RecInfo              string `json:"rec_info,omitempty"`
	TrafficSource        string `json:"traffic_source,omitempty"`
	PurchaseCount        int32  `json:"purchase_count,omitempty"`
	DetailPageStayTime   int32  `json:"detail_page_stay_time,omitempty"`
	CustomField          string `json:"custom_field,omitempty"`
}
