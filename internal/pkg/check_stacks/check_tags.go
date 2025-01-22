package checkstacks

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"slices"
	"strings"

	"github.com/tmtf-stacker/stacker/internal/pkg/helper"
	"github.com/tmtf-stacker/stacker/internal/pkg/types"
	updatestacks "github.com/tmtf-stacker/stacker/internal/pkg/update_stacks"
)

func CheckTags(Stack types.DecodedStack) {
	// take stack tags and convert to string.
	tags, _ := json.Marshal(Stack.Tags)
	tagsStr := string(tags)

	// create terramate command to grab stack tags
	command := "terramate experimental get-config-value -C " + Stack.Path + " 'terramate.stack.tags'"
	cmd := exec.Command("sh", "-c", command)

	// run the command
	output, _ := cmd.CombinedOutput()

	// convert output to string
	outputStr := string(output)

	sanitizeOutput := helper.RemoveChars(outputStr, "[] \n")
	sanitizeTags := helper.RemoveChars(tagsStr, "[] \n")

	outputArr := strings.Split(sanitizeOutput, ",")
	slices.Sort(outputArr)
	tagsArr := strings.Split(sanitizeTags, ",")
	slices.Sort(tagsArr)

	if !slices.Equal(outputArr, tagsArr) {
		updatestacks.UpdateTags(Stack.Path, tagsStr)
		fmt.Println("Updating Tags!")
	} else {
		fmt.Println("Tags Match!")
	}
}
