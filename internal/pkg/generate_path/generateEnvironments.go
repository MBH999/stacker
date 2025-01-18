package generatepath

import (
	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GenerateEnvironments(Environment types.Environments) []string {
	var paths []string

	for _, environment := range Environment.Config {
		path := "./stacks/" + environment
		paths = append(paths, path)
	}
	return paths
}
