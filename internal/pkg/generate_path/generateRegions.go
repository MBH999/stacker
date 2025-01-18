package generatepath

import "github.com/tmtf-stacker/stacker/internal/pkg/types"

func GenerateRegions(Regions types.Regions, EnvironmentPaths []string) []string {
	var paths []string

	for _, environment := range EnvironmentPaths {
		for _, region := range Regions.Config {
			path := environment + "/" + region
			paths = append(paths, path)
		}
	}
	return paths
}
