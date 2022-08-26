package util_os

import (
	"os"
	"strconv"
)

func GetEnvironmentVariable[T any](name string, convertor func(value string) (T, error), fallback T) T {
	value, ok := os.LookupEnv(name)

	if ok {
		result, err := convertor(value)

		if err == nil {
			return result
		}
	}

	return fallback
}

func ToInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}
