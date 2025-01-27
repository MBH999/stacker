package generatestacks

import (
	"fmt"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

// func GenerateStacks(Stacks types.Stacks, RegionStacks types.DecodedStacks) types.DecodedStacks {
// 	var DecodedStacks types.DecodedStacks
//
// 	for _, region := range RegionStacks.DecodedStack {
// 		for _, stack := range Stacks.Stack {
//
// 			var DecodedStack types.DecodedStack
//
// 			DecodedStack.Path = region.Path + "/" + stack.Name
// 			DecodedStack.Tags = append(DecodedStack.Tags, stack.Tags...)
// 			DecodedStack.Tags = append(DecodedStack.Tags, region.Tags...)
//
// 			if stack.Description != "" {
// 				DecodedStack.Description = stack.Description
// 				fmt.Println(DecodedStack.Description)
// 			} else {
// 				DecodedStack.Description = stack.Name
// 			}
//
// 			DecodedStacks.DecodedStack = append(DecodedStacks.DecodedStack, DecodedStack)
//
// 			if stack.ChildStack != nil {
// 				for _, childStack := range stack.ChildStack {
//
// 					var ChildDecodedStack types.DecodedStack
//
// 					ChildDecodedStack.Path = DecodedStack.Path + "/" + childStack.Name
// 					ChildDecodedStack.Tags = append(ChildDecodedStack.Tags, childStack.Tags...)
// 					ChildDecodedStack.Tags = append(ChildDecodedStack.Tags, region.Tags...)
//
// 					if childStack.Description != "" {
// 						ChildDecodedStack.Description = childStack.Description
// 						fmt.Println(childStack.Description)
// 					} else {
// 						ChildDecodedStack.Description = childStack.Name
// 					}
//
// 					DecodedStacks.DecodedStack = append(DecodedStacks.DecodedStack, ChildDecodedStack)
//
// 				}
// 			}
// 		}
// 	}
// 	return DecodedStacks
// }

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

	// Construct the path for the current stack.
	DecodedStack.Path = parentPath + "/" + stack.Name
	DecodedStack.Tags = append(DecodedStack.Tags, stack.Tags...)
	DecodedStack.Tags = append(DecodedStack.Tags, region.Tags...)

	// Set the description for the current stack.
	if stack.Description != "" {
		DecodedStack.Description = stack.Description
		fmt.Println(DecodedStack.Description)
	} else {
		DecodedStack.Description = stack.Name
	}

	// Append the current stack to the result.
	DecodedStacks.DecodedStack = append(DecodedStacks.DecodedStack, DecodedStack)

	// Recursively process child stacks, if any.
	if stack.ChildStack != nil {
		for _, childStack := range stack.ChildStack {
			DecodedStacks = processStack(childStack, region, DecodedStack.Path, DecodedStacks)
		}
	}

	return DecodedStacks
}
