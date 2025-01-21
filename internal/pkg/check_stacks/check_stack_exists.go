package checkstacks

import "os"

func CheckStacks(path string) bool {
	stackPath := path + "/stack.tm.hcl"
	if _, err := os.Stat(stackPath); err == nil {
		return true
	}
	return false
}
