package generatestacks

import "github.com/tmtf-stacker/stacker/internal/pkg/types"

func GenerateStacks(Stacks types.Stacks, RegionStacks types.DecodedStacks) types.DecodedStacks {
	var DecodedStacks types.DecodedStacks

	for _, region := range RegionStacks.DecodedStack {
		for _, stack := range Stacks.Stack {

			var DecodedStack types.DecodedStack

			DecodedStack.Path = region.Path + "/" + stack.Name
			DecodedStack.Tags = append(DecodedStack.Tags, stack.Tags...)

			DecodedStacks.DecodedStack = append(DecodedStacks.DecodedStack, DecodedStack)

			if stack.ChildStack != nil {
				for _, childStack := range stack.ChildStack {

					var ChildDecodedStack types.DecodedStack

					ChildDecodedStack.Path = DecodedStack.Path + "/" + childStack.Name
					ChildDecodedStack.Tags = append(ChildDecodedStack.Tags, childStack.Tags...)

					DecodedStacks.DecodedStack = append(DecodedStacks.DecodedStack, ChildDecodedStack)

				}
			}
		}
	}
	return DecodedStacks
}
