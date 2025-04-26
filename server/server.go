package server

import (
	"auth-service/internal/services"
	"fmt"
	"log"
	"net/http"
)

func Start() error {
	http.HandleFunc("/auth/signup", signupHandler)
	log.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	services.IsUsernameExist("qwes")

	fmt.Fprint(w, "signup")
}
