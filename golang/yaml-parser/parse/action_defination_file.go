package parse

import (
	"fmt"
	"go-lang-learning/yaml-parser/dto"
)

func ParseActionDefinationFile() dto.ActionDefination {
	fmt.Println("Testing")
	action_defination := dto.ActionDefination{
		ADFVersion: "1.0",
		Actions:    []dto.Action{},
	}
	return action_defination
}
