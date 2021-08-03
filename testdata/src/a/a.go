package main

import (
	"fmt"
)

type Account struct {
	Name        string
	Email       string
	Permissions []Permission
	Deactivated bool
}

type Permission struct {
	Domain string
	Access string
}

func main() {
	acc := Account{ // want "Found 4 non-labeled fields on struct literal \\(> 0\\)"
		"Nuruddin Ashr",
		"uudashr@gmail.com",
		[]Permission{
			Permission{"account", "read"},  // want "Found 2 non-labeled fields on struct literal \\(> 0\\)"
			Permission{"account", "write"}, // want "Found 2 non-labeled fields on struct literal \\(> 0\\)"
		},
		false,
	}
	fmt.Printf("%+v", acc)
}
