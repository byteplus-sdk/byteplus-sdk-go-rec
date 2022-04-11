package retail

import (
	"errors"
	"fmt"
	"strings"

	core "github.com/byteplus-sdk/byteplus-sdk-go-rec-core"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec-core/logs"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec-core/option"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/retail/protocol"
)

var (
	writeMsgFormat  = "Only can receive max to %d items in one write request"
	writeTooManyErr = errors.New(fmt.Sprintf(writeMsgFormat, maxWriteCount))
)

type clientImpl struct {
	httpClient *core.HTTPClient
	projectID  string
}

func (c *clientImpl) Release() {
	c.httpClient.Shutdown()
}

func checkPredictRequest(projectId string, modelId string) error {
	const (
		errMsgFormat      = "%s field can't be empty"
		errFieldProjectId = "projectId"
		errFieldModelId   = "modelId"
	)
	if projectId != "" && modelId != "" {
		return nil
	}
	emptyParams := make([]string, 0)
	if projectId == "" {
		emptyParams = append(emptyParams, errFieldProjectId)
	}
	if modelId == "" {
		emptyParams = append(emptyParams, errFieldModelId)
	}
	return errors.New(fmt.Sprintf(errMsgFormat, strings.Join(emptyParams, ",")))
}

func checkUploadDataRequest(request *protocol.WriteDataRequest) error {
	const (
		errMsgFormat      = "%s field can't' be empty"
		errFieldProjectId = "projectId"
		errFieldStage     = "stage"
	)
	if request.GetProjectId() != "" && request.GetStage() != "" {
		return nil
	}
	emptyParams := make([]string, 0)
	if request.GetProjectId() == "" {
		emptyParams = append(emptyParams, errFieldProjectId)
	}
	if request.GetStage() == "" {
		emptyParams = append(emptyParams, errFieldStage)
	}
	return errors.New(fmt.Sprintf(errMsgFormat, strings.Join(emptyParams, ",")))
}

func (c *clientImpl) doWrite(request *protocol.WriteDataRequest,
	path string, opts ...option.Option) (*protocol.WriteResponse, error) {
	if len(c.projectID) > 0 && len(request.ProjectId) == 0 {
		request.ProjectId = c.projectID
	}
	if err := checkUploadDataRequest(request); err != nil {
		return nil, err
	}
	if len(request.GetData()) > maxWriteCount {
		return nil, writeTooManyErr
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
	return c.doWrite(writeRequest, UserUri, opts...)
}

func (c *clientImpl) WriteProducts(writeRequest *protocol.WriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	return c.doWrite(writeRequest, ProductUri, opts...)
}

func (c *clientImpl) WriteUserEvents(writeRequest *protocol.WriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	return c.doWrite(writeRequest, UserEventUri, opts...)
}

func (c *clientImpl) WriteOthers(writeRequest *protocol.WriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	return c.doWrite(writeRequest, OthersUri, opts...)
}

func checkFinishDataRequest(request *protocol.FinishWriteDataRequest) error {
	const (
		errMsgFormat      = "%s field can't' be empty"
		errFieldProjectId = "projectId"
		errFieldStage     = "stage"
		errFieldTopic     = "topic"
	)
	if request.GetProjectId() != "" && request.GetStage() != "" && request.GetTopic() != "" {
		return nil
	}
	emptyParams := make([]string, 0)
	if request.GetProjectId() == "" {
		emptyParams = append(emptyParams, errFieldProjectId)
	}
	if request.GetStage() == "" {
		emptyParams = append(emptyParams, errFieldStage)
	}
	if request.GetTopic() == "" {
		emptyParams = append(emptyParams, errFieldTopic)
	}
	return errors.New(fmt.Sprintf(errMsgFormat, strings.Join(emptyParams, ",")))
}

func (c *clientImpl) doFinish(request *protocol.FinishWriteDataRequest,
	path string, opts ...option.Option) (*protocol.WriteResponse, error) {
	if len(c.projectID) > 0 && len(request.ProjectId) == 0 {
		request.ProjectId = c.projectID
	}
	if err := checkFinishDataRequest(request); err != nil {
		return nil, err
	}
	if len(request.GetDataDates()) > maxWriteCount {
		return nil, writeTooManyErr
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
	return c.doFinish(finishRequest, FinishUserUri, opts...)
}

func (c *clientImpl) FinishWriteProducts(finishRequest *protocol.FinishWriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	return c.doFinish(finishRequest, FinishProductUri, opts...)
}

func (c *clientImpl) FinishWriteUserEvents(finishRequest *protocol.FinishWriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	return c.doFinish(finishRequest, FinishUserEventUri, opts...)
}

func (c *clientImpl) FinishWriteOthers(finishRequest *protocol.FinishWriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	return c.doFinish(finishRequest, FinishOthersUri, opts...)
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
	err := c.httpClient.DoPBRequest(PredictUri,
		request, response, option.Conv2Options(opts...))
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
