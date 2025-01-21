package checkstacks

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func CheckTags(Stack types.DecodedStack) {
	// tags := strings.Join(Stack.Tags, ",")
	tags, _ := json.Marshal(Stack.Tags)

	command := "terramate experimental get-config-value -C " + Stack.Path + " 'terramate.stack.tags'"
	cmd := exec.Command("sh", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
	}
	// fmt.Println(string(output))
	// fmt.Println(string(tags))

	if string(output) == string(tags) {
		fmt.Println("true")
	}
}
