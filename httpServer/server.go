package httpServer

import (
	loginPassword "github.com/Glaucus/Bracelet/httpServer/controllers/login/password"
	registerPassword "github.com/Glaucus/Bracelet/httpServer/controllers/register/password"

	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func Start() {
	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.HandleFunc("/login/password", loginPassword.Controller)
	apiRouter.HandleFunc("/register/password", registerPassword.Controller)
	log.Fatal(http.ListenAndServe(":80", apiRouter))
}
