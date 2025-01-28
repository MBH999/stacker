package generatestacks

import (
	"fmt"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GenerateStacks(Stacks types.Stacks, RegionStacks types.DecodedStacks) types.DecodedStacks {
	var DecodedStacks types.DecodedStacks

	for _, region := range RegionStacks.DecodedStack {
		for _, stack := range Stacks.Stack {
			DecodedStacks = processStack(stack, region, region.Path, DecodedStacks)
		}
	}

	return DecodedStacks
}

// Helper function to recursively process a stack and its child stacks.
func processStack(stack types.Stack, region types.DecodedStack, parentPath string, DecodedStacks types.DecodedStacks) types.DecodedStacks {
	var DecodedStack types.DecodedStack

	DecodedStack.Name = stack.Name
	DecodedStack.Path = parentPath + "/" + stack.Name
	DecodedStack.ParentPath = parentPath + "/"
	DecodedStack.Tags = append(DecodedStack.Tags, stack.Tags...)
	DecodedStack.Tags = append(DecodedStack.Tags, region.Tags...)

	if stack.Description != "" {
		DecodedStack.Description = stack.Description
		fmt.Println(DecodedStack.Description)
	} else {
		DecodedStack.Description = stack.Name
	}

	DecodedStacks.DecodedStack = append(DecodedStacks.DecodedStack, DecodedStack)

	if stack.ChildStack != nil {
		for _, childStack := range stack.ChildStack {
			DecodedStacks = processStack(childStack, region, DecodedStack.Path, DecodedStacks)
		}
	}

	return DecodedStacks
}
