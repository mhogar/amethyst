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

	users := []*example.User{
		example.CreateNewUser("aaaaaaaaaa", "abc"),
		example.CreateNewUser("aaaaaaaaaa", "abcdefgdfgds"),
	}

	for _, user := range users {
		fmt.Printf("%s:\n", user.Password)
		p.Run(user)
		fmt.Println()
	}
}
