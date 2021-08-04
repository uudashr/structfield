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
	acc := Account{ // want "Found 5 non-labeled fields on struct literal \\(> 2\\)"
		"John Smith",
		"john.smith@example.com",
		[]Permission{
			Permission{"account", "read"},
			Permission{"account", "write"},
		},
		true,
		false,
	}
	fmt.Printf("%+v", acc)
}
