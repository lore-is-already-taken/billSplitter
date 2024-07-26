package main

import (
	"billAPI/cmd/reposiroty"
	"billAPI/cmd/routes"
	"context"
	"log"
	"log/slog"
	"net/http"
)

func initroutes() {
	routes.AddUser()

}
func main() {
	initroutes()
	reposiroty.InitializeBD()

	slog.Log(
		context.TODO(),
		slog.LevelInfo,
		"Starting server on port :4000",
	)
	err := http.ListenAndServe(":4000", routes.GetMuxInstance())
	log.Fatal(err)
}
