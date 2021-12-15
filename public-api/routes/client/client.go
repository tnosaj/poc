package client

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tnosaj/poc/public-api/internals"
	"github.com/tnosaj/poc/public-api/transport"
)

type ClientApi struct {
	Metrics        internals.Metrics
	AsyncTransport transport.AsyncTransport
	SyncTransport  transport.SyncTransport
}

func NewClientRoutes(runtime internals.Runtime) *ClientApi {
	logrus.Infof("Initialize clients api")
	return &ClientApi{
		Metrics:        runtime.Metrics,
		AsyncTransport: runtime.Async,
		SyncTransport:  runtime.Sync,
	}
}

func (c *ClientApi) RegisterClientRoutes(router *mux.Router) {
	logrus.Infof("Registering client api routes")
	sub := router.PathPrefix("/client").Subrouter()
	sub.HandleFunc("/status", c.Status).Methods("GET")
	sub.HandleFunc("/new", c.NewClient).Methods("POST")
	sub.HandleFunc("/update", c.UpdateClient).Methods("POST")
	sub.HandleFunc("/delete", c.DeleteClient).Methods("POST")
	sub.HandleFunc("/get/{key}", c.GetClient).Methods("GET")
	sub.HandleFunc("/get/{key}/details", c.GetClientDetails).Methods("GET")

}

func (c *ClientApi) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (c *ClientApi) NewClient(w http.ResponseWriter, r *http.Request) {
	c.SyncTransport.Post()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (c *ClientApi) UpdateClient(w http.ResponseWriter, r *http.Request) {
	c.SyncTransport.Post()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (c *ClientApi) DeleteClient(w http.ResponseWriter, r *http.Request) {
	c.SyncTransport.Post()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (c *ClientApi) GetClient(w http.ResponseWriter, r *http.Request) {
	c.SyncTransport.Get()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (c *ClientApi) GetClientDetails(w http.ResponseWriter, r *http.Request) {
	c.SyncTransport.Get()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}
