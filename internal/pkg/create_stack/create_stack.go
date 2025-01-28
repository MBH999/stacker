package createstack

import (
	"fmt"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func CreateStack(Stack types.DecodedStack) error {
	fmt.Println(Stack.Path)
	return nil
	// tags := strings.Join(Stack.Tags, ",")
	//
	// command := "terramate create --tags " + tags + " " + "--description \"" + Stack.Description + "\" " + Stack.Path
	// cmd := exec.Command("sh", "-c", command)
	//
	// _, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Errorf("error: %v", err.Error())
	// 	return err
	// }
	//
	// log.Infof("Created stack: %v", Stack.Path)
	// return err
}
