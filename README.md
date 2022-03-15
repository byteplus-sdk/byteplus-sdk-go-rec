## byteplus rec go sdk

#### Install the client library
```go
// add this to "go.mod"
require github.com/byteplus-sdk/byteplus-sdk-go-rec latest

//or
go get github.com/byteplus-sdk/byteplus-sdk-go-rec
```

#### How to run example
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