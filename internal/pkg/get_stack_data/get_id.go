package getstackdata

import "os/exec"

func GetStackID(path string) string {
	command := "terramate experimental get-config-value -C " + path + " 'terramate.stack.id'"
	cmd := exec.Command("sh", "-c", command)

	// run the command
	output, _ := cmd.CombinedOutput()

	return string(output)
}
