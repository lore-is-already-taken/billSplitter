package reposiroty

import (
	db_config "billAPI/cmd/reposiroty/config"
	"log/slog"
)

func InitializeBD() {
	_, err := db_config.InitializeDB()
	if err != nil {
		slog.Error(err.Message)
		panic(err.Err)
	}
}
