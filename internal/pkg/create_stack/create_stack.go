package createstack

import (
	"fmt"
	"os/exec"

	"github.com/charmbracelet/log"
)

func CreateStacks(stackPaths []string) {
	for _, path := range stackPaths {
		fmt.Println(path)
		command := "terramate create " + path
		cmd := exec.Command("sh", "-c", command)

		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Errorf("error: %v", err)
		}
	}
}
