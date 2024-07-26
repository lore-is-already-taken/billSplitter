package routes

import (
	"net/http"
)

func AddUser() {
	mux.HandleFunc("GET "+"/adduser", showHello)

}
func showHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User added"))
}
