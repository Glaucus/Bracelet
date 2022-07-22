package httpServer

import (
	"github.com/Glaucus/Bracelet/httpServer/controllers/login/password"

	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func Start() {
	apiRouter := mux.NewRouter().StrictSlash(true)
	apiRouter.HandleFunc("/login/password", password.Controller)
	log.Fatal(http.ListenAndServe(":80", apiRouter))
}
