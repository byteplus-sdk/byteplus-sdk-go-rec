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

const (
	dataInitOptionCount    = 1
	predictInitOptionCount = 1
)

type clientImpl struct {
	httpClient *core.HTTPClient
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

func checkUploadDataRequest(projectId string, stage string) error {
	const (
		errMsgFormat      = "%s field can't' be empty"
		errFieldProjectId = "projectId"
		errFieldStage     = "stage"
	)
	if projectId != "" && stage != "" {
		return nil
	}
	emptyParams := make([]string, 0)
	if projectId == "" {
		emptyParams = append(emptyParams, errFieldProjectId)
	}
	if stage == "" {
		emptyParams = append(emptyParams, errFieldStage)
	}
	return errors.New(fmt.Sprintf(errMsgFormat, strings.Join(emptyParams, ",")))
}

func (c *clientImpl) doWrite(request *protocol.WriteDataRequest,
	url string, opts ...option.Option) (*protocol.WriteResponse, error) {
	if err := checkUploadDataRequest(request.ProjectId, request.Stage); err != nil {
		return nil, err
	}
	if len(request.GetData()) > maxWriteCount {
		return nil, writeTooManyErr
	}
	if len(opts) == 0 {
		opts = make([]option.Option, 0, dataInitOptionCount)
	}
	response := &protocol.WriteResponse{}
	opts = addSaasFlag(opts)
	err := c.httpClient.DoPbRequest(url, request, response, option.Conv2Options(opts...))
	if err != nil {
		return nil, err
	}
	logs.Debug("[WriteData] rsp:\n%s\n", response)
	return response, nil
}

func addSaasFlag(opts []option.Option) []option.Option {
	return append(opts, withSaasHeader())
}

func withSaasHeader() option.Option {
	const (
		HTTPHeaderServerFrom = "Server-From"
		SaasFlag             = "saas"
	)
	return func(opt *option.Options) {
		if len(opt.Headers) == 0 {
			opt.Headers = map[string]string{HTTPHeaderServerFrom: SaasFlag}
			return
		}
		opt.Headers[HTTPHeaderServerFrom] = SaasFlag
	}
}

func (c *clientImpl) WriteUsers(writeRequest *protocol.WriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	return c.doWrite(writeRequest, "/RetailSaaS/WriteUsers", opts...)
}

func (c *clientImpl) WriteProducts(writeRequest *protocol.WriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	return c.doWrite(writeRequest, "/RetailSaaS/WriteProducts", opts...)
}

func (c *clientImpl) WriteUserEvents(writeRequest *protocol.WriteDataRequest,
	opts ...option.Option) (*protocol.WriteResponse, error) {
	return c.doWrite(writeRequest, "/RetailSaaS/WriteUserEvents", opts...)
}

func (c *clientImpl) Predict(request *protocol.PredictRequest,
	opts ...option.Option) (*protocol.PredictResponse, error) {
	if err := checkPredictRequest(request.ProjectId, request.ModelId); err != nil {
		return nil, err
	}
	if len(opts) == 0 {
		opts = make([]option.Option, 0, predictInitOptionCount)
	}
	response := &protocol.PredictResponse{}
	opts = addSaasFlag(opts)
	err := c.httpClient.DoPbRequest("/RetailSaaS/Predict",
		request, response, option.Conv2Options(opts...))
	if err != nil {
		return nil, err
	}
	logs.Debug("[Predict] rsp:\n%s\n", response)
	return response, nil
}

func (c *clientImpl) AckServerImpressions(request *protocol.AckServerImpressionsRequest,
	opts ...option.Option) (*protocol.AckServerImpressionsResponse, error) {
	if err := checkPredictRequest(request.ProjectId, request.ModelId); err != nil {
		return nil, err
	}
	if len(opts) == 0 {
		opts = make([]option.Option, 0, predictInitOptionCount)
	}
	response := &protocol.AckServerImpressionsResponse{}
	opts = addSaasFlag(opts)
	err := c.httpClient.DoPbRequest("/RetailSaaS/AckServerImpressions",
		request, response, option.Conv2Options(opts...))
	if err != nil {
		return nil, err
	}
	logs.Debug("[AckImpressions] rsp:\n%s\n", response)
	return response, nil
}
