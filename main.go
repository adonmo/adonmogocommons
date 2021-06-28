package main

import (
	"time"

	"github.com/adonmo/adonmogocommons/env"
	"github.com/adonmo/adonmogocommons/logger"
)

func main() {
	envVarReader := env.EnvironmentReaderWrapper{}
	appLogLevelStr := envVarReader.GetenvStrStrict("APP_LOG_LEVEL")
	logger.SetupLogger(appLogLevelStr)

	logger.DebugF("Printing this message at %v", time.Now())
	logger.InfoF("Printing this message at %v", time.Now())
	logger.WarnF("Printing this message at %v", time.Now())
	logger.ErrorF("Printing this message at %v", time.Now())
}
