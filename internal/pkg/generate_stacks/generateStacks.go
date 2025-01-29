package generatestacks

import (
	"fmt"
	"slices"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GenerateStacks(Stacks types.Stacks, Regions types.Regions, Environments types.Environments) types.DecodedStacks {
	var DecodedStacks types.DecodedStacks

	for _, environment := range Environments.Environment {
		for _, region := range Regions.Region {
			for _, stack := range Stacks.Stack {
				if slices.Contains(stack.ExcludeRegions, region.Name) {
					continue
				}
				if slices.Contains(stack.ExcludeEnvironments, environment.Name) {
					continue
				}
				if !slices.Contains(stack.ExcludeRegions, region.Name) && !slices.Contains(stack.ExcludeEnvironments, environment.Name) {
					parentPath := fmt.Sprintf("./stacks/%s/%s", environment.Name, region.Name)
					DecodedStacks = processStack(stack, region.Name, environment.Name, parentPath, DecodedStacks, stack.ExcludeRegions, stack.ExcludeEnvironments)
				}
			}
		}
	}

	return DecodedStacks
}

// Helper function to recursively process a stack and its child stacks.
func processStack(
	stack types.Stack,
	region string,
	environment string,
	parentPath string,
	DecodedStacks types.DecodedStacks,
	ExcludeRegions []string,
	ExcludeEnvironments []string) types.DecodedStacks {
	Stack := types.DecodedStack{
		Name:                 stack.Name,
		Path:                 parentPath + "/" + stack.Name,
		ParentPath:           parentPath,
		Tags:                 stack.Tags,
		Description:          stack.Name,
		Region:               region,
		Environment:          environment,
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
