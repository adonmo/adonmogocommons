package main

import (
	"fmt"
	"time"

	"github.com/adonmo/adonmogocommons/env"
	"github.com/adonmo/adonmogocommons/logger"
	"github.com/adonmo/adonmogocommons/set"
)

func main() {
	envVarReader := env.EnvironmentReaderWrapper{}
	appLogLevelStr := envVarReader.GetenvStrStrict("APP_LOG_LEVEL")
	logger.SetupLogger(appLogLevelStr)

	logger.DebugF("Printing this message at %v", time.Now())
	logger.InfoF("Printing this message at %v", time.Now())
	logger.WarnF("Printing this message at %v", time.Now())
	logger.ErrorF("Printing this message at %v", time.Now())

	sEt := set.NewSet([]interface{}{"hey", "hello"})
	setInt := set.NewSetInt([]int{1, 2, 3, 4, 4})
	setString := set.NewSetString([]string{"hello", "world", "world"})

	logger.DebugF("length of set with interfaces %v", sEt.Len())
	logger.DebugF("length of set with ints %v", setInt.Len())
	logger.DebugF("length of set with strings %v", setString.Len())

	logger.DebugF("values of set with ints before mutation %v", setInt.Values())
	logger.DebugF("values of set with string before mutation %v", setString.Values())

	setInt.Map(func(i int) int {
		return i * 2
	})

	setString.Map(func(s string) string {
		return fmt.Sprintf("%s_mutated", s)
	})

	logger.DebugF("values of set with ints after mutation %v", setInt.Values())
	logger.DebugF("values of set with string after mutation %v", setString.Values())
}
