package main

import (
	"database/sql"
	"fmt"
	"github.com/fyllekanin/com.monitier.server/api"
	"github.com/fyllekanin/com.monitier.server/application"
	"github.com/fyllekanin/com.monitier.server/config"
	"github.com/fyllekanin/com.monitier.server/database"
	"github.com/fyllekanin/com.monitier.server/task"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	var conf = config.GetAppConfig()
	var db = getDatabaseConnection(conf)
	runSqlFiles(db)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	var app = application.GetNewApplication(apiRouter, db, conf)
	task.StartScheduler(app)

	api.GetApi(app)
	// auth_api.GetApi(application)

	fmt.Println(fmt.Sprintf("Server running on %s", app.Config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", app.Config.Port), router))
}

func runSqlFiles(db *sql.DB) {
	files, err := os.ReadDir("sql/")
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, file := range files {
		c, err := os.ReadFile(fmt.Sprintf("sql/%s", file.Name()))
		if err != nil {
			log.Fatalln(err.Error())
		}
		sqlString := string(c)
		_, err = db.Exec(sqlString)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func getDatabaseConnection(config *config.AppConfig) *sql.DB {
	switch config.DatabaseType {
	case "SQLite":
		return database.GetSqliteDatabase()
		break
	case "Postgres":
		return database.GetPostgresDatabase(config)
		break
	default:
		log.Fatalln(fmt.Sprintf("Database type %s is not supported", config.DatabaseType))
	}
	return nil
}
