package getstackdata

import (
	"os/exec"
	"strings"
)

func GetStackName(path string) string {
	command := "terramate experimental get-config-value -C " + path + " 'terramate.stack.name'"
	cmd := exec.Command("sh", "-c", command)

	// run the command
	output, _ := cmd.CombinedOutput()

	cleanedOutput := strings.TrimSpace(string(output))

	return cleanedOutput
}
