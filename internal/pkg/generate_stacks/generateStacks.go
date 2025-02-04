package generatestacks

import (
	"fmt"
	"slices"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GenerateStacks(Stacks types.Stacks, Regions types.Regions) types.DecodedStacks {
	var DecodedStacks types.DecodedStacks

	// for _, environment := range Environments.Environment {
		for _, region := range Regions.Region {
			for _, stack := range Stacks.Stack {
				if slices.Contains(stack.ExcludeRegions, region.Name) {
					continue
				}
				// if slices.Contains(stack.ExcludeEnvironments, environment.Name) {
				// 	continue
				// }
				if !slices.Contains(stack.ExcludeRegions, region.Name) {
					parentPath := fmt.Sprintf("./stacks/%s", region.Name)
					DecodedStacks = processStack(stack, region, parentPath, DecodedStacks, stack.ExcludeRegions, stack.ExcludeEnvironments)
				}
			}
		}
	// }

	return DecodedStacks
}

// Helper function to recursively process a stack and its child stacks.
func processStack(
	stack types.Stack,
	region types.Region,
	// environment types.Environment,
	parentPath string,
	DecodedStacks types.DecodedStacks,
	ExcludeRegions []string,
	ExcludeEnvironments []string) types.DecodedStacks {

	tags := append(region.Tags, region.Name)
	// tags = append(tags, region.Name, environment.Name)

	Stack := types.DecodedStack{
		Name:                 stack.Name,
		Path:                 parentPath + "/" + stack.Name,
		ParentPath:           parentPath,
		Tags:                 append(tags, stack.Tags...),
		Description:          stack.Name,
		Region:               region.Name,
		Environment:          region.Name,
		DeployAsStack:        true,
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
			DecodedStacks = processStack(childStack, region, Stack.Path, DecodedStacks, ExcludeRegions, ExcludeEnvironments)
		}
	}

	return DecodedStacks
}
