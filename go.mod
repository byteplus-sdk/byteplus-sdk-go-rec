module github.com/byteplus-sdk/byteplus-sdk-go-rec

go 1.16

//replace github.com/byteplus-sdk/byteplus-sdk-go-rec-core => ../byteplus-sdk-go-rec-core

require (
	github.com/byteplus-sdk/byteplus-sdk-go-rec-core v0.1.8
	github.com/google/uuid v1.3.0
	google.golang.org/protobuf v1.27.1
)
