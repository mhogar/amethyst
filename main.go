package main

import (
	"fmt"
	"net/http"

	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/example/session"
	"github.com/mhogar/kiwi/example/user"
	"github.com/mhogar/kiwi/nodes/web"

	"github.com/julienschmidt/httprouter"
)

func createRouter(adapter adapter.DataAdapter) *httprouter.Router {
	r := httprouter.New()

	// user routes
	r.GET("/user",
		web.NewHandler(adapter, user.GetUserEndpoint()).ServeHTTPRouter,
	)
	r.GET("/user/:username",
		web.NewHandler(adapter, user.GetUserEndpoint()).ServeHTTPRouter,
	)
	r.POST("/user",
		web.NewHandler(adapter, user.CreateUserEndpoint()).ServeHTTPRouter,
	)
	r.PUT("/user",
		web.NewHandler(adapter, user.UpdateUserEndpoint()).ServeHTTPRouter,
	)
	r.PATCH("/user/password",
		web.NewHandler(adapter, user.UpdateUserAuthEndpoint()).ServeHTTPRouter,
	)
	r.DELETE("/user",
		web.NewHandler(adapter, user.DeleteUserEndpoint()).ServeHTTPRouter,
	)

	// session routes
	r.POST("/session",
		web.NewHandler(adapter, session.CreateSessionEndpoint()).ServeHTTPRouter,
	)
	r.DELETE("/session",
		web.NewHandler(adapter, session.DeleteSessionEndpoint()).ServeHTTPRouter,
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
