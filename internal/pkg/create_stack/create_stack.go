package createstack

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/tmtf-stacker/stacker/internal/pkg/types"
	"os/exec"
	"strings"
)

func CreateStack(stack types.DecodedStack) error {
	// Join tags into a comma-separated string
	tags := strings.Join(stack.Tags, ",")
	log.Debugf("Prepared tags: %s", tags)

	// Build the command string
	command := fmt.Sprintf(
		`terramate create --tags "%s" --description "%s" %s`,
		tags, stack.Description, stack.Path,
	)
	log.Debugf("Executing command: %s", command)

	// Prepare the command
	cmd := exec.Command("sh", "-c", command)

	// Run the command and capture its combined output
	output, err := cmd.CombinedOutput()

	// Log both error and combined output if something goes wrong
	if err != nil {
		log.Errorf("Failed to create stack. Error: %v", err)
		log.Errorf("Command output: %s", string(output))
		return err
	}

	// Log both the success message and the command output
	log.Infof("Successfully created stack at path '%s'", stack.Path)
	log.Debugf("Command output: %s", string(output))

	return nil
}
