package generatestacks

import (
	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GenerateStacks(Stacks types.Stacks, RegionStacks types.DecodedStacks, EnvironmentStacks types.DecodedStacks) types.DecodedStacks {
	var DecodedStacks types.DecodedStacks

	for _, environment := range EnvironmentStacks.DecodedStack {
		for _, region := range RegionStacks.DecodedStack {
			for _, stack := range Stacks.Stack {
				DecodedStacks = processStack(stack, region, environment, region.Path, DecodedStacks, stack.ExcludeRegions, stack.ExcludeEnvironments)
			}
		}
	}
	return DecodedStacks
}

// Helper function to recursively process a stack and its child stacks.
func processStack(
	stack types.Stack,
	region types.DecodedStack,
	environment types.DecodedStack,
	parentPath string,
	DecodedStacks types.DecodedStacks,
	ExcludeRegions []string,
	ExcludeEnvironments []string) types.DecodedStacks {
	Stack := types.DecodedStack{
		Name:                 stack.Name,
		Path:                 parentPath + "/" + stack.Name,
		ParentPath:           parentPath,
		Tags:                 append(region.Tags, stack.Tags...),
		Description:          stack.Name,
		Region:               region.Name,
		Environment:          environment.Name,
		ExcludedEnvironments: ExcludeEnvironments,
		ExcludedRegions:      ExcludeRegions,
	}

	// if stack.Description != "" {
	// 	DecodedStack.Description = stack.Description
	// 	fmt.Println(DecodedStack.Description)
	// } else {
	// 	DecodedStack.Description = stack.Name
	// }

	DecodedStacks.DecodedStack = append(DecodedStacks.DecodedStack, Stack)

	if stack.ChildStack != nil {
		for _, childStack := range stack.ChildStack {
			DecodedStacks = processStack(childStack, region, environment, Stack.Path, DecodedStacks, ExcludeRegions, ExcludeEnvironments)
		}
	}

	return DecodedStacks
}
