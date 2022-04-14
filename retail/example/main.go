package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	core "github.com/byteplus-sdk/byteplus-sdk-go-rec-core"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec-core/logs"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec-core/option"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/region"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/retail"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/retail/protocol"
	"github.com/google/uuid"
)

const (
	DefaultWriteTimeout = 800 * time.Millisecond

	DefaultPredictTimeout = 800 * time.Millisecond

	DefaultAckImpressionsTimeout = 800 * time.Millisecond

	DefaultFinishTimeout = 800 * time.Millisecond

	DefaultRetryTimes = 2
)

var (
	client retail.Client
)

const (
	// A unique identity assigned by Bytedance.
	projectID = "***********"

	// Unique id for this model.
	// The saas model id that can be used to get rec results from predict api, which is need to fill in URL.
	modelID = "***********"
)

func init() {
	logs.Level = logs.LevelDebug
	var err error
	client, err = retail.NewClientBuilder().
		AccountID("***********"). // Required. The account id of byteplus.
		ProjectID(projectID).
		Region(region.SG).     // Required. The region of the server used to provide service.
		AuthAK("***********"). // Required. Access Key, used to generate request signature.
		AuthSK("***********"). // Required. Secure key, used to generate request signature.
		//Schema("https"). // Optional.
		//Hosts([]string{"rec-api-sg1.recplusapi.com"}). // Optional.
		Build()
	if err != nil {
		panic(fmt.Sprintf("byteplus rec sdk init err:%s", err.Error()))
	}
}

func main() {
	// Write real-time user data
	writeUsersExample()

	// Finish write real-time user data
	//finishWriteUsersExample()

	// Write real-time product data
	writeProductsExample()

	// Finish write real-time product data
	//finishWriteProductsExample()

	// Write real-time user event data
	writeUserEventsExample()

	// Finish write real-time user event data
	//finishWriteUserEventsExample()

	// Write self defined topic data
	//writeOthersExample()

	// Finish write self defined topic data
	//finishWriteOthersExample()

	// Get recommendation results
	recommendExample()

	// Pause for 5 seconds until the asynchronous import task completes
	time.Sleep(5 * time.Second)
	client.Release()
	os.Exit(0)
}

func writeUsersExample() {
	// The "WriteXXX" api can transfer max to 2000 items at one request
	request := buildWriteUsersRequest(1)
	opts := defaultOptions(DefaultWriteTimeout)
	response, err := client.WriteUsers(request, opts...)
	if err != nil {
		logs.Error("write user occur err, msg:%v", err)
		return
	}
	if core.IsUploadSuccess(response.GetStatus().GetCode()) {
		logs.Info("write user success")
		return
	}
	logs.Error("write user find failure info, msg:%s errItems:%+v",
		response.GetStatus(), response.GetErrors())
}

func buildWriteUsersRequest(count int) *protocol.WriteDataRequest {
	users := mockUsers(count)
	marshalUsers := make([]string, 0, len(users))
	for _, user := range users {
		marshalUser, _ := json.Marshal(user)
		marshalUsers = append(marshalUsers, string(marshalUser))
	}
	return &protocol.WriteDataRequest{
		Stage: retail.StageTrial,
		Data:  marshalUsers,
		Extra: map[string]string{"extra_info": "extra"},
	}
}

func finishWriteUsersExample() {
	request := buildFinishUserRequest()
	opts := defaultOptions(DefaultFinishTimeout)
	response, err := client.FinishWriteUsers(request, opts...)
	if err != nil {
		logs.Error("run finish occur error, msg:%v", err)
		return
	}
	if core.IsUploadSuccess(response.GetStatus().GetCode()) {
		logs.Info("finish write user data")
		return
	}
	logs.Error("fail to finish write user data, msg:%s errItems:%+v",
		response.GetStatus(), response.GetErrors())
}

func buildFinishUserRequest() *protocol.FinishWriteDataRequest {
	return &protocol.FinishWriteDataRequest{
		Stage: retail.StageIncrementalDaily,
	}
}

func writeProductsExample() {
	// The "WriteXXX" api can transfer max to 2000 items at one request
	request := buildWriteProductsRequest(1)
	opts := defaultOptions(DefaultWriteTimeout)
	response, err := client.WriteProducts(request, opts...)
	if err != nil {
		logs.Error("write product occur err, msg:%v", err)
		return
	}
	if core.IsUploadSuccess(response.GetStatus().GetCode()) {
		logs.Info("write product success")
		return
	}
	logs.Error("write product find failure info, msg:%s errItems:%+v",
		response.GetStatus(), response.GetErrors())
}

