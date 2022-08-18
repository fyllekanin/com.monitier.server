package repository

import (
	"database/sql"
	"errors"
	"github.com/fyllekanin/com.monitier.server/config"
	"github.com/fyllekanin/com.monitier.server/database/entity"
	"log"
	"time"
)

type PingRepository struct {
	db *sql.DB
}

func (pingRepository *PingRepository) InsertPing(service config.AppService, responseTime int) error {
	_, err := pingRepository.db.Exec("INSERT INTO pings (serviceName, responseTime, createdAt, updatedAt) VALUES ($1, $2, $3, $4)",
		service.Name, responseTime, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to query statement")
	}
	return nil
}

func (pingRepository *PingRepository) GetPings() ([]entity.PingEntity, error) {
	statement, err := pingRepository.db.Prepare("SELECT * FROM pings")
	if err != nil {
		log.Println(err)
		return []entity.PingEntity{}, errors.New("failed to prepare statement")
	}
	rows, err := statement.Query()
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return []entity.PingEntity{}, errors.New("failed to query statement")
	}

	var response = []entity.PingEntity{}
	for rows.Next() {
		var item entity.PingEntity
		rows.Scan(&item.ServiceName, &item.ResponseTime, &item.CreatedAt, &item.UpdatedAt)
		response = append(response, item)
	}
	return response, nil
}

func NewPingRepository(db *sql.DB) *PingRepository {
	return &PingRepository{
		db: db,
	}
}
