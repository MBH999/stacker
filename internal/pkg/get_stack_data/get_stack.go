package getstackdata

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"log"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

func GetStack(stackPath string) types.UpdateStackConfig {
	var Stack types.UpdateStackConfig

	fullPath := fmt.Sprintf("%s/%s", stackPath, "stack.tm.hcl")

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCLFile(fullPath)
	if diags.HasErrors() {
		log.Fatalf("Error parsing HCL: %v", diags)
	}

	decodeDiags := gohcl.DecodeBody(file.Body, nil, &Stack)
	if decodeDiags.HasErrors() {
		log.Fatalf("Error decoding HCL: %v", decodeDiags)
	}

	return Stack

}
