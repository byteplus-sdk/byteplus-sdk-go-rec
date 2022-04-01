## byteplus rec go sdk

#### Install the client library
```go
// add this to "go.mod"
require github.com/byteplus-sdk/byteplus-sdk-go-rec latest

//or
go get github.com/byteplus-sdk/byteplus-sdk-go-rec
```

#### Example
```go
import (
	sdklog "github.com/byteplus-sdk/byteplus-sdk-go-rec-core/logs"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/region"
	sdk "github.com/byteplus-sdk/byteplus-sdk-go-rec/retail"
)

var client sdk.Client

func init() {
	sdklog.Level = logs.LevelDebug
	var err error
	client, err = sdk.NewClientBuilder().
		TenantID("***********"). // Required. The account id of byteplus.
		ProjectID("*************").
		Region(region.SG). // Required. The region of the server used to provide service.
		AuthAK("***********"). // Required. Access Key, used to generate request signature.
		AuthSK("***********=="). // Required. Secure key, used to generate request signature.
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
* [code sample](https://docs.byteplus.com/recommend/docs/code-samplesaas)