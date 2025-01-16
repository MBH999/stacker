package stackdecode

import (
	"log"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	generatepath "github.com/tmtf-stacker/stacker/internal/pkg/generate_path"

	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func StackDecode() {
	// HCL file path
	filename := "./config/examples/basic.hcl"

	// Parse the HCL file
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCLFile(filename)
	if diags.HasErrors() {
		log.Fatalf("Error parsing HCL: %v", diags)
	}

	// Decode the HCL file into the Config struct
	var config types.Config
	decodeDiags := gohcl.DecodeBody(file.Body, nil, &config)
	if decodeDiags.HasErrors() {
		log.Fatalf("Error decoding HCL: %v", decodeDiags)
	}

	// Print the parsed Config
	// log.Printf("Parsed Config: %+v", config)

	// Print details for each stack
	for _, stack := range config.Stacks.Stack {
		generatepath.GeneratePath(stack)

		if stack.Stack != nil {
			for _, childStack := range stack.Stack {
				generatepath.GeneratePath(childStack, stack)
			}
		}

		// log.Printf("Stack Name: %s", stack.Name)
		// log.Printf("Tags: %v", stack.Tags)
		// log.Printf("Exclude Regions: %v", stack.ExcludeRegions)
		// log.Printf("Exclude Environments: %v", stack.ExcludeEnvironments)
		// for _, childStack := range stack.Stack {
		// 	log.Printf("Child Stacks: %v", childStack.Name)
		// }
	}
}
