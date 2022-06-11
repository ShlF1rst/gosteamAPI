package API

import "fmt"

type Parameter struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description"`
}

func (param Parameter) String() string {
	return fmt.Sprintf("Parameter %s\n\tType: %s\n\tOptional: %v\n\tDescription: %s", param.Name, param.Type, param.Optional, param.Description)
}
