package generatestacks

import (
	// "fmt"

	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GenerateEnvironments(Environments types.Environments) types.DecodedStacks {
	var Stacks types.DecodedStacks

	for _, environment := range Environments.Environment {
		Stack := types.DecodedStack{
			Name:                 environment.Name,
			Path:                 "./stacks/" + environment.Name,
			ParentPath:           "./stacks/",
			Tags:                 []string{environment.Name},
			Description:          environment.Name,
			Region:               "N/A",
			Environment:          environment.Name,
			DeployAsStack:        environment.DeployAsStack,
			ExcludedEnvironments: []string{},
			ExcludedRegions:      []string{},
		}
		Stacks.DecodedStack = append(Stacks.DecodedStack, Stack)
	}
	return Stacks
}
