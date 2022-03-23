package main

import (
	"github.com/mhogar/kiwi/data"
	sqladapter "github.com/mhogar/kiwi/data/adapter/sql_adapter"
	"github.com/mhogar/kiwi/data/query"
	"github.com/mhogar/kiwi/example"
)

func main() {
	// f := nodes.NodeFactory{}

	// w := f.Workflow(
	// 	f.Validation(example.UserValidator{}),
	// 	f.Converter(example.UserConverter{}),
	// )

	// fmt.Println(w.Run(
	// 	example.CreateNewUserInput("username", "Password123!"),
	// ))

	adapter := &sqladapter.SqlAdapter{}
	handle := data.GetHandle[example.User](adapter)

	users, _ := handle.Read(
		query.Where("username", "=", "username").And(query.Where("rank", ">", 0)),
	)

	println(users[0].Username)
}
