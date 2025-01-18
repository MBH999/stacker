package stackdecode

import (
	"log"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"

	generatestacks "github.com/tmtf-stacker/stacker/internal/pkg/generate_stacks"
	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func StackDecode() {
	filename := "./config/examples/basic.hcl"

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCLFile(filename)
	if diags.HasErrors() {
		log.Fatalf("Error parsing HCL: %v", diags)
	}

	var config types.Config
	decodeDiags := gohcl.DecodeBody(file.Body, nil, &config)
	if decodeDiags.HasErrors() {
		log.Fatalf("Error decoding HCL: %v", decodeDiags)
	}

	generatestacks.Init(config)

	// for _, region := range config.Regions.Config {
	// 	fmt.Printf("Regions: %s\n", region)
	// }
	// fmt.Println(config.Regions.CreateRegionStacks)
	//
	// for _, environment := range config.Environments.Config {
	// 	fmt.Printf("Environments: %s\n", environment)
	// }
	// fmt.Println(config.Environments.CreateEnvironmentStacks)
	//
	// // Print details for each stack
	// for _, stack := range config.Stacks.Stack {
	// 	generatepath.GeneratePath(stack)
	//
	// 	if stack.Stack != nil {
	// 		for _, childStack := range stack.Stack {
	// 			generatepath.GeneratePath(childStack, stack)
	// 		}
	// 	}
	// }
}
