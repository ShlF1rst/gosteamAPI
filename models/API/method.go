package API

import (
	"fmt"
	"strings"
)

type Method struct {
	Name       string      `json:"name"`
	Version    uint8       `json:"version"`
	HTTPMethod string      `json:"httpmethod"`
	Parameters []Parameter `json:"parameters"`
}

func (method Method) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Method %s\n\tVersion: %d\n\tHTTP method: %s\n\tParameters:\n", method.Name, method.Version, method.HTTPMethod))
	lastIndex := len(method.Parameters) - 1
	for index, param := range method.Parameters {
		builder.WriteString(fmt.Sprintf("\t\t%s", strings.Replace(param.String(), "\n\t", "\n\t\t\t", -1)))
		if index != lastIndex {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}
