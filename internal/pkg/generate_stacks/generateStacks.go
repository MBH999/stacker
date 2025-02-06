package generatestacks

import (
	"fmt"
	"slices"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GenerateStacks(Stacks types.Stacks, Regions types.Regions) types.DecodedStacks {
	var DecodedStacks types.DecodedStacks

	for _, region := range Regions.Region {

		for _, stack := range Stacks.Stack {

			if slices.Contains(stack.ExcludeRegions, region.Name) {
				continue
			}

			if !slices.Contains(stack.ExcludeRegions, region.Name) {
				parentPath := fmt.Sprintf("./stacks/%s", region.Name)
				DecodedStacks = processStack(stack, region, parentPath, DecodedStacks, stack.ExcludeRegions, region.Tags)
			}

		}

	}

	return DecodedStacks

}

func processStack(
	stack types.Stack,
	region types.Region,
	parentPath string,
	DecodedStacks types.DecodedStacks,
	ExcludeRegions []string,
	ParentTags []string) types.DecodedStacks {

	tags := append(region.Tags, region.Name)
	tags = append(tags, stack.Tags...)
	tags = append(tags, ParentTags...)
	tags = append(tags, stack.Name)

	tags = removeDuplicates(tags)

	var stackDescription string
	if stack.Description != "" {
		stackDescription = stack.Description
	} else {
		stackDescription = stack.Name
	}

	Stack := types.DecodedStack{
		Name:            stack.Name,
		Path:            parentPath + "/" + stack.Name,
		ParentPath:      parentPath,
		Tags:            tags,
		Description:     stackDescription,
		Region:          region.Name,
		DeployAsStack:   true,
		ExcludedRegions: ExcludeRegions,
	}

	DecodedStacks.DecodedStack = append(DecodedStacks.DecodedStack, Stack)

	if stack.ChildStack != nil {
		for _, childStack := range stack.ChildStack {
			DecodedStacks = processStack(childStack, region, Stack.Path, DecodedStacks, ExcludeRegions, tags)
		}
	}

	return DecodedStacks
}

func removeDuplicates(tags []string) []string {
	seen := make(map[string]struct{})
	var unique []string
	for _, tag := range tags {
		if _, ok := seen[tag]; !ok {
			seen[tag] = struct{}{}
			unique = append(unique, tag)
		}
	}
	return unique
}
