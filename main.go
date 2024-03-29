package main

import (
	"log"
	"net/http"

	"github.com/phanorcoll/http_package/middleware"
	"github.com/phanorcoll/http_package/user"
)

func main() {
	userH := &user.Handler{}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", userH.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", userH.FindUserByID)

	adminRouter := http.NewServeMux()
	adminRouter.HandleFunc("POST /users", userH.CreateUser)
	adminRouter.HandleFunc("PUT /users/{id}", userH.UpdateUserByID)
	adminRouter.HandleFunc("DELETE /users/{id}", userH.DeleteUserByID)
	mux.Handle("/admin/", http.StripPrefix("/admin", adminRouter))

	frontEnd := http.NewServeMux()
	frontEnd.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	mux.Handle("/", frontEnd)

	chain := middleware.MiddlewareChain(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":3001",
		Handler: chain(mux), // v1, adm,
	}

	log.Println("Server started on port :3001")
	log.Fatal(server.ListenAndServe())
}
