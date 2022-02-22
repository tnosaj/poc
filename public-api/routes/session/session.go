package session

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

type SessionApi struct {
	Metrics           internals.Metrics
	AsyncTransport    transport.AsyncTransport
	SyncTransport     transport.SyncTransport
	AsyncTransportUrl string
	SyncTransportUrl  string
}

func NewSessionRoutes(runtime runtime.RuntimeSettings) *SessionApi {
	logrus.Infof("Initialize sessions api")
	return &SessionApi{
		Metrics:        runtime.Metrics,
		AsyncTransport: runtime.Async,
		SyncTransport:  runtime.Sync,
	}
}

func (s *SessionApi) RegisterSessionRoutes(sub *mux.Router, backends map[string]string) {
	for backend := range backends {
		switch backend {
		case "async":
			s.AsyncTransportUrl = backends[backend]
			s.registerAsyncSessionRoutes(sub)

		case "sync":
			s.SyncTransportUrl = backends[backend]
			s.registerSyncSessionRoutes(sub)

		}
	}
}

func (s *SessionApi) registerSyncSessionRoutes(sub *mux.Router) {
	logrus.Infof("Registering SYNC session api routes with %s", s.SyncTransportUrl)
	sub.HandleFunc("/status", s.Status).Methods("GET")
	sub.HandleFunc("/new", s.NewSession).Methods("POST")
	sub.HandleFunc("/end", s.EndSession).Methods("POST")
	sub.HandleFunc("/delete/{key}", s.DeleteSession).Methods("POST")
	sub.HandleFunc("/get/{key}", s.GetSession).Methods("GET")
	sub.HandleFunc("/get/{key}/details", s.GetSessionDetails).Methods("GET")
}

func (s *SessionApi) registerAsyncSessionRoutes(sub *mux.Router) {
	logrus.Infof("Registering ASYNC session api routes with %s", s.AsyncTransportUrl)
	sub.HandleFunc("/update", s.UpdateSession).Methods("POST")
}

func (s *SessionApi) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) NewSession(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	s.SyncTransport.Post(
		fmt.Sprintf("%s/new/%s", s.SyncTransportUrl, id),
		[]byte{},
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) UpdateSession(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	s.AsyncTransport.Send(
		fmt.Sprintf("%s with data for %s", s.AsyncTransportUrl, id),
		[]byte{},
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) EndSession(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	s.SyncTransport.Post(
		fmt.Sprintf("%s/finish/%s", s.SyncTransportUrl, id),
		[]byte{},
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) DeleteSession(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	s.SyncTransport.Post(
		fmt.Sprintf("%s/delete/%s", s.SyncTransportUrl, id),
		[]byte{},
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) GetSession(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	s.SyncTransport.Get(
		fmt.Sprintf("%s/get/%s", s.SyncTransportUrl, id),
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) GetSessionDetails(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	s.SyncTransport.Get(
		fmt.Sprintf("%s/get/%s/details", s.SyncTransportUrl, id),
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}
