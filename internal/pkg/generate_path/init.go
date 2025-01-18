package generatepath

import (
	createstack "github.com/tmtf-stacker/stacker/internal/pkg/create_stack"
	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func Init(Config types.Config) {
	environmentPaths := GenerateEnvironments(Config.Environments)
	regionPaths := GenerateRegions(Config.Regions, environmentPaths)
	stackPaths := GenerateStacks(Config.Stacks, regionPaths)

	var joinedPaths []string

	if Config.Environments.CreateEnvironmentStacks {
		joinedPaths = append(joinedPaths, environmentPaths...)
	}

	if Config.Regions.CreateRegionStacks {
		joinedPaths = append(joinedPaths, regionPaths...)
	}

	joinedPaths = append(joinedPaths, stackPaths...)

	createstack.CreateStacks(joinedPaths)
}
