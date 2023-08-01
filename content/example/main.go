package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	core "github.com/byteplus-sdk/byteplus-sdk-go-rec-core"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec-core/logs"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec-core/option"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/content"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/content/protocol"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/region"
	"github.com/google/uuid"
)

const (
	DefaultWriteTimeout = 800 * time.Millisecond

	DefaultFinishTimeout = 800 * time.Millisecond

	DefaultPredictTimeout = 800 * time.Millisecond

	DefaultAckImpressionsTimeout = 800 * time.Millisecond

	DefaultRetryTimes = 2
)

var (
	client content.Client
)

const (
	// A unique identity assigned by Bytedance.
	projectID = "***********"

	// Unique id for this model.
	// The saas model id that can be used to get rec results from predict api, which is need to fill in URL.
	modelID = "***********"
)

func init() {
	//// Customize the caller config, the parameters in Example are the parameters currently used by default,
	//// you can customize them according to your own needs.
	//callerConfig := &core.CallerConfig{
	//	// fasthttp.Client MaxIdleConnDuration param.
	//	KeepAliveDuration: 60 * time.Second,
	//	// Only takes effect when contentClient.keepAlive(true), heartbeat packet sending interval.
	//	KeepAlivePingInterval: 45 * time.Second,
	//	// Default is 512. The maximum number of connections for a single host, if the maximum number of connections is exceeded, an error will be reported.
	//	MaxConnections: 512,
	//}
	//
	//// Metrics configuration, when Metrics and Metrics Log are turned on,
	//// the metrics and logs at runtime will be collected and sent to the byteplus server.
	//// During debugging, byteplus can help customers troubleshoot problems.
	//metricsConfig := &metrics.Config{
	//	// enable metrics, default is false.
	//	EnableMetrics: true,
	//	// enable metrics log, default is false.
	//	EnableMetricsLog: true,
	//	// The time interval for reporting metrics to the byteplus server, the default is 15s.
	//	// When the QPS is high, the value of the reporting interval can be reduced to prevent
	//	// loss of metrics.
	//	// The longest should not exceed 30s, otherwise it will cause the loss of metrics accuracy.
	//	ReportInterval: 15 * time.Second,
	//}

	logs.Level = logs.LevelDebug
	var err error
	client, err = content.NewClientBuilder().
		AccountID("***********"). // Required. The account id of byteplus.
		ProjectID(projectID).
		Region(region.SG).     // Required. The region of the server used to provide service.
		AuthAK("***********"). // Required. Access Key, used to generate request signature. Saas Standard projects should use.
		AuthSK("***********"). // Required. Secure key, used to generate request signature. Saas Standard projects should use.
		// AirAuthToken("***********"). // Required. The token of this project. Saas Premium projects should use.
		//KeepAlive(true). // Optional.
		//CallerConfig(callerConfig). // Optional.
		//MetricsConfig(metricsConfig). // Optional.
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

	// Write real-time content data
	writeContentsExample()

	// Finish write real-time content data
	//finishWriteContentsExample()

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
		Stage: content.StageIncremental,
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
		Stage: content.StageIncremental,
	}
}

func writeContentsExample() {
	// The "WriteXXX" api can transfer max to 2000 items at one request
	request := buildWriteContentsRequest(1)
	opts := defaultOptions(DefaultWriteTimeout)
	response, err := client.WriteContents(request, opts...)
	if err != nil {
		logs.Error("write content occur err, msg:%v", err)
		return
	}
	if core.IsUploadSuccess(response.GetStatus().GetCode()) {
		logs.Info("write content success")
		return
	}
	logs.Error("write content find failure info, msg:%s errItems:%+v",
		response.GetStatus(), response.GetErrors())
}

func buildWriteContentsRequest(count int) *protocol.WriteDataRequest {
	contents := mockContents(count)
	marshalContents := make([]string, 0, len(contents))
	for _, content := range contents {
		marshalContent, _ := json.Marshal(content)
		marshalContents = append(marshalContents, string(marshalContent))
	}
	return &protocol.WriteDataRequest{
		Stage: content.StageIncremental,
		Data:  marshalContents,
		Extra: map[string]string{"extra_info": "extra"},
	}
}

func finishWriteContentsExample() {
	// The "FinishXXX" api can mark max to 100 dates at one request
	request := buildFinishContentRequest()
	opts := defaultOptions(DefaultFinishTimeout)
	response, err := client.FinishWriteContents(request, opts...)
	if err != nil {
		logs.Error("run finish occur error, msg:%v", err)
		return
	}
	if core.IsUploadSuccess(response.GetStatus().GetCode()) {
		logs.Info("finish write content data")
		return
	}
	logs.Error("fail to finish write content data, msg:%s errItems:%+v",
		response.GetStatus(), response.GetErrors())
}

func buildFinishContentRequest() *protocol.FinishWriteDataRequest {
	return &protocol.FinishWriteDataRequest{
		Stage: content.StageIncremental,
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
		Stage: content.StageIncremental,
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
		logs.Info("finish write user event data")
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
		Stage:     content.StageIncremental,
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
		Stage: content.StageIncremental,
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
		Stage:     content.StageIncremental,
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
	alteredContents := recommendWithPredictResult(response.GetContentValue())
	ackRequest := buildAckRequest(response.GetRequestId(), predictRequest, alteredContents)
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
	rootContent := mockPredictContent()
	device := mockPredictDevice()
	context := &protocol.PredictRequest_Context{
		RootContent:       rootContent,
		Device:            device,
		CandidateContents: []*protocol.Content{mockPredictContent()},
	}
	return &protocol.PredictRequest{
		ModelId:        modelID,
		UserId:         "1457789",
		Size:           20,
		Scene:          scene,
		ContentContext: context,
		// Extra:     map[string]string{"extra_info": "extra"},
	}
}

func recommendWithPredictResult(
	predictResult *protocol.PredictResult) []*protocol.AckServerImpressionsRequest_AlteredContent {
	// You can handle recommend results here,
	// such as filter, insert other items, sort again, etc.
	// The list of goods finally displayed to user and the filtered goods
	// should be sent back to bytedance for deduplication
	return conv2AlteredContents(predictResult.GetResponseContents())
}

func conv2AlteredContents(
	contents []*protocol.PredictResult_ResponseContent) []*protocol.AckServerImpressionsRequest_AlteredContent {
	if len(contents) == 0 {
		return nil
	}
	alteredContents := make([]*protocol.AckServerImpressionsRequest_AlteredContent, len(contents))
	for i, aContent := range contents {
		alteredContents[i] = &protocol.AckServerImpressionsRequest_AlteredContent{
			AlteredReason: "kept",
			ContentId:     aContent.GetContentId(),
			Rank:          int32(i + 1),
		}
	}
	return alteredContents
}

func buildAckRequest(predictRequestId string, predictRequest *protocol.PredictRequest,
	alteredContents []*protocol.AckServerImpressionsRequest_AlteredContent) *protocol.AckServerImpressionsRequest {

	return &protocol.AckServerImpressionsRequest{
		ModelId:          predictRequest.GetModelId(),
		PredictRequestId: predictRequestId,
		UserId:           predictRequest.GetUserId(),
		// If it is the recommendation result from byteplus, traffic_source is byteplus,
		// if it is the customer's own recommendation result, traffic_source is self.
		TrafficSource:   "byteplus",
		Scene:           predictRequest.GetScene(),
		AlteredContents: alteredContents,
		// Extra:            map[string]string{"ip": "127.0.0.1"},
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
