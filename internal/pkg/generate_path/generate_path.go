package generatepath

import (
	"fmt"

	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GeneratePath(Stack types.Stack, ParentStack ...types.Stack) {
	if len(ParentStack) == 1 {
		fmt.Printf("Parent Stack: %s, Child Stack %s\n", ParentStack[0].Name, Stack.Name)
	}
	fmt.Printf("Stack Name: %s\n", Stack.Name)
}
