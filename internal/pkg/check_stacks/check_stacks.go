package checkstacks

import "os"

func CheckStacks(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
