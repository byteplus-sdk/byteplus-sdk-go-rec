package retail

const (
	maxWriteCount  = 2000
	maxFinishCount = 100
)

const (
	// StageTrial In this stage, data will be only used to test
	StageTrial = "trial"

	// StageProduction In this stage, data will be used to train model
	StageProduction = "production"

	// StageIncrementalRealtime In this stage, data will be used to realtime update model
	StageIncrementalRealtime = "incremental_sync_streaming"

	// StageIncrementalDaily In this stage, data will be used to daily update model
	StageIncrementalDaily = "incremental_sync_daily"

	// TopicUser Topic if write users and finish write users
	TopicUser = "user"

	// TopicProduct Topic if write products and finish write products
	TopicProduct = "goods"

	// TopicUserEvent Topic if write user events and finish write user events
	TopicUserEvent = "behavior"

	// UserUri in user topic, url path is end with WriteUsers
	UserUri = "/RetailSaaS/WriteUsers"

	// FinishUserUri The URL format of finish information
	FinishUserUri = "/RetailSaaS/FinishWriteUsers"

	ProductUri = "/RetailSaaS/WriteProducts"

	FinishProductUri = "/RetailSaaS/FinishWriteProducts"

	UserEventUri = "/RetailSaaS/WriteUserEvents"

	FinishUserEventUri = "/RetailSaaS/FinishWriteUserEvents"

	OthersUri = "/RetailSaaS/WriteOthers"

	FinishOthersUri = "/RetailSaaS/FinishWriteOthers"

	PredictUri = "/RetailSaaS/Predict"
)
