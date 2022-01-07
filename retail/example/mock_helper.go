package main

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/byteplus-sdk/byteplus-sdk-go-rec/retail/protocol"
)

func mockUsers(count int) []*DemoUser {
	users := make([]*DemoUser, count)
	for i := 0; i < count; i++ {
		user := mockUser()
		user.UserId += strconv.Itoa(i)
		users[i] = user
	}
	return users
}

func mockUser() *DemoUser {
	tags, _ := json.Marshal([]string{"tag1", "tag2", "tag3"})
	return &DemoUser{
		UserId:                 "user_id",
		Gender:                 "male",
		Age:                    "23",
		Tags:                   string(tags),
		ActivationChannel:      "AppStore",
		MembershipLevel:        "silver",
		RegistrationTimestamp:  time.Now().Unix(),
		LocationCity:           "beijing",
		LocationCountry:        "china",
		LocationDistrictOrArea: "haidian",
		LocationPostcode:       "123456",
		//CustomField:                   "custom",
	}
}

func mockProducts(count int) []*DemoProduct {
	products := make([]*DemoProduct, count)
	for i := 0; i < count; i++ {
		product := mockProduct()
		product.ProductId += strconv.Itoa(i)
		products[i] = product
	}
	return products
}

func mockProduct() *DemoProduct {
	category1Node1 := &ProductCategoryCategoryNode{
		IdOrName: "cate_1_1",
	}
	category1 := &Category{
		CategoryDepth: 1,
		CategoryNodes: []*ProductCategoryCategoryNode{category1Node1},
	}
	category2Node1 := &ProductCategoryCategoryNode{
		IdOrName: "cate_2_1",
	}
	category2Node2 := &ProductCategoryCategoryNode{
		IdOrName: "cate_2_2",
	}
	category2 := &Category{
		CategoryDepth: 2,
		CategoryNodes: []*ProductCategoryCategoryNode{category2Node1, category2Node2},
	}
	category := []*Category{category1, category2}
	categoryStr, _ := json.Marshal(category)
	tags := []string{"tag1", "tag2", "tag3"}
	tagsStr, _ := json.Marshal(tags)
	displayDetailPageDisplayTags := []string{"tag1", "tag2"}
	displayDetailPageDisplayTagsStr, _ := json.Marshal(displayDetailPageDisplayTags)
	displayListingPageDisplayTags := []string{"taga", "tagb"}
	displayListingPageDisplayTagsStr, _ := json.Marshal(displayListingPageDisplayTags)
	return &DemoProduct{
		ProductId:                     "product_id",
		Category:                      string(categoryStr),
		Brands:                        "brand_1",
		PriceCurrentPrice:             1000,
		PriceOriginPrice:              1000,
		IsRecommendable:               true,
		Title:                         "title",
		QualityScore:                  3.4,
		Tags:                          string(tagsStr),
		DisplayDetailPageDisplayTags:  string(displayDetailPageDisplayTagsStr),
		DisplayCoverMultimediaUrl:     "https://www.google.com",
		DisplayListingPageDisplayTags: string(displayListingPageDisplayTagsStr),
		DisplayListingPageDisplayType: "image",
		ProductSpecProductGroupId:     "group_id",
		ProductSpecCommentCount:       100,
		ProductSpecPublishTimestamp:   time.Now().Unix(),
		ProductSpecSource:             "self",
		ProductSpecUserRating:         0.23,
		SellerId:                      "seller_id",
		SellerSellerLevel:             "level1",
		SellerSellerRating:            3.5,
		//CustomField:                   "custom",
	}
}

func mockUserEvents(count int) []*DemoUserEvent {
	userEvents := make([]*DemoUserEvent, count)
	for i := 0; i < count; i++ {
		userEvents[i] = mockUserEvent()
	}
	return userEvents
}

func mockUserEvent() *DemoUserEvent {
	return &DemoUserEvent{
		UserId:               "user_id",
		EventType:            "purchase",
		EventTimestamp:       time.Now().Unix(),
		Scene:                "scene_name",
		ScenePageNumber:      2,
		SceneOffset:          10,
		ProductId:            "product_id",
		DevicePlatform:       "app",
		DeviceOsType:         "android",
		DeviceAppVersion:     "app_version",
		DeviceDeviceModel:    "device_model",
		DeviceDeviceBrand:    "device_brand",
		DeviceOsVersion:      "os_version",
		DeviceBrowserType:    "firefox",
		DeviceUserAgent:      "user_agent",
		DeviceNetwork:        "3g",
		ContextQuery:         "query",
		ContextRootProductId: "root_product_id",
		AttributionToken:     "attribution_token",
		RecInfo:              "trans_data",
		TrafficSource:        "self",
		PurchaseCount:        20,
		// 10s
		DetailPageStayTime: 10,
		//CustomField:                   "custom",
	}
}

func mockPredictProduct() *protocol.Product {
	category1Node1 := &protocol.Product_Category_CategoryNode{
		IdOrName: "cate_1_1",
	}
	category1 := &protocol.Product_Category{
		CategoryDepth: 1,
		CategoryNodes: []*protocol.Product_Category_CategoryNode{category1Node1},
	}
	category2Node1 := &protocol.Product_Category_CategoryNode{
		IdOrName: "cate_2_1",
	}
	category2Node2 := &protocol.Product_Category_CategoryNode{
		IdOrName: "cate_2_2",
	}
	category2 := &protocol.Product_Category{
		CategoryDepth: 2,
		CategoryNodes: []*protocol.Product_Category_CategoryNode{category2Node1, category2Node2},
	}

	brand1 := &protocol.Product_Brand{
		BrandDepth: 1,
		IdOrName:   "brand_1",
	}
	brand2 := &protocol.Product_Brand{
		BrandDepth: 2,
		IdOrName:   "brand_2",
	}

	price := &protocol.Product_Price{
		CurrentPrice: 10,
		OriginPrice:  10,
	}

	display := &protocol.Product_Display{
		DetailPageDisplayTags:  []string{"tag1", "tag2"},
		ListingPageDisplayTags: []string{"taga", "tagb"},
		ListingPageDisplayType: "image",
		CoverMultimediaUrl:     "https://www.google.com",
	}

	spec := &protocol.Product_ProductSpec{
		ProductGroupId:   "group_id",
		UserRating:       0.23,
		CommentCount:     100,
		Source:           "self",
		PublishTimestamp: time.Now().Unix(),
	}

	seller := &protocol.Product_Seller{
		Id:           "seller_id",
		SellerLevel:  "level1",
		SellerRating: 3.5,
	}

	return &protocol.Product{
		ProductId:       "product_id",
		Categories:      []*protocol.Product_Category{category1, category2},
		Brands:          []*protocol.Product_Brand{brand1, brand2},
		Price:           price,
		IsRecommendable: true,
		Title:           "title",
		QualityScore:    3.4,
		Tags:            []string{"tag1", "tag2", "tag3"},
		Display:         display,
		ProductSpec:     spec,
		Seller:          seller,
		Extra:           map[string]string{"count": "20"},
	}
}

func mockPredictDevice() *protocol.Device {
	return &protocol.Device{
		Platform:    "app",
		OsType:      "android",
		AppVersion:  "app_version",
		DeviceModel: "device_model",
		DeviceBrand: "device_brand",
		OsVersion:   "os_version",
		BrowserType: "firefox",
		UserAgent:   "user_agent",
		Network:     "3g",
	}
}
