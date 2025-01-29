package generatestacks

import "github.com/tmtf-stacker/stacker/internal/pkg/types"

func GenerateRegions(Regions types.Regions, EnvironmentStacks types.DecodedStacks) types.DecodedStacks {
	var Stacks types.DecodedStacks

	for _, environment := range EnvironmentStacks.DecodedStack {
		for _, region := range Regions.Config {
			Stack := types.DecodedStack{
				Name:                 region,
				Path:                 environment.Path + "/" + region,
				ParentPath:           environment.ParentPath,
				Tags:                 append(environment.Tags, region),
				Description:          region,
				Region:               region,
				Environment:          environment.Environment,
				ExcludedEnvironments: []string{},
				ExcludedRegions:      []string{},
			}
			Stacks.DecodedStack = append(Stacks.DecodedStack, Stack)
		}
	}
	return Stacks
}
