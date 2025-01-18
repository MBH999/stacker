package generatestacks

import (
	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GenerateEnvironments(Environment types.Environments) types.DecodedStacks {
	var Stacks types.DecodedStacks

	for _, environment := range Environment.Config {
		var Stack types.DecodedStack
		Stack.Path = "./stacks/" + environment
		Stack.Tags = append(Stack.Tags, environment)

		Stacks.DecodedStack = append(Stacks.DecodedStack, Stack)
	}
	return Stacks
}
