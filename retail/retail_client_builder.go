package retail

import (
	"errors"

	core "github.com/byteplus-sdk/byteplus-sdk-go-rec-core"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec-core/metrics"
	_ "github.com/byteplus-sdk/byteplus-sdk-go-rec/region"
)

type clientBuilder struct {
	tenantID              string
	projectID             string
	airAuthToken          string
	authAK                string
	authSK                string
	schema                string
	mainHost              string
	hosts                 []string
	region                core.IRegion
	callerConfig          *core.CallerConfig
	hostAvailablerFactory core.HostAvailablerFactory
	metricsConfig         *metrics.Config
	keepAlive             bool
}

func NewClientBuilder() *clientBuilder {
	return &clientBuilder{}
}

func (receiver *clientBuilder) AccountID(accountID string) *clientBuilder {
	receiver.tenantID = accountID
	return receiver
}

func (receiver *clientBuilder) TenantID(tenantID string) *clientBuilder {
	receiver.tenantID = tenantID
	return receiver
}

func (receiver *clientBuilder) ProjectID(projectID string) *clientBuilder {
	receiver.projectID = projectID
	return receiver
}

func (receiver *clientBuilder) AirAuthToken(airAuthToken string) *clientBuilder {
	receiver.airAuthToken = airAuthToken
	return receiver
}

func (receiver *clientBuilder) AuthAK(authAK string) *clientBuilder {
	receiver.authAK = authAK
	return receiver
}

func (receiver *clientBuilder) AuthSK(authSK string) *clientBuilder {
	receiver.authSK = authSK
	return receiver
}

func (receiver *clientBuilder) Schema(schema string) *clientBuilder {
	receiver.schema = schema
	return receiver
}

func (receiver *clientBuilder) Hosts(hosts []string) *clientBuilder {
	receiver.hosts = hosts
	return receiver
}

func (receiver *clientBuilder) MainHost(host string) *clientBuilder {
	receiver.mainHost = host
	return receiver
}

func (receiver *clientBuilder) Region(region core.IRegion) *clientBuilder {
	receiver.region = region
	return receiver
}

func (receiver *clientBuilder) CallerConfig(callerConfig *core.CallerConfig) *clientBuilder {
	receiver.callerConfig = callerConfig
	return receiver
}

func (receiver *clientBuilder) HostAvailablerFactory(hostAvailablerFactory core.HostAvailablerFactory) *clientBuilder {
	receiver.hostAvailablerFactory = hostAvailablerFactory
	return receiver
}

func (receiver *clientBuilder) MetricsConfig(metricsConfig *metrics.Config) *clientBuilder {
	receiver.metricsConfig = metricsConfig
	return receiver
}

func (receiver *clientBuilder) KeepAlive(keepAlive bool) *clientBuilder {
	receiver.keepAlive = keepAlive
	return receiver
}

const (
	byteplusAuthService = "byteplus_recommend"
)

func (receiver *clientBuilder) Build() (Client, error) {
	err := receiver.checkRequiredField()
	if err != nil {
		return nil, err
	}
	httpClient, err := core.NewHTTPClientBuilder().
		TenantID(receiver.tenantID).
		ProjectID(receiver.projectID).
		UseAirAuth(receiver.isUseAirAuth()).
		AirAuthToken(receiver.airAuthToken).
		AuthAK(receiver.authAK).
		AuthSK(receiver.authSK).
		Schema(receiver.schema).
		Hosts(receiver.hosts).
		MainHost(receiver.mainHost).
		Region(receiver.region).
		AuthService(byteplusAuthService).
		CallerConfig(receiver.callerConfig).
		HostAvailablerFactory(receiver.hostAvailablerFactory).
		MetricsCfg(receiver.metricsConfig).
		KeepAlive(receiver.keepAlive).Build()
	if err != nil {
		return nil, err
	}
	client := &clientImpl{
		httpClient: httpClient,
		projectID:  receiver.projectID,
	}
	return client, nil
}

func (receiver *clientBuilder) checkRequiredField() error {
	if len(receiver.projectID) < 0 {
		return errors.New("project id is empty")
	}
	return nil
}

func (receiver *clientBuilder) isUseAirAuth() bool {
	return len(receiver.authAK) == 0 && len(receiver.authSK) == 0 && len(receiver.airAuthToken) > 0
}
