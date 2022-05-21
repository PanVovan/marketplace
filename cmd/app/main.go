package main

import (
	"makretplace/internal/user"
	marketplace "makretplace/pkg/httpserver"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	routes := mux.NewRouter()
	userModule := &user.Module{}
	userModule.Configure(nil, routes)

	server := new(marketplace.Server)
	server.Run("localhost", "8080", routes)
}
