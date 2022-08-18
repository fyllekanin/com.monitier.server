package task

import (
	"github.com/fyllekanin/com.monitier.server/application"
	"github.com/fyllekanin/com.monitier.server/config"
	"github.com/fyllekanin/com.monitier.server/database/repository"
	"github.com/go-co-op/gocron"
	"net/http"
	"time"
)

func runPingFor(service config.AppService, pingRepository *repository.PingRepository) {
	var start = time.Now()
	_, err := http.Get(service.Host)
	if err != nil {
		pingRepository.InsertPing(service, -1)
	} else {
		var responseTime = int(time.Since(start).Milliseconds())
		pingRepository.InsertPing(service, responseTime)
	}
}

func StartScheduler(app *application.Application) {
	scheduler := gocron.NewScheduler(time.UTC)
	pingRepository := repository.NewPingRepository(app.Db)

	scheduler.Every(15).Seconds().Do(func() {
		for _, s := range app.Config.Services {
			go runPingFor(s, pingRepository)
		}
	})

	scheduler.StartAsync()
}
