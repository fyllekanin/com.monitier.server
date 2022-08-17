package application

import (
	"database/sql"
	"github.com/fyllekanin/com.monitier.server/config"
	"github.com/gorilla/mux"
)

type Application struct {
	Db     *sql.DB           `json:"db"`
	Router *mux.Router       `json:"router"`
	Config *config.AppConfig `json:"config"`
}

func GetNewApplication(router *mux.Router, db *sql.DB, config *config.AppConfig) *Application {
	return &Application{
		Db:     db,
		Router: router,
		Config: config,
	}
}
