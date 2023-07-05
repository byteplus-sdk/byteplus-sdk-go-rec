package main

import (
	"strconv"

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
	return &DemoUser{
		UserId:                "1457789",
		Gender:                "male",
		Age:                   "23",
		Tags:                  `["new user","low purchasing power","bargain seeker"]`,
		ActivationChannel:     "AppStore",
		MembershipLevel:       "silver",
		RegistrationTimestamp: 1651818792,
		City:                  "Kirkland",
		Country:               "USA",
		District:              "King County",
		Province:              "Texas",
		Language:              "English",
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
	return &DemoProduct{
		ProductId:                 "632461",
		IsRecommendable:           1,
		CurrentPrice:              49.99,
		Categories:                `[{"category_depth":1,"category_nodes":[{"id_or_name":"Shoes"}]},{"category_depth":2,"category_nodes":[{"id_or_name":"Men's Shoes"}]}]`,
		Brands:                    "Adidas",
		OriginalPrice:             69.98,
		Title:                     "adidas Men's Yeezy Boost 350 V2 Grey/Borang/Dgsogr",
		Tags:                      `["New Product","Summer Product"]`,
		DisplayCoverMultimediaUrl: `["https://images-na.ssl-images-amazon.com/images/I/81WmojBxvbL._AC_UL1500_.jpg"]`,
		ProductGroupId:            "1356",
		CommentCount:              100,
		SoldCount:                 60,
		PublishTimestamp:          1623193487,
		Source:                    "self",
		UserRating:                0.25,
		SellerId:                  "43485",
		SellerLevel:               "1",
		SellerRating:              3.5,
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
		UserId:           "1457789",
		EventType:        "purchase",
		EventTimestamp:   1686883465,
		SceneName:        "product detail page",
		PageNumber:       2,
		Offset:           10,
		ProductId:        "632461",
		Platform:         "app",
		OsType:           "android",
		AppVersion:       "9.2.0",
		DeviceModel:      "huawei-mate30",
		OsVersion:        "10",
		Network:          "3g",
		Query:            "iPad",
		ParentProductId:  "441356",
		AttributionToken: "eyJpc3MiOiJuaW5naGFvLm5ldCIsImV4cCI6IjE0Mzg5NTU0NDUiLCJuYW1lIjoid2FuZ2hhbyIsImFkbWluIjp0cnVlfQ",
		TrafficSource:    "self",
		PurchaseCount:    20,
		PaidPrice:        12.23,
		Currency:         "USD",
		City:             "Kirkland",
		Country:          "USA",
		District:         "King County",
		Province:         "Texas",
		//CustomField:                   "custom",
	}
}

func mockPredictProduct() *protocol.Product {
	return &protocol.Product{
		ProductId: "632461",
	}
}

func mockPredictDevice() *protocol.Device {
	return &protocol.Device{
		Platform:    "app",
		OsType:      "android",
		AppVersion:  "9.2.0",
		DeviceModel: "huawei-mate30",
		DeviceBrand: "huawei",
		OsVersion:   "10",
		BrowserType: "chrome",
		UserAgent:   "Mozilla/5.0 (Linux; Android 10; TAS-AN00; HMSCore 5.3.0.312) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 HuaweiBrowser/11.0.8.303 Mobile Safari/537.36",
		Network:     "3g",
	}
}
