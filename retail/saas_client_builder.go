package retail

import (
	core "github.com/byteplus-sdk/byteplus-sdk-go-rec-core"
	_ "github.com/byteplus-sdk/byteplus-sdk-go-rec/region"
)

type clientBuilder struct {
	tenantID   string
	token      string
	schema     string
	hosts      []string
	hostHeader string
	headers    map[string]string
	region     core.Region
	ak         string
	sk         string
}

func NewClientBuilder() *clientBuilder {
	return &clientBuilder{}
}

func (receiver *clientBuilder) AK(ak string) *clientBuilder {
	receiver.ak = ak
	return receiver
}

func (receiver *clientBuilder) SK(sk string) *clientBuilder {
	receiver.sk = sk
	return receiver
}

func (receiver *clientBuilder) TenantID(tenantID string) *clientBuilder {
	receiver.tenantID = tenantID
	return receiver
}

func (receiver *clientBuilder) Token(token string) *clientBuilder {
	receiver.token = token
	return receiver
}

func (receiver *clientBuilder) Schema(schema string) *clientBuilder {
	receiver.schema = schema
	return receiver
}

func (receiver *clientBuilder) HostHeader(hostHeader string) *clientBuilder {
	receiver.hostHeader = hostHeader
	return receiver
}

func (receiver *clientBuilder) Hosts(hosts []string) *clientBuilder {
	receiver.hosts = hosts
	return receiver
}

func (receiver *clientBuilder) Headers(headers map[string]string) *clientBuilder {
	receiver.headers = headers
	return receiver
}

func (receiver *clientBuilder) Region(region core.Region) *clientBuilder {
	receiver.region = region
	return receiver
}

const (
	saasTenant      = "saas"
	volcAuthService = "air"
)

func (receiver *clientBuilder) Build() (Client, error) {
	httpClient, err := core.NewHTTPClientBuilder().
		AK(receiver.ak).
		SK(receiver.sk).
		TenantID(receiver.tenantID).
		Token(receiver.token).
		Schema(receiver.schema).
		HostHeader(receiver.hostHeader).
		Hosts(receiver.hosts).
		Headers(receiver.headers).
		Region(receiver.region).
		AuthService(volcAuthService).Build()
	if err != nil {
		return nil, err
	}
	client := &clientImpl{
		httpClient: httpClient,
	}
	return client, nil
}