func buildWriteProductsRequest(count int) *protocol.WriteDataRequest {
	products := mockProducts(count)
	marshalProducts := make([]string, 0, len(products))
	for _, product := range products {
		marshalProduct, _ := json.Marshal(product)
		marshalProducts = append(marshalProducts, string(marshalProduct))
	}
	return &protocol.WriteDataRequest{
		Stage: retail.StageTrial,
		Data:  marshalProducts,
		Extra: map[string]string{"extra_info": "extra"},
	}
}

func finishWriteProductsExample() {
	// The "FinishXXX" api can mark max to 100 dates at one request
	request := buildFinishProductRequest()
	opts := defaultOptions(DefaultFinishTimeout)
	response, err := client.FinishWriteProducts(request, opts...)
	if err != nil {
		logs.Error("run finish occur error, msg:%v", err)
		return
	}
	if core.IsUploadSuccess(response.GetStatus().GetCode()) {
		logs.Info("finish write product data")
		return
	}
	logs.Error("fail to finish write product data, msg:%s errItems:%+v",
		response.GetStatus(), response.GetErrors())
}

func buildFinishProductRequest() *protocol.FinishWriteDataRequest {
	return &protocol.FinishWriteDataRequest{
		Stage: retail.StageIncrementalDaily,
	}
}

func writeUserEventsExample() {
	// The "WriteXXX" api can transfer max to 2000 items at one request
	request := buildWriteUserEventsRequest(1)
	opts := defaultOptions(DefaultWriteTimeout)
	response, err := client.WriteUserEvents(request, opts...)
	if err != nil {
		logs.Error("write user event occur err, msg:%v", err)
		return
	}
	if core.IsUploadSuccess(response.GetStatus().GetCode()) {
		logs.Info("write user event success")
		return
	}
	logs.Error("write user event find failure info, msg:%s errItems:%+v",
		response.GetStatus(), response.GetErrors())
}

func buildWriteUserEventsRequest(count int) *protocol.WriteDataRequest {
	userEvents := mockUserEvents(count)
	marshalUserEvents := make([]string, 0, len(userEvents))
	for _, userEvent := range userEvents {
		marshalUserEvent, _ := json.Marshal(userEvent)
		marshalUserEvents = append(marshalUserEvents, string(marshalUserEvent))
	}
	return &protocol.WriteDataRequest{
		Stage: retail.StageTrial,
		Data:  marshalUserEvents,
		Extra: map[string]string{"extra_info": "extra"},
	}
}

func finishWriteUserEventsExample() {
	// The "FinishXXX" api can mark max to 100 dates at one request
	request := buildFinishUserEventRequest()
	opts := defaultOptions(DefaultFinishTimeout)
	response, err := client.FinishWriteUserEvents(request, opts...)
	if err != nil {
		logs.Error("run finish occur error, msg:%v", err)
		return
	}
	if core.IsUploadSuccess(response.GetStatus().GetCode()) {
		logs.Info("finish user event product data")
		return
	}
	logs.Error("fail to finish write user event data, msg:%s errItems:%+v",
		response.GetStatus(), response.GetErrors())
}

func buildFinishUserEventRequest() *protocol.FinishWriteDataRequest {
	// dates should be passed when finishing others
	dates := []*protocol.Date{
		{
			Year:  2022,
			Month: 2,
			Day:   1,
		}}
	return &protocol.FinishWriteDataRequest{
		Stage:     retail.StageIncrementalDaily,
		DataDates: dates,
	}
}

func writeOthersExample() {
	// The "WriteXXX" api can transfer max to 2000 items at one request
	// The `topic` is datatype, which specify the type of data users are going to write.
	// It is temporarily set to "video", the specific value depends on your need.
	topic := "video"
	request := buildWriteOthersRequest(topic)
	opts := defaultOptions(DefaultWriteTimeout)
	response, err := client.WriteOthers(request, opts...)
	if err != nil {
		logs.Error("write others occur err, msg:%v", err)
		return
	}
	if core.IsUploadSuccess(response.GetStatus().GetCode()) {
		logs.Info("write others success")
		return
	}
	logs.Error("write others find failure info, msg:%s errItems:%+v",
		response.GetStatus(), response.GetErrors())
}

func buildWriteOthersRequest(topic string) *protocol.WriteDataRequest {
	data := map[string]interface{}{
		"field1": 1,
		"field2": "value2",
	}
	marshalData, _ := json.Marshal(data)
	datas := []string{string(marshalData)}
	return &protocol.WriteDataRequest{
		Stage: retail.StageTrial,
		Topic: topic,
		Data:  datas,
		Extra: map[string]string{"extra_info": "extra"},
	}
}

