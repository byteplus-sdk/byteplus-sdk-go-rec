## byteplus rec go sdk

#### Install the client library
```go
// add this to "go.mod"
require github.com/byteplus-sdk/byteplus-sdk-go-rec latest

//or
go get github.com/byteplus-sdk/byteplus-sdk-go-rec
```

#### Saas E-Commerce Example
```go
import (
	sdklog "github.com/byteplus-sdk/byteplus-sdk-go-rec-core/logs"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/region"
	sdk "github.com/byteplus-sdk/byteplus-sdk-go-rec/retail"
)

var client sdk.Client

func init() {
	sdklog.Level = sdklog.LevelDebug
	var err error
	client, err = sdk.NewClientBuilder().
		TenantID("***********"). // Required. The account id of byteplus.
		ProjectID("*************").
		Region(region.SG). // Required. The region of the server used to provide service.
        AuthAK("***********"). // Required. Access Key, used to generate request signature. Saas Standard project use this.
        AuthSK("***********"). // Required. Secure key, used to generate request signature. Saas Standard project use this.
        // AirAuthToken("***********"). // Required. The token of this project. Saas Premium project use this.
		//Schema("https"). // Optional.
		//Hosts([]string{"rec-api-sg1.recplusapi.com"}). // Optional.
		Build()
	if err != nil {
		panic(fmt.Sprintf("byteplus rec sdk init err:%s", err.Error()))
	}
}

func main() {
	client.WriteUsers()
	client.Predict()
}
```

#### Saas Content(Short-Video/Image/Doc) Example
```go
import (
	sdklog "github.com/byteplus-sdk/byteplus-sdk-go-rec-core/logs"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/region"
	sdk "github.com/byteplus-sdk/byteplus-sdk-go-rec/content"
)

var client sdk.Client

func init() {
	sdklog.Level = sdklog.LevelDebug
	var err error
	client, err = sdk.NewClientBuilder().
		TenantID("***********"). // Required. The account id of byteplus.
		ProjectID("*************").
		Region(region.SG). // Required. The region of the server used to provide service.
        AuthAK("***********"). // Required. Access Key, used to generate request signature. Saas Standard projects should use.
        AuthSK("***********"). // Required. Secure key, used to generate request signature. Saas Standard projects should use.
        // AirAuthToken("***********"). // Required. The token of this project. Saas Premium projects should use.
		//Schema("https"). // Optional.
		//Hosts([]string{"rec-api-sg1.recplusapi.com"}). // Optional.
		Build()
	if err != nil {
		panic(fmt.Sprintf("byteplus rec sdk init err:%s", err.Error()))
	}
}

func main() {
	client.WriteUsers()
	client.Predict()
}
```

#### How to run example project
Take the E-commerce industry as an example:
* clone the project.
* enter the example directory.
* fill necessary parameters.
* build the binary executable file.
* run executable file.

```shell
git clone https://github.com/byteplus-sdk/byteplus-sdk-go-rec.git
cd byteplus-sdk-go-rec
go mod tidy
cd retail/example
# fill in projectID, modelID, tenantID, AK, SK and other parameters.
go build
./example
```

#### For more details
* [Saas E-Commerce Code Sample](https://docs.byteplus.com/recommend/docs/code-samples)
* [Saas E-Commerce API Reference](https://docs.byteplus.com/recommend/reference/byteplussaasservice_writusers-2)
* [Saas Content Code Sample](https://docs.byteplus.com/recommend/docs/content-code-samples)
* [Saas Content API Reference](https://docs.byteplus.com/recommend/reference/byteplussaasservice_writusers)