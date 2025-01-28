package createstack

import (
	"fmt"
	"slices"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func CreateStack(Stack types.DecodedStack) error {
	fmt.Println(Stack.ExcludedRegions)

	if slices.Contains(Stack.ExcludedRegions, Stack.Region) {
		fmt.Printf("Not Deploying Stack %s in %s\n", Stack.Name, Stack.Region)
		return nil
	}

	if slices.Contains(Stack.ExcludedEnvironments, Stack.Environment) {
		fmt.Printf("Not Deploying Stack %s in %s\n", Stack.Name, Stack.Environment)
		return nil
	}

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