func finishWriteOthersExample() {
	// The "FinishXXX" api can mark max to 100 dates at one request
	// The `topic` is datatype, which specify the type of data users are going to finish writing.
	// It is temporarily set to "video", the specific value depends on your need.
	topic := "video"
	request := buildFinishOthersRequest(topic)
	opts := defaultOptions(DefaultFinishTimeout)
	response, err := client.FinishWriteOthers(request, opts...)
	if err != nil {
		logs.Error("run finish occur error, msg:%v", err)
		return
	}
	if core.IsUploadSuccess(response.GetStatus().GetCode()) {
		logs.Info("finish writing data")
		return
	}
	logs.Error("fail to finish write data, msg:%s errItems:%+v",
		response.GetStatus(), response.GetErrors())
}

func buildFinishOthersRequest(topic string) *protocol.FinishWriteDataRequest {
	// dates should be passed when finishing others
	dates := []*protocol.Date{
		{
			Year:  2022,
			Month: 1,
			Day:   1,
		}}
	return &protocol.FinishWriteDataRequest{
		Stage:     retail.StageIncrementalDaily,
		DataDates: dates,
		Topic:     topic,
	}
}

func recommendExample() {
	predictRequest := buildPredictRequest()
	predictOpts := defaultOptions(DefaultPredictTimeout)
	// The "home" is scene name, which provided by ByteDance, usually is "home"
	response, err := client.Predict(predictRequest, predictOpts...)
	if err != nil {
		logs.Error("predict occur error, msg:%v", err)
		return
	}
	if !core.IsSuccess(response.GetStatus().GetCode()) {
		logs.Error("predict find failure info, msg:%s", response.GetStatus())
		return
	}
	logs.Info("predict success")
	// The items, which is eventually shown to user,
	// should send back to Bytedance for deduplication
	alteredProducts := recommendWithPredictResult(response.GetValue())
	ackRequest := buildAckRequest(response.GetRequestId(), predictRequest, alteredProducts)
	ackOpts := defaultOptions(DefaultAckImpressionsTimeout)
	// async ack the actual impressions after this recommendation
	core.AsyncExecute(func() {
		_ = core.DoWithRetry(DefaultRetryTimes, func() error {
			_, err := client.AckServerImpressions(ackRequest, ackOpts...)
			return err
		})
	})
}

func buildPredictRequest() *protocol.PredictRequest {
	scene := &protocol.Scene{
		Offset: 10,
	}
	rootProduct := mockPredictProduct()
	device := mockPredictDevice()
	context := &protocol.PredictRequest_Context{
		RootProduct:         rootProduct,
		Device:              device,
		CandidateProductIds: []string{"632462", "632463"},
	}
	return &protocol.PredictRequest{
		ModelId: modelID,
		UserId:  "1457789",
		Size:    20,
		Scene:   scene,
		Context: context,
		// Extra:     map[string]string{"extra_info": "extra"},
	}
}

func recommendWithPredictResult(
	predictResult *protocol.PredictResult) []*protocol.AckServerImpressionsRequest_AlteredProduct {
	// You can handle recommend results here,
	// such as filter, insert other items, sort again, etc.
	// The list of goods finally displayed to user and the filtered goods
	// should be sent back to bytedance for deduplication
	return conv2AlteredProducts(predictResult.GetResponseProducts())
}

func conv2AlteredProducts(
	products []*protocol.PredictResult_ResponseProduct) []*protocol.AckServerImpressionsRequest_AlteredProduct {
	if len(products) == 0 {
		return nil
	}
	alteredProducts := make([]*protocol.AckServerImpressionsRequest_AlteredProduct, len(products))
	for i, product := range products {
		alteredProducts[i] = &protocol.AckServerImpressionsRequest_AlteredProduct{
			AlteredReason: "kept",
			ProductId:     product.GetProductId(),
			Rank:          int32(i + 1),
		}
	}
	return alteredProducts
}

func buildAckRequest(predictRequestId string, predictRequest *protocol.PredictRequest,
	alteredProducts []*protocol.AckServerImpressionsRequest_AlteredProduct) *protocol.AckServerImpressionsRequest {

	return &protocol.AckServerImpressionsRequest{
		ModelId:          predictRequest.GetModelId(),
		PredictRequestId: predictRequestId,
		UserId:           predictRequest.GetUserId(),
		Scene:            predictRequest.GetScene(),
		AlteredProducts:  alteredProducts,
	}
}

func defaultOptions(timeout time.Duration) []option.Option {
	// All options are optional
	opts := []option.Option{
		option.WithRequestID(uuid.NewString()),
		option.WithTimeout(timeout),
		// Optional. Add a header to a custom http header collection.
		//option.WithHTTPHeader("key", "value"),
		// Optional. Add a query to a custom query collection.
		//option.WithHTTPQuery("key", "value"),
		// Optional. It is expected that the server will process the data for the maximum time.
		// If the processing time exceeds this time, the server will return the result immediately,
		// regardless of whether there is any remaining data that has not been processed.
		//option.WithServerTimeout(5000 * time.Millisecond),
	}
	return opts
}
