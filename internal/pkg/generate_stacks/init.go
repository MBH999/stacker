package generatestacks

import (
	// "fmt"
	// "slices"

	"fmt"

	"github.com/charmbracelet/log"
	checkstacks "github.com/tmtf-stacker/stacker/internal/pkg/check_stacks"
	createstack "github.com/tmtf-stacker/stacker/internal/pkg/create_stack"
	getstackdata "github.com/tmtf-stacker/stacker/internal/pkg/get_stack_data"
	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
	updatestacks "github.com/tmtf-stacker/stacker/internal/pkg/update_stacks"
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
			existingStack := getstackdata.GetStack(Stack.Path)
			existingStack = checkstacks.Init(Stack, existingStack)
			updatestacks.UpdateStack(existingStack, fmt.Sprintf("%s/%s", Stack.Path, "stack.tm.hcl"))

		} else {
			err := createstack.CreateStack(Stack)
			if err != nil {
				log.Errorf("unable to create stack %s", Stack.Path)
			}
		}
	}
}
