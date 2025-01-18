package generatepath

import "github.com/tmtf-stacker/stacker/internal/pkg/types"

func GenerateStacks(Stacks types.Stacks, RegionPaths []string) []string {
	var paths []string

	for _, region := range RegionPaths {
		for _, stack := range Stacks.Stack {
			path := region + "/" + stack.Name
			paths = append(paths, path)
			if stack.ChildStack != nil {
				for _, childStack := range stack.ChildStack {
					childStack := path + "/" + childStack.Name
					paths = append(paths, childStack)
				}
			}
		}
	}
	return paths
}
