package generatestacks

import "github.com/tmtf-stacker/stacker/internal/pkg/types"

func GenerateRegions(Regions types.Regions, EnvironmentStacks types.DecodedStacks) types.DecodedStacks {
	var Stacks types.DecodedStacks

	for _, environment := range EnvironmentStacks.DecodedStack {
		for _, region := range Regions.Config {
			var Stack types.DecodedStack
			Stack.Path = environment.Path + "/" + region
			Stack.Tags = append(Stack.Tags, region)
			Stack.Tags = append(Stack.Tags, environment.Tags...)
			Stacks.DecodedStack = append(Stacks.DecodedStack, Stack)
		}
	}
	return Stacks
}
