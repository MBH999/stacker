package updatestacks

import (
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

//go:embed templates/stack.tm.hcl.templ
var templFile embed.FS

func toHCLList(slice []string) string {
	var quoted []string
	for _, s := range slice {
		quoted = append(quoted, fmt.Sprintf("%q", s))
	}
	return fmt.Sprintf("[%s]", strings.Join(quoted, ", "))
}

func UpdateStack(existingStack types.UpdateStackConfig, stackPath string) error {
	funcMap := template.FuncMap{
		"toHCLList": toHCLList,
	}

	tmpl, err := template.New("stack.tm.hcl.templ").Funcs(funcMap).ParseFS(templFile, "templates/stack.tm.hcl.templ")
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	outFile, err := os.Create(stackPath)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, existingStack); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	return nil
}

