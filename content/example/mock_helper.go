package main

import (
	"strconv"

	"github.com/byteplus-sdk/byteplus-sdk-go-rec/content/protocol"
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
		Tags:                  `["new user","low purchasing power","bargain seeker"]`, // Json Array, use json.Marshal()
		Language:              "English",
		SubscriberType:        "free",
		MembershipLevel:       "silver",
		RegistrationTimestamp: 1659958007,
		Country:               "USA",
		Province:              "Texas",
		City:                  "Kirkland",
		District:              "King County",
		//CustomField:                   "custom",
	}
}

func mockContents(count int) []*DemoContent {
	contents := make([]*DemoContent, count)
	for i := 0; i < count; i++ {
		content := mockProduct()
		content.ContentId += strconv.Itoa(i)
		contents[i] = content
	}
	return contents
}

func mockProduct() *DemoContent {
	return &DemoContent{
		ContentId:        "632461",
		IsRecommendable:  1,
		Categories:       `[{"category_depth":1,"category_nodes":[{"id_or_name":"Movie"}]},{"category_depth":2,"category_nodes":[{"id_or_name":"Comedy"}]}]`, // Json Array, use json.Marshal().
		ContentType:      "video",
		VideoDuration:    120000,
		ContentTitle:     "Green Book Movie Explanation",
		Description:      "A brief summary of the main content of the Green Book movie",
		ContentOwnerId:   "1457789",
		CollectionId:     "1342",
		Tags:             `["New","Trending"]`,                                                               // Json Array, use json.Marshal().
		ImageUrls:        `["https://images-na.ssl-images-amazon.com/images/I/81WmojBxvbL._AC_UL1500_.jpg"]`, // Json Array, use json.Marshal().
		VideoUrls:        `["https://test_video.mov"]`,                                                       // Json Array, use json.Marshal().
		UserRating:       4.9,
		CurrentPrice:     1300.12,
		OriginalPrice:    1600.12,
		PublishTimestamp: 1660035734,
		IsPaidContent:    true,
		Language:         "English",
		LinkedProductId:  `["632462","632463"]`, // Json Array, use json.Marshal().
		Source:           "self",
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
		UserId:           "1457787",
		EventType:        "purchase",
		EventTimestamp:   1686883465,
		ContentId:        "632461",
		TrafficSource:    "byteplus",
		AttributionToken: "eyJpc3MiOiJuaW5naGFvLm5ldCIsImV4cCI6IjE0Mzg5NTU0NDUiLCJuYW1lIjoid2FuZ2hhbyIsImFkbWluIjp0cnVlfQ",
		SceneName:        "Home page",
		PageNumber:       2,
		Offset:           10,
		StayDuration:     150000,
		ParentContentId:  "632431",
		ContentOwnerId:   "1457789",
		Query:            "comedy",
		Platform:         "app",
		OsType:           "ios",
		AppVersion:       "1.0.1",
		DeviceModel:      "iPhone10",
		OsVersion:        "14.4.2",
		Network:          "4g",
		Country:          "USA",
		Province:         "Texas",
		City:             "Kirkland",
		District:         "King County",
		PurchaseCount:    20,
		PaidPrice:        12.23,
		Currency:         "USD",

		//CustomField:                   "custom",
	}
}

func mockPredictContent() *protocol.Content {
	return &protocol.Content{
		ContentId: "632461",
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
