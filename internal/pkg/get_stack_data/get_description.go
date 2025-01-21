package getstackdata

import (
	"os/exec"
	"strings"
)

func GetStackDescription(path string) string {
	command := "terramate experimental get-config-value -C " + path + " 'terramate.stack.description'"
	cmd := exec.Command("sh", "-c", command)

	// run the command
	output, _ := cmd.CombinedOutput()

	cleanedOutput := strings.TrimSpace(string(output))

	return cleanedOutput
}
