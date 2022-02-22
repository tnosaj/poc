package routes

import (
	"github.com/gorilla/mux"
	"github.com/tnosaj/poc/public-api/routes/client"
	"github.com/tnosaj/poc/public-api/routes/session"
	"github.com/tnosaj/poc/public-api/runtime"
)

type Routers struct {
	Client  client.ClientApi
	Session session.SessionApi
}

func InitializeNewRoutes(runtime runtime.RuntimeSettings, baseurls map[string]map[string]string, router *mux.Router) {

	r := Routers{}

	for backend := range baseurls {
		switch backend {
		case "client":
			r.Client = *client.NewClientRoutes(runtime)
			r.Client.RegisterClientRoutes(
				router.PathPrefix("/client").Subrouter(),
				baseurls["client"],
			)

		case "session":
			r.Session = *session.NewSessionRoutes(runtime)
			r.Session.RegisterSessionRoutes(
				router.PathPrefix("/session").Subrouter(),
				baseurls["session"],
			)
		}
	}

}
