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
		Region(region.SG).       // Required. The region of the server used to provide service.
		AuthAK("***********").   // Required. Access Key, used to generate request signature.
		AuthSK("***********=="). // Required. Secure key, used to generate request signature.
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

	// Write real-time product data˚
	writeProductsExample()

	// Write real-time user event data
	writeUserEventsExample()

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
		ProjectId: projectID,
		Stage:     retail.StageTrial,
		Data:      marshalUsers,
		Extra:     map[string]string{"extra_info": "extra"},
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
		ProjectId: projectID,
		Stage:     retail.StageTrial,
		Data:      marshalProducts,
		Extra:     map[string]string{"extra_info": "extra"},
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
		ProjectId: projectID,
		Stage:     retail.StageTrial,
		Data:      marshalUserEvents,
		Extra:     map[string]string{"extra_info": "extra"},
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
		ProjectId: projectID,
		ModelId:   modelID,
		UserId:    "1457789",
		Size:      20,
		Scene:     scene,
		Context:   context,
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
		ProjectId:        predictRequest.GetProjectId(),
		ModelId:          predictRequest.GetModelId(),
		PredictRequestId: predictRequestId,
		UserId:           predictRequest.GetUserId(),
		Scene:            predictRequest.GetScene(),
		AlteredProducts:  alteredProducts,
	}
}

func defaultOptions(timeout time.Duration) []option.Option {
	// All options are optional
	// var customerHeaders map[string]string
	// var customerQueries map[string]string
	opts := []option.Option{
		option.WithRequestID(uuid.NewString()),
		option.WithTimeout(timeout),
		// Optional. Add a set of customer headers to the request, which will be overwritten by multiple calls.
		//option.WithHeaders(customerHeaders),
		// Optional. Add a set of customer queries to the request, which will be overwritten by multiple calls.
		//option.WithQueries(customerQueries),
		// Optional. Add a header to an existing custom header collection.
		//option.WithHeader("key", "value"),
		// Optional. Add a query to an existing custom query collection.
		//option.WithQuery("key", "value"),
		// Optional. It is expected that the server will process the data for the maximum time.
		// If the processing time exceeds this time, the server will return the result immediately,
		// regardless of whether there is any remaining data that has not been processed.
		//option.WithServerTimeout(5000 * time.Millisecond),
	}
	return opts
}
