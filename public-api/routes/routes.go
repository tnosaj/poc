package routes

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tnosaj/poc/public-api/internals"
	"github.com/tnosaj/poc/public-api/routes/client"
	"github.com/tnosaj/poc/public-api/routes/session"
)

type Routers struct {
	Client  client.ClientApi
	Session session.SessionApi
}

func InitializeNewRoutes(runtime internals.Runtime) *Routers {
	return &Routers{
		Client:  *client.NewClientRoutes(runtime),
		Session: *session.NewsessionRoutes(runtime),
	}

}

func (r *Routers) RegisterRoutes(router *mux.Router) {
	logrus.Infof("Registering all api routes")
	r.Session.RegisterSessionRoutes(router)
	r.Client.RegisterClientRoutes(router)
}
