package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

//EnvironmentReaderWrapper Wraps the os dependency that reads environment variables
type EnvironmentReaderWrapper struct{}

//errEnvVarEmpty Error to indicate that an environment variable is missing
var errEnvVarEmpty = errors.New("getenv: environment variable empty")

//getenvStr Get string value from environment internal method
func (EnvironmentReaderWrapper) getenvStr(key string, shouldPanicOnError bool) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		if shouldPanicOnError {
			panic(fmt.Sprintf("Missing environment variable: %s", key))
		}

		return v, errEnvVarEmpty
	}
	return v, nil
}

//GetenvStr Get string value from environment
func (reader EnvironmentReaderWrapper) GetenvStr(key string, defaultVal string) (string, error) {
	s, err := reader.getenvStr(key, false)
	if err != nil {
		s, err = defaultVal, nil
	}

	return s, err
}

//GetenvInt Get integer value from environment
func (reader EnvironmentReaderWrapper) getenvInt(key string, shouldPanicOnError bool) (int, error) {
	s, err := reader.getenvStr(key, shouldPanicOnError)

	if err != nil {
		return 0, err
	}

	v, err := strconv.Atoi(s)

	if err != nil {
		if shouldPanicOnError {
			panic(fmt.Sprintf("Bad Value for environment variable: [%s] => %s", key, s))
		}
		return 0, err
	}
	return v, nil
}

//GetenvInt Get integer value from environment
func (reader EnvironmentReaderWrapper) GetenvInt(key string, defaultVal int) (int, error) {
	v, err := reader.getenvInt(key, false)
	if err != nil {
		v, err = defaultVal, nil
	}

	return v, err
}

//GetenvStrStrict Get string value from environment with panic on failure
func (reader EnvironmentReaderWrapper) GetenvStrStrict(key string) string {
	v, _ := reader.getenvStr(key, true)
	return v
}

//GetenvIntStrict Get integer value from environment with panic on failure
func (reader EnvironmentReaderWrapper) GetenvIntStrict(key string) int {
	v, _ := reader.getenvInt(key, true)
	return v
}

//GetenvBytes Get bytes value from environment
func (reader *EnvironmentReaderWrapper) GetenvBytes(key string, shouldPanicOnError bool) ([]byte, error) {
	s, err := reader.getenvStr(key, shouldPanicOnError)
	if err != nil {
		return nil, err
	}

	return []byte(s), nil
}
