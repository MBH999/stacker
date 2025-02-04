package generatestacks

import (
	"fmt"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GenerateRegions(Regions types.Regions) types.DecodedStacks {
	var Stacks types.DecodedStacks

	// for _, environment := range Environments.Environment {
		for _, region := range Regions.Region {
			regionTags := append(region.Tags, region.Name)

			Stack := types.DecodedStack{
				Name:                 region.Name,
				Path:                 fmt.Sprintf("./stacks/%s/", region.Name),
				ParentPath:           fmt.Sprintf("./stacks/"),
				Tags:                 regionTags,
				Description:          region.Name,
				Region:               region.Name,
				Environment:          region.Name,
				DeployAsStack:        region.DeployAsStack,
				ExcludedEnvironments: []string{},
				ExcludedRegions:      []string{},
			}
			Stacks.DecodedStack = append(Stacks.DecodedStack, Stack)
		}
	// }
	return Stacks
}
