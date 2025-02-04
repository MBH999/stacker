package generatestacks

import (
	// "fmt"
	// "slices"

	"github.com/charmbracelet/log"
	checkstacks "github.com/tmtf-stacker/stacker/internal/pkg/check_stacks"
	createstack "github.com/tmtf-stacker/stacker/internal/pkg/create_stack"
	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func Init(Config types.Config) {
	RegionStacks := GenerateRegions(Config.Regions)
	Stacks := GenerateStacks(Config.Stacks, Config.Regions)

	var StackConfig []types.DecodedStack

	for _, region := range RegionStacks.DecodedStack {
		if region.DeployAsStack {
			StackConfig = append(StackConfig, region)
		}
	}

	StackConfig = append(StackConfig, Stacks.DecodedStack...)

	for _, Stack := range StackConfig {
		stackExists := checkstacks.CheckStacks(Stack.Path)

		if stackExists {
			log.Warnf("Stack %s exists!", Stack.Path)
			checkstacks.CheckTags(Stack)
		} else {
			err := createstack.CreateStack(Stack)
			if err != nil {
				log.Errorf("unable to create stack %s", Stack.Path)
			}
		}
	}
}
