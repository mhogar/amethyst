package main

import (
	"fmt"
	"net/http"

	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/example/user"
	"github.com/mhogar/kiwi/nodes/web"

	"github.com/julienschmidt/httprouter"
)

func createRouter(adapter adapter.DataAdapter) *httprouter.Router {
	r := httprouter.New()

	r.GET("/user",
		web.NewHandler(adapter, user.GetUsersWorkflow()).ServeHTTPRouter,
	)
	r.GET("/user/:username",
		web.NewHandler(adapter, user.GetUserWorkflow()).ServeHTTPRouter,
	)
	r.POST("/user",
		web.NewHandler(adapter, user.CreateUserWorkflow()).ServeHTTPRouter,
	)
	r.PUT("/user/:username",
		web.NewHandler(adapter, user.UpdateUserWorkflow()).ServeHTTPRouter,
	)
	r.PATCH("/user/:username/password",
		web.NewHandler(adapter, user.UpdateUserAuthWorkflow()).ServeHTTPRouter,
	)
	r.DELETE("/user/:username",
		web.NewHandler(adapter, user.DeleteUserWorkflow()).ServeHTTPRouter,
	)

	return r
}

func main() {
	adapter := dependencies.DataAdapter.Resolve()

	err := adapter.Setup()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer adapter.CleanUp()

	server := http.Server{
		Addr:    ":8080",
		Handler: createRouter(adapter),
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
