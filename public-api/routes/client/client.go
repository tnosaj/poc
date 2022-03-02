package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tnosaj/poc/public-api/internals"
	"github.com/tnosaj/poc/public-api/runtime"
	"github.com/tnosaj/poc/public-api/transport"
)

type ClientApi struct {
	Metrics           internals.Metrics
	AsyncTransport    transport.AsyncTransport
	SyncTransport     transport.SyncTransport
	AsyncTransportUrl string
	SyncTransportUrl  string
}

func NewClientRoutes(runtime runtime.RuntimeSettings) *ClientApi {
	logrus.Infof("Initialize clients api")
	return &ClientApi{
		Metrics:        runtime.Metrics,
		AsyncTransport: runtime.Async,
		SyncTransport:  runtime.Sync,
	}
}

func (c *ClientApi) RegisterClientRoutes(sub *mux.Router, backends map[string]string) {
	for backend := range backends {
		switch backend {
		case "async":
			c.AsyncTransportUrl = backends[backend]
			c.registerAsyncClientRoutes(sub)

		case "sync":
			c.SyncTransportUrl = backends[backend]
			c.registerSyncClientRoutes(sub)

		}
	}
}

func (c *ClientApi) registerSyncClientRoutes(sub *mux.Router) {
	logrus.Infof("Registering SYNC client api routes with %s", c.SyncTransportUrl)
	sub.HandleFunc("/status", c.Status).Methods("GET")
	sub.HandleFunc("/new", c.NewClient).Methods("POST")
	sub.HandleFunc("/update", c.UpdateClient).Methods("POST")
	sub.HandleFunc("/delete", c.DeleteClient).Methods("POST")
	sub.HandleFunc("/get/{id}", c.GetClient).Methods("GET")
	sub.HandleFunc("/get/{id}/details", c.GetClientDetails).Methods("GET")

}

func (c *ClientApi) registerAsyncClientRoutes(sub *mux.Router) {
	logrus.Infof("Registering ASYC client api routes with %s", c.AsyncTransportUrl)
}

func (c *ClientApi) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (c *ClientApi) NewClient(w http.ResponseWriter, r *http.Request) {
	//params := r.URL.Query()
	//= params.Get("id")
	c.SyncTransport.Post(
		fmt.Sprintf("%s/new", c.SyncTransportUrl),
		[]byte{},
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (c *ClientApi) UpdateClient(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	c.SyncTransport.Post(
		fmt.Sprintf("%s/update/%s", c.SyncTransportUrl, id),
		[]byte{},
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (c *ClientApi) DeleteClient(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")

	c.SyncTransport.Post(
		fmt.Sprintf("%s/delete/%s", c.SyncTransportUrl, id),
		[]byte{},
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (c *ClientApi) GetClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c.SyncTransport.Get(
		fmt.Sprintf("%s/get/%s", c.SyncTransportUrl, vars["id"]),
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (c *ClientApi) GetClientDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c.SyncTransport.Get(
		fmt.Sprintf("%s/get/%s/details", c.SyncTransportUrl, vars["id"]),
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}
