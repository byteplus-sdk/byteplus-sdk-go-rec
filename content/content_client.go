package content

import (
	"github.com/byteplus-sdk/byteplus-sdk-go-rec-core/option"
	"github.com/byteplus-sdk/byteplus-sdk-go-rec/content/protocol"
)

type Client interface {
	// WriteUsers
	//
	// Writes at most 2000 users data at a time. Exceeding 2000 in a request results in
	// a rejection.Each element of dataList array is a json serialized string of data.
	// One can use this to upload new data, or update existing data.
	WriteUsers(writeRequest *protocol.WriteDataRequest, opts ...option.Option) (*protocol.WriteResponse, error)

	// FinishWriteUsers
	//
	// Recording that user data has been written. Mark at most 100 dates at a time
	// No need to finish real-time data, the system will automatically finish when entering the next day
	FinishWriteUsers(request *protocol.FinishWriteDataRequest, opts ...option.Option) (*protocol.WriteResponse, error)

	// WriteContents
	//
	// Writes at most 2000 contents data at a time. Exceeding 2000 in a request results in
	// a rejection.Each element of dataList array is a json serialized string of data.
	// One can use this to upload new data, or update existing data.
	WriteContents(writeRequest *protocol.WriteDataRequest, opts ...option.Option) (*protocol.WriteResponse, error)

	// FinishWriteContents
	//
	// Recording that content data has been written. Mark at most 100 dates at a time
	// No need to finish real-time data, the system will automatically finish when entering the next day
	FinishWriteContents(request *protocol.FinishWriteDataRequest, opts ...option.Option) (*protocol.WriteResponse, error)

	// WriteUserEvents
	//
	// Writes at most 2000 user events data at a time. Exceeding 2000 in a request results in
	// a rejection.Each element of dataList array is a json serialized string of data.
	// One can use this to upload new data, or update existing data (by providing all the fields,
	// some data type not support update, e.g. user event).
	WriteUserEvents(writeRequest *protocol.WriteDataRequest, opts ...option.Option) (*protocol.WriteResponse, error)

	// FinishWriteUserEvents
	//
	// Recording that user event data has been written. Mark at most 100 dates at a time
	// In general, you need to pass the date list in FinishWriteDataRequest. While if the date list is empty,
	// the data of the previous day will be finished by default.
	// No need to finish real-time data, the system will automatically finish when entering the next day
	FinishWriteUserEvents(request *protocol.FinishWriteDataRequest, opts ...option.Option) (*protocol.WriteResponse, error)

	// WriteOthers
	//
	// Writes at most 2000 data at a time, the topic of these data is set by users
	// One can use this to upload new data, or update existing data.
	WriteOthers(request *protocol.WriteDataRequest, opts ...option.Option) (*protocol.WriteResponse, error)

	// FinishWriteOthers
	//
	// Recording that some data has been written, the topic of these data is set by users.
	// Mark at most 100 dates at a time
	// No need to finish real-time data, the system will automatically finish when entering the next day
	FinishWriteOthers(request *protocol.FinishWriteDataRequest, opts ...option.Option) (*protocol.WriteResponse, error)

	// Predict
	//
	// Gets the list of contents (ranked).
	// The updated user data will take effect in 24 hours.
	// The updated content data will take effect in 30 mins.
	// Depending on how (realtime or batch) the UserEvents are sent back, it will
	// be fed into the models and take effect after that.
	Predict(request *protocol.PredictRequest, opts ...option.Option) (*protocol.PredictResponse, error)

	// AckServerImpressions
	//
	// Sends back the actual content list shown to the users based on the
	// customized changes from `PredictResponse`.
	// example: our Predict call returns the list of items [1, 2, 3, 4].
	// Your custom logic have decided that content 3 has been sold out and
	// content 10 needs to be inserted before 2 based on some promotion rules,
	// the AckServerImpressionsRequest content items should looks like
	// [
	//   {id:1, altered_reason: "kept", rank:1},
	//   {id:10, altered_reason: "inserted", rank:2},
	//   {id:2, altered_reason: "kept", rank:3},
	//   {id:4, altered_reason: "kept", rank:4},
	//   {id:3, altered_reason: "filtered", rank:0},
	// ].
	AckServerImpressions(request *protocol.AckServerImpressionsRequest,
		opts ...option.Option) (*protocol.AckServerImpressionsResponse, error)

	//Release resource used by client
	Release()
}
