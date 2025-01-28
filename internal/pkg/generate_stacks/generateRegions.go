package generatestacks

import "github.com/tmtf-stacker/stacker/internal/pkg/types"

func GenerateRegions(Regions types.Regions, EnvironmentStacks types.DecodedStacks) types.DecodedStacks {
	var Stacks types.DecodedStacks

	for _, environment := range EnvironmentStacks.DecodedStack {
		for _, region := range Regions.Config {
			var Stack types.DecodedStack
			Stack.Name = region
			Stack.Path = environment.Path + "/" + region
			Stack.Tags = append(Stack.Tags, region)
			Stack.ParentPath = environment.Path
			Stack.Tags = append(Stack.Tags, environment.Tags...)
			Stack.Description = region
			Stacks.DecodedStack = append(Stacks.DecodedStack, Stack)
		}
	}
	return Stacks
}
