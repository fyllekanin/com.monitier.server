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

func GetApi(application *application.Application) *PingApi {
	var api = &PingApi{
		application: application,
	}
	var subRouter = application.Router.PathPrefix("/pings").Subrouter()

	subRouter.HandleFunc("", api.getPings).Methods("GET")
	return api
}
