package retail

import (
	"errors"
	"fmt"

	core "github.com/byteplus-sdk/byteplus-sdk-go-rec-core"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec-core/logs"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec-core/option"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/retail/protocol"
)

var (
	writeMsgFormat   = "Only can receive max to %d items in one request"
	writeTooManyErr  = errors.New(fmt.Sprintf(writeMsgFormat, maxWriteCount))
	finishTooManyErr = errors.New(fmt.Sprintf(writeMsgFormat, maxFinishCount))
)

type clientImpl struct {
	httpClient *core.HTTPClient
	projectID  string
}

func (c *clientImpl) Release() {
	c.httpClient.Shutdown()
}

func checkWriteDataRequest(request *protocol.WriteDataRequest) error {
	if request.GetProjectId() == "" {
		return errors.New("project id is empty")
	}
	if request.GetStage() == "" {
		return errors.New("stage is empty")
	}
	if len(request.GetData()) > maxWriteCount {
		return writeTooManyErr
	}
	return nil
}

func (c *clientImpl) doWrite(request *protocol.WriteDataRequest,
	path string, opts ...option.Option) (*protocol.WriteResponse, error) {
	if len(c.projectID) > 0 && len(request.ProjectId) == 0 {
		request.ProjectId = c.projectID
	}
	if err := checkWriteDataRequest(request); err != nil {
		return nil, err
	}
	response := &protocol.WriteResponse{}
	err := c.httpClient.DoPBRequest(path, request, response, option.Conv2Options(opts...))
	if err != nil {
		return nil, err
	}
	logs.Debug("[WriteData] rsp:\n%s\n", response)
	return response, nil
}

func (c *clientImpl) WriteUsers(writeRequest *protocol.WriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	writeRequest.Topic = TopicUser
	return c.doWrite(writeRequest, UserUri, opts...)
}

func (c *clientImpl) WriteProducts(writeRequest *protocol.WriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	writeRequest.Topic = TopicProduct
	return c.doWrite(writeRequest, ProductUri, opts...)
}

func (c *clientImpl) WriteUserEvents(writeRequest *protocol.WriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	writeRequest.Topic = TopicUserEvent
	return c.doWrite(writeRequest, UserEventUri, opts...)
}

func (c *clientImpl) WriteOthers(writeRequest *protocol.WriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	return c.doWrite(writeRequest, OthersUri, opts...)
}

func checkFinishDataRequest(request *protocol.FinishWriteDataRequest) error {
	if request.GetProjectId() == "" {
		return errors.New("project id is empty")
	}
	if request.GetStage() == "" {
		return errors.New("stage is empty")
	}
	if request.GetTopic() == "" {
		return errors.New("topic is empty")
	}
	if len(request.GetDataDates()) > maxFinishCount {
		return finishTooManyErr
	}
	return nil
}

func (c *clientImpl) doFinish(request *protocol.FinishWriteDataRequest,
	path string, opts ...option.Option) (*protocol.WriteResponse, error) {
	if len(c.projectID) > 0 && len(request.ProjectId) == 0 {
		request.ProjectId = c.projectID
	}
	if err := checkFinishDataRequest(request); err != nil {
		return nil, err
	}
	response := &protocol.WriteResponse{}
	err := c.httpClient.DoPBRequest(path, request, response, option.Conv2Options(opts...))
	if err != nil {
		return nil, err
	}
	logs.Debug("[WriteData] rsp:\n%s\n", response)
	return response, nil
}

func (c *clientImpl) FinishWriteUsers(finishRequest *protocol.FinishWriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	finishRequest.Topic = TopicUser
	return c.doFinish(finishRequest, FinishUserUri, opts...)
}

func (c *clientImpl) FinishWriteProducts(finishRequest *protocol.FinishWriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	finishRequest.Topic = TopicProduct
	return c.doFinish(finishRequest, FinishProductUri, opts...)
}

func (c *clientImpl) FinishWriteUserEvents(finishRequest *protocol.FinishWriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	finishRequest.Topic = TopicUserEvent
	return c.doFinish(finishRequest, FinishUserEventUri, opts...)
}

func (c *clientImpl) FinishWriteOthers(finishRequest *protocol.FinishWriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	return c.doFinish(finishRequest, FinishOthersUri, opts...)
}

func checkPredictRequest(projectId string, modelId string) error {
	if projectId == "" {
		return errors.New("project id is empty")
	}
	if modelId == "" {
		return errors.New("model id is empty")
	}
	return nil
}

func (c *clientImpl) Predict(request *protocol.PredictRequest,
	opts ...option.Option) (*protocol.PredictResponse, error) {
	if len(c.projectID) > 0 && len(request.ProjectId) == 0 {
		request.ProjectId = c.projectID
	}
	if err := checkPredictRequest(request.ProjectId, request.ModelId); err != nil {
		return nil, err
	}
	response := &protocol.PredictResponse{}
	err := c.httpClient.DoPBRequest(PredictUri, request, response, option.Conv2Options(opts...))
	if err != nil {
		return nil, err
	}
	logs.Debug("[Predict] rsp:\n%s\n", response)
	return response, nil
}

func (c *clientImpl) AckServerImpressions(request *protocol.AckServerImpressionsRequest,
	opts ...option.Option) (*protocol.AckServerImpressionsResponse, error) {
	if len(c.projectID) > 0 && len(request.ProjectId) == 0 {
		request.ProjectId = c.projectID
	}
	if err := checkPredictRequest(request.ProjectId, request.ModelId); err != nil {
		return nil, err
	}
	response := &protocol.AckServerImpressionsResponse{}
	err := c.httpClient.DoPBRequest("/RetailSaaS/AckServerImpressions",
		request, response, option.Conv2Options(opts...))
	if err != nil {
		return nil, err
	}
	logs.Debug("[AckImpressions] rsp:\n%s\n", response)
	return response, nil
}
