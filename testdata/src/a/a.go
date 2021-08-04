package main

import (
	"fmt"
)

type Account struct {
	Name        string
	Email       string
	Permissions []Permission
	Verified    bool
	Deactivated bool
}

type Permission struct {
	Domain string
	Access string
}

func main() {
	acc := Account{ // want "Found 5 non-labeled fields on struct literal \\(> 0\\)"
		"John Smith",
		"john.smith@example.com",
		[]Permission{
			Permission{"account", "read"},  // want "Found 2 non-labeled fields on struct literal \\(> 0\\)"
			Permission{"account", "write"}, // want "Found 2 non-labeled fields on struct literal \\(> 0\\)"
		},
		true,
		false,
	}
	fmt.Printf("%+v", acc)
}
