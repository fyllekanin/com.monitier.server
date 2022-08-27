package api

import (
	"encoding/json"
	"github.com/fyllekanin/com.monitier.server/application"
	"github.com/fyllekanin/com.monitier.server/database/repository"
	"net/http"
)

type PingApi struct {
	application *application.Application
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
	subRouter.HandleFunc("/stop", api.stopService).Methods("GET")
	return api
}
