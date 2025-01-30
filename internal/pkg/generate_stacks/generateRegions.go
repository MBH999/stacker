package generatestacks

import (
	"fmt"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GenerateRegions(Regions types.Regions, Environments types.Environments) types.DecodedStacks {
	var Stacks types.DecodedStacks

	for _, environment := range Environments.Environment {
		for _, region := range Regions.Region {
			Stack := types.DecodedStack{
				Name:                 region.Name,
				Path:                 fmt.Sprintf("./stacks/%s/%s/", environment.Name, region.Name),
				ParentPath:           fmt.Sprintf("./stacks/%s", environment.Name),
				Tags:                 append(environment.Tags, region.Name),
				Description:          region.Name,
				Region:               region.Name,
				Environment:          environment.Name,
				DeployAsStack:        region.DeployAsStack,
				ExcludedEnvironments: []string{},
				ExcludedRegions:      []string{},
			}
			Stacks.DecodedStack = append(Stacks.DecodedStack, Stack)
		}
	}
	return Stacks
}
