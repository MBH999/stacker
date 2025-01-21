package updatestacks

import (
	"embed"
	"os"
	"text/template"

	"github.com/charmbracelet/log"
	getstackdata "github.com/tmtf-stacker/stacker/internal/pkg/get_stack_data"
	"github.com/tmtf-stacker/stacker/internal/pkg/types"
)

//go:embed templates/stack.tm.hcl.templ
var tmplFiles embed.FS

func UpdateTags(path string, tags string) {
	tmplContent, err := tmplFiles.ReadFile("templates/stack.tm.hcl.templ")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("stack").Parse(string(tmplContent))
	if err != nil {
		log.Error(err)
	}

	stackData := types.UpdateStack{
		Name:        getstackdata.GetStackName(path),
		ID:          getstackdata.GetStackID(path),
		Description: getstackdata.GetStackDescription(path),
		Tags:        tags,
	}

	file, err := os.Create(path + "/stack.tm.hcl")
	if err != nil {
		log.Error(err)
	}
	defer file.Close()

	tmpl.Execute(file, stackData)
}
