package createstack

import (
	"os/exec"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func CreateStack(Stack types.DecodedStack) int {
	tags := strings.Join(Stack.Tags, ",")

	command := "terramate create --tags " + tags + " " + Stack.Path
	// fmt.Println(Stack.Tags)
	cmd := exec.Command("sh", "-c", command)

	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("error: %v", err.Error())
	}

	log.Infof("Created stack: %v", Stack.Path)
	return 0
}
