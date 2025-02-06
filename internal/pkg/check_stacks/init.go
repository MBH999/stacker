package checkstacks

import (
	"slices"
	"strings"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func Init(StackConfig types.DecodedStack, ExistingStack types.UpdateStackConfig) types.UpdateStackConfig {
	slices.Sort(StackConfig.Tags)
	slices.Sort(ExistingStack.Stack.Tags)

	if !slices.Equal(StackConfig.Tags, ExistingStack.Stack.Tags) {
		ExistingStack.Stack.Tags = StackConfig.Tags
	}

	if strings.Compare(StackConfig.Description, ExistingStack.Stack.Description) != 0 {
		ExistingStack.Stack.Description = StackConfig.Description
	}

	// if StackConfig.Description != ExistingStack.Stack.Description {
	// 	ExistingStack.Stack.Description = StackConfig.Description
	// }

	return ExistingStack

}
