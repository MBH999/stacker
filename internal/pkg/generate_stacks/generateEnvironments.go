package generatestacks

import (
	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GenerateEnvironments(Environment types.Environments) types.DecodedStacks {
	var Stacks types.DecodedStacks

	for _, environment := range Environment.Config {
		Stack := types.DecodedStack{
			Name:                 environment,
			Path:                 "./stacks/" + environment,
			ParentPath:           "./stacks/",
			Tags:                 []string{environment},
			Description:          environment,
			Region:               "N/A",
			Environment:          environment,
			ExcludedEnvironments: []string{},
			ExcludedRegions:      []string{},
		}
		Stacks.DecodedStack = append(Stacks.DecodedStack, Stack)
	}
	return Stacks
}
