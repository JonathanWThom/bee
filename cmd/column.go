package cmd

import "fmt"

type Column struct {
	Name string
	Type string
}

func (c Column) String() string {
	return fmt.Sprintf("%s:%s\n", c.Name, c.Type)
}

func validType(typ string) bool {
	valid := []string{"int", "float", "string", "bool"}

	return includes(valid, typ)
}
