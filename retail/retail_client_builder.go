package retail

import (
	"errors"

	core "github.com/byteplus-sdk/byteplus-sdk-go-rec-core"
	_ "github.com/byteplus-sdk/byteplus-sdk-go-rec/region"
)

type clientBuilder struct {
	tenantID     string
	projectID    string
	airAuthToken string
	authAK       string
	authSK       string
	schema       string
	hosts        []string
	region       core.IRegion
}

func NewClientBuilder() *clientBuilder {
	return &clientBuilder{}
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

func (receiver *clientBuilder) Region(region core.IRegion) *clientBuilder {
	receiver.region = region
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
		Region(receiver.region).
		AuthService(byteplusAuthService).Build()
	if err != nil {
		return nil, err
	}
	client := &clientImpl{
		httpClient: httpClient,
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
