package retail

const (
	maxWriteCount = 2000
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
)
