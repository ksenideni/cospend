package util

import (
	"os"
)

func Getenv(key string) (fallback string) {
	var value string
	if key == "" {
		return ""
	} else {
		value = os.Getenv(key)
		if len(value) == 0 {
			return ""
		}
		return value
	}
}
