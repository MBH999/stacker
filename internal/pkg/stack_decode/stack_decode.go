package stackdecode

import (
	"log"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"

	generatestacks "github.com/tmtf-stacker/stacker/internal/pkg/generate_stacks"
	types "github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func StackDecode(configPath string) {
	filename := configPath

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
}
