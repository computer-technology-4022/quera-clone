package main

import (
	"fmt"
	"net/http"

	"github.com/computer-technology-4022/goera/internal/config"
	handler "github.com/computer-technology-4022/goera/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle(config.StaticRouter, http.
		StripPrefix(config.StaticRouter, http.FileServer(http.Dir(config.StaticRouterDir))))
	r.HandleFunc("/", handler.WelcomeHandler)
	r.HandleFunc("/login", handler.LoginHandler)
	r.HandleFunc("/signUp", handler.SignUpHandler)
	r.HandleFunc("/questions", handler.QuestionsHandler)
	r.HandleFunc("/question", handler.QuestionHandler)
	fmt.Println("Server is running on http://localhost:8080")

	// s := r.PathPrefix("/api").Subrouter()
	// s.HandleFunc("/login", api.LoginHandler).Methods("GET", "POST")
	// s.HandleFunc("/register", api.RegisterHandler).Methods("GET", "POST")
	// s.HandleFunc("/user/{ID}", auth.JWTMiddleware(api.UsersHandler)).Methods("GET", "POST")

	http.Handle("/", r)
	http.ListenAndServe(config.ServerPort, nil)
}
