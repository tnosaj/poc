package session

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tnosaj/poc/public-api/internals"
	"github.com/tnosaj/poc/public-api/transport"
)

type SessionApi struct {
	Metrics        internals.Metrics
	AsyncTransport transport.AsyncTransport
	SyncTransport  transport.SyncTransport
}

func NewsessionRoutes(runtime internals.Runtime) *SessionApi {
	logrus.Infof("Initialize sessions api")
	return &SessionApi{
		Metrics:        runtime.Metrics,
		AsyncTransport: runtime.Async,
		SyncTransport:  runtime.Sync,
	}
}

func (s *SessionApi) RegisterSessionRoutes(router *mux.Router) {
	logrus.Infof("Registering session api routes")
	sub := router.PathPrefix("/session").Subrouter()
	sub.HandleFunc("/status", s.Status).Methods("GET")
	sub.HandleFunc("/new", s.NewSession).Methods("POST")
	sub.HandleFunc("/update", s.UpdateSession).Methods("POST")
	sub.HandleFunc("/end", s.EndSession).Methods("POST")
	sub.HandleFunc("/end/{key}", s.DeleteSession).Methods("POST")
	sub.HandleFunc("/get/{key}", s.GetSession).Methods("GET")
	sub.HandleFunc("/get/{key}/details", s.GetSessionDetails).Methods("GET")

}

func (s *SessionApi) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) NewSession(w http.ResponseWriter, r *http.Request) {
	s.AsyncTransport.Send()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) UpdateSession(w http.ResponseWriter, r *http.Request) {
	s.AsyncTransport.Send()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) EndSession(w http.ResponseWriter, r *http.Request) {
	s.AsyncTransport.Send()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) DeleteSession(w http.ResponseWriter, r *http.Request) {
	s.SyncTransport.Post()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) GetSession(w http.ResponseWriter, r *http.Request) {
	s.SyncTransport.Get()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}

func (s *SessionApi) GetSessionDetails(w http.ResponseWriter, r *http.Request) {
	s.SyncTransport.Get()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": "ok"}`)
}
