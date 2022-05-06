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

	// StageIncremental In this stage, data will be used to update model
	StageIncremental = "incremental_sync_streaming"

	// StageIncrementalRealtime In this stage, data will be used to realtime update model
	// Please use `STAGE_INCREMENTAL` instead `STAGE_INCREMENTAL_REALTIME` in most cases
	StageIncrementalRealtime = "incremental_sync_streaming"

	// StageIncrementalDaily In this stage, data will be used to daily update model
	// Please use `STAGE_INCREMENTAL` instead `STAGE_INCREMENTAL_DAILY` in most cases
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
