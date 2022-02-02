package main

import (
	"fmt"

	"github.com/amethyst/example"
	"github.com/amethyst/nodes"
)

func main() {
	p := nodes.Pipeline{}
	p = append(p, nodes.ValidationNode{
		Validator: &example.UserValidator{},
	})

	usernames := []string{
		"aaaa",
		"aaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}

	for _, name := range usernames {
		fmt.Printf("%s:\n", name)
		p.Run(example.CreateNewUser(name))
		fmt.Println()
	}
}
