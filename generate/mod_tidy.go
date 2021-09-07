package generate

import (
	"log"
	"strings"
)

func ModTidy(moduleName string)  {
	template := strings.Replace(modTidyTemplate, "{:MODULE:}", moduleName, -1)
	log.Println(template)
}

var modTidyTemplate = `Run the following command to finish !!
    cd {:MODULE:} && go mod tidy
`
