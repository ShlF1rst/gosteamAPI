package API

import (
	"fmt"
	"strings"
)

type Interface struct {
	Name    string   `json:"name"`
	Methods []Method `json:"methods"`
}

func (inter Interface) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Interface %s\nMethods:\n", inter.Name))
	lastIndex := len(inter.Methods) - 1
	for index, method := range inter.Methods {
		builder.WriteString(fmt.Sprintf("\t%s", strings.Replace(method.String(), "\n\t", "\n\t\t", -1)))
		if index != lastIndex {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
