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
		UserId:                 "1457789",
		Gender:                 "male",
		Age:                    "23",
		Tags:                   "[\"new user\",\"low purchasing power\",\"bargain seeker\"]",
		ActivationChannel:      "AppStore",
		MembershipLevel:        "silver",
		RegistrationTimestamp:  1623593487,
		LocationCity:           "Kirkland",
		LocationCountry:        "USA",
		LocationDistrictOrArea: "King County",
		LocationPostcode:       "98033",
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
		ProductId:                     "632461",
		Category:                      "[{\"category_depth\":1,\"category_nodes\":[{\"id_or_name\":\"Shoes\"}]},{\"category_depth\":2,\"category_nodes\":[{\"id_or_name\":\"Men's Shoes\"}]}]",
		Brands:                        "Adidas",
		PriceCurrentPrice:             49900,
		PriceOriginPrice:              69900,
		IsRecommendable:               true,
		Title:                         "adidas Men's Yeezy Boost 350 V2 Grey/Borang/Dgsogr",
		QualityScore:                  4.4,
		Tags:                          "[\"New Product\",\"Summer Product\"]",
		DisplayDetailPageDisplayTags:  "[\"FreeShipping\",\"Return in 7 days without any reasons\"]",
		DisplayCoverMultimediaUrl:     "https://images-na.ssl-images-amazon.com/images/I/81WmojBxvbL._AC_UL1500_.jpg",
		DisplayListingPageDisplayTags: "[\"best seller\",\"hot sales\"]",
		DisplayListingPageDisplayType: "image",
		ProductSpecProductGroupId:     "1356",
		ProductSpecCommentCount:       100,
		ProductSpecPublishTimestamp:   1623193487,
		ProductSpecSource:             "self",
		ProductSpecUserRating:         0.25,
		SellerId:                      "43485",
		SellerSellerLevel:             "1",
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
		UserId:               "1457789",
		EventType:            "purchase",
		EventTimestamp:       1623681888,
		Scene:                "product detail page",
		ScenePageNumber:      2,
		SceneOffset:          10,
		ProductId:            "632461",
		DevicePlatform:       "app",
		DeviceOsType:         "android",
		DeviceAppVersion:     "9.2.0",
		DeviceDeviceModel:    "huawei-mate30",
		DeviceDeviceBrand:    "huawei",
		DeviceOsVersion:      "10",
		DeviceBrowserType:    "chrome",
		DeviceUserAgent:      "Mozilla/5.0 (Linux; Android 10; TAS-AN00; HMSCore 5.3.0.312) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 HuaweiBrowser/11.0.8.303 Mobile Safari/537.36",
		DeviceNetwork:        "3g",
		ContextQuery:         "iPad",
		ContextRootProductId: "441356",
		AttributionToken:     "eyJpc3MiOiJuaW5naGFvLm5ldCIsImV4cCI6IjE0Mzg5NTU0NDUiLCJuYW1lIjoid2FuZ2hhbyIsImFkbWluIjp0cnVlfQ",
		RecInfo:              "CiRiMjYyYjM1YS0xOTk1LTQ5YmMtOGNkNS1mZTVmYTczN2FkNDASJAobcmVjZW50X2hvdF9jbGlja3NfcmV0cmlldmVyFQAAAAAYDxoKCgNjdHIdog58PBoKCgNjdnIdANK2OCIHMjcyNTgwMg==",
		TrafficSource:        "self",
		PurchaseCount:        20,
		// 10s
		DetailPageStayTime: 10,
		//CustomField:                   "custom",
	}
}

func mockPredictProduct() *protocol.Product {
	category1Node1 := &protocol.Product_Category_CategoryNode{
		IdOrName: "Shoes",
	}
	category1 := &protocol.Product_Category{
		CategoryDepth: 1,
		CategoryNodes: []*protocol.Product_Category_CategoryNode{category1Node1},
	}
	category2Node1 := &protocol.Product_Category_CategoryNode{
		IdOrName: "Men's Shoes",
	}
	category2 := &protocol.Product_Category{
		CategoryDepth: 2,
		CategoryNodes: []*protocol.Product_Category_CategoryNode{category2Node1},
	}

	brand1 := &protocol.Product_Brand{
		BrandDepth: 1,
		IdOrName:   "Adidas",
	}
	brand2 := &protocol.Product_Brand{
		BrandDepth: 2,
		IdOrName:   "Yeezy",
	}

	price := &protocol.Product_Price{
		CurrentPrice: 49900,
		OriginPrice:  69900,
	}

	display := &protocol.Product_Display{
		DetailPageDisplayTags:  []string{"FreeShipping", "Return in 7 days without any reasons"},
		ListingPageDisplayTags: []string{"best seller", "hot sales"},
		ListingPageDisplayType: "image",
		CoverMultimediaUrl:     "https://images-na.ssl-images-amazon.com/images/I/81WmojBxvbL._AC_UL1500_.jpg",
	}

	spec := &protocol.Product_ProductSpec{
		ProductGroupId:   "1356",
		UserRating:       0.25,
		CommentCount:     100,
		Source:           "self",
		PublishTimestamp: 1623193487,
	}

	seller := &protocol.Product_Seller{
		Id:           "43485",
		SellerLevel:  "1",
		SellerRating: 3.5,
	}

	return &protocol.Product{
		ProductId:       "632461",
		Categories:      []*protocol.Product_Category{category1, category2},
		Brands:          []*protocol.Product_Brand{brand1, brand2},
		Price:           price,
		IsRecommendable: true,
		Title:           "adidas Men's Yeezy Boost 350 V2 Grey/Borang/Dgsogr",
		QualityScore:    4.4,
		Tags:            []string{"New Product", "Summer Product"},
		Display:         display,
		ProductSpec:     spec,
		Seller:          seller,
		Extra:           map[string]string{"count": "20", "color": "white"},
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
