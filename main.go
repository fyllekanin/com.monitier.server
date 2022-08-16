package main

import (
	"database/sql"
	"fmt"
	"github.com/fyllekanin/com.monitier.server/configs"
	"github.com/fyllekanin/com.monitier.server/databases"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Application struct {
	Db     *sql.DB            `json:"db"`
	Router *mux.Router        `json:"router"`
	Config *configs.AppConfig `json:"config"`
}

func GetNewApplication(router *mux.Router, db *sql.DB, config *configs.AppConfig) *Application {
	return &Application{
		Db:     db,
		Router: router,
		Config: config,
	}
}

func main() {
	var config = configs.GetAppConfig()
	var db = getDatabaseConnection(config)
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	var application = GetNewApplication(apiRouter, db, config)

	// product_api.GetApi(application)
	// auth_api.GetApi(application)

	fmt.Println(fmt.Sprintf("Server running on %d", application.Config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", application.Config.Port), router))
}

func getDatabaseConnection(config *configs.AppConfig) *sql.DB {
	switch config.DatabaseType {
	case "SQLite":
		break
	case "Postgres":
		return databases.GetPostgresDatabase(config)
		break
	default:
		log.Fatalln(fmt.Sprintf("Database type %s is not supported", config.DatabaseType))
	}
	return nil
}
