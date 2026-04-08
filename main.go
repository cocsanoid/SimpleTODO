package main

import (
	"fmt"
	"todo/http"
	"todo/structures"
)

func main() {
	list := &structures.TodoList{}
	httpHandlers := http.NewHttpHandlers(list)
	httpServer := http.NewHTTPServer(httpHandlers)
	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
