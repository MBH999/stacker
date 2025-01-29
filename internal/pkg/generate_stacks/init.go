package generatestacks

import (
	"fmt"
	"slices"

	"github.com/charmbracelet/log"
	checkstacks "github.com/tmtf-stacker/stacker/internal/pkg/check_stacks"
	createstack "github.com/tmtf-stacker/stacker/internal/pkg/create_stack"
	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func Init(Config types.Config) {
	EnvironmentStacks := GenerateEnvironments(Config.Environments)
	RegionStacks := GenerateRegions(Config.Regions, Config.Environments)
	Stacks := GenerateStacks(Config.Stacks, Config.Regions, Config.Environments)

	var StackConfig types.DecodedStacks

	for i := 0; i < len(Config.Environments.Environment); i++ {
		fmt.Printf("Environment %s - DeployAsStack %v\n", Config.Environments.Environment[i].Name, Config.Environments.Environment[i])

		if !Config.Environments.Environment[i].DeployAsStack {
			EnvironmentStacks.DecodedStack = slices.Delete(EnvironmentStacks.DecodedStack, i, i+1)
		}
	}

	StackConfig.DecodedStack = append(StackConfig.DecodedStack, EnvironmentStacks.DecodedStack...)

	for i := 0; i < len(Config.Regions.Region); i++ {
		fmt.Printf("Region %s - DeployAsStack %v\n", Config.Regions.Region[i].Name, Config.Regions.Region[i])

		if !Config.Regions.Region[i].DeployAsStack {
			RegionStacks.DecodedStack = slices.Delete(RegionStacks.DecodedStack, i, i+1)
		}
	}

	StackConfig.DecodedStack = append(StackConfig.DecodedStack, RegionStacks.DecodedStack...)

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
