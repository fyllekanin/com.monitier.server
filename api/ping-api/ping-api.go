package ping_api

import (
	"encoding/json"
	"github.com/fyllekanin/com.monitier.server/application"
	"github.com/fyllekanin/com.monitier.server/database/repository"
	"net/http"
	"time"
)

type PingApi struct {
	application *application.Application
}

type Overview struct {
	Day                 int `json:"day"`
	AverageResponseTime int `json:"averageResponseTime"`
	DataPoints          int `json:"dataPoints"`
}

func (pingApi *PingApi) getServices(w http.ResponseWriter, r *http.Request) {
	var serviceNames []string
	for _, item := range pingApi.application.Config.Services {
		serviceNames = append(serviceNames, item.Name)
	}
	json.NewEncoder(w).Encode(serviceNames)
}

func (pingApi *PingApi) getOverviewDays(w http.ResponseWriter, r *http.Request) {
	var serviceName = r.URL.Query().Get("serviceName")
	var pingRepository = repository.NewPingRepository(pingApi.application.Db)

	pings, err := pingRepository.GetPingsSinceFor(serviceName, time.Now().Unix()-time.Now().Add(90*(-time.Hour*24)).Unix())
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(GetOverviewDays(pings))
}

func (pingApi *PingApi) getOverviewHours(w http.ResponseWriter, r *http.Request) {
	var serviceName = r.URL.Query().Get("serviceName")
	var pingRepository = repository.NewPingRepository(pingApi.application.Db)

	pings, err := pingRepository.GetPingsSinceFor(serviceName, time.Now().Unix()-time.Now().Add(24*-time.Hour).Unix())
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(GetOverviewHours(pings))
}

func (pingApi *PingApi) stopService(w http.ResponseWriter, r *http.Request) {
	var serviceName = r.URL.Query().Get("serviceName")
	for _, service := range pingApi.application.Config.Services {
		if service.Name == serviceName {
			service.Stop()
			break
		}
	}

	w.WriteHeader(200)
}

func (pingApi *PingApi) startService(w http.ResponseWriter, r *http.Request) {
	var serviceName = r.URL.Query().Get("serviceName")
	for _, service := range pingApi.application.Config.Services {
		if service.Name == serviceName {
			service.Start(repository.NewPingRepository(pingApi.application.Db))
			break
		}
	}

	w.WriteHeader(200)
}

func (pingApi *PingApi) getPings(w http.ResponseWriter, r *http.Request) {
	var pingRepository = repository.NewPingRepository(pingApi.application.Db)
	pings, err := pingRepository.GetPings()
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("gg")
		return
	}

	json.NewEncoder(w).Encode(pings)
}

func GetPingApi(application *application.Application) *PingApi {
	var api = &PingApi{
		application: application,
	}
	var subRouter = application.Router.PathPrefix("/pings").Subrouter()

	subRouter.HandleFunc("", api.getPings).Methods("GET")
	subRouter.HandleFunc("/services", api.getServices).Methods("GET")
	subRouter.HandleFunc("/overview-days", api.getOverviewDays).Methods("GET")
	subRouter.HandleFunc("/overview-hours", api.getOverviewHours).Methods("GET")
	subRouter.HandleFunc("/stop", api.stopService).Methods("GET")
	return api
}
