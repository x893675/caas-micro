package util

import "os"

func GetEnvironment(key, value string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return value
	}
	return val
}
