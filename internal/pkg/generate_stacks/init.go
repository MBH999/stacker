package generatestacks

import (
	"github.com/charmbracelet/log"
	checkstacks "github.com/tmtf-stacker/stacker/internal/pkg/check_stacks"
	createstack "github.com/tmtf-stacker/stacker/internal/pkg/create_stack"
	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func Init(Config types.Config) {
	EnvironmentStacks := GenerateEnvironments(Config.Environments)
	RegionStacks := GenerateRegions(Config.Regions, EnvironmentStacks)
	Stacks := GenerateStacks(Config.Stacks, RegionStacks, EnvironmentStacks)

	var StackConfig types.DecodedStacks

	if Config.Environments.CreateEnvironmentStacks {
		StackConfig.DecodedStack = append(StackConfig.DecodedStack, EnvironmentStacks.DecodedStack...)
	}

	if Config.Regions.CreateRegionStacks {
		StackConfig.DecodedStack = append(StackConfig.DecodedStack, RegionStacks.DecodedStack...)
	}

	StackConfig.DecodedStack = append(StackConfig.DecodedStack, Stacks.DecodedStack...)

	for _, Stack := range StackConfig.DecodedStack {
		stackExists := checkstacks.CheckStacks(Stack.Path)

		if stackExists {
			log.Warnf("Stack %s exists!", Stack.Path)
			// checkstacks.CheckTags(Stack)
		} else {
			err := createstack.CreateStack(Stack)
			if err != nil {
				log.Errorf("unable to create stack %s", Stack.Path)
			}
		}
	}
}
