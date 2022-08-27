package task

import (
	"fmt"
	"github.com/fyllekanin/com.monitier.server/database/repository"
	"github.com/go-co-op/gocron"
	"log"
	"net/http"
	"time"
)

var scheduler = func() *gocron.Scheduler {
	var s = gocron.NewScheduler(time.UTC)
	s.StartAsync()
	return s
}()

type Service struct {
	Name      string      `json:"name"`
	Host      string      `json:"host"`
	IsRunning bool        `json:"isRunning"`
	job       *gocron.Job `json:"job""`
}

func (service *Service) Start(pingRepository *repository.PingRepository) {
	if service.IsRunning {
		log.Fatalln(fmt.Sprintf("Trying to start service %s but its already running", service.Name))
	}

	service.job, _ = scheduler.Every(15).Seconds().Do(func() {
		var start = time.Now()
		_, err := http.Get(service.Host)
		if err != nil {
			pingRepository.InsertPing(service.Name, -1)
		} else {
			var responseTime = int(time.Since(start).Milliseconds())
			pingRepository.InsertPing(service.Name, responseTime)
		}

		log.Println(fmt.Sprintf("Service %s pinged", service.Name))
	})
	service.job.Tag(service.Name)
	log.Println("Started")
	service.IsRunning = true
}

func (service *Service) Stop() {
	if service.IsRunning != true {
		log.Fatalln(fmt.Sprintf("Trying to stop service %s but its not running", service.Name))
	}
	scheduler.RemoveByTag(service.Name)
	service.IsRunning = false
	log.Println(fmt.Sprintf("Stopped service %s", service.Name))
}
