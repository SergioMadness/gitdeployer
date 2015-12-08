package helpers

import (
	"os"
)

func IsFileExists(path string) bool {
	result := false
	if _, err := os.Stat(path); err == nil {
		result = true
	}
	return result
}
