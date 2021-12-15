package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/tnosaj/poc/public-api/internals"
	"github.com/tnosaj/poc/public-api/metrics"
	"github.com/tnosaj/poc/public-api/routes"
)

func main() {
	var wait time.Duration
	s, err := evaluateInputs()
	if err != nil {
		log.Fatalf("could not evaluate inputs: %q", err)
	}

	runtime := internals.NewRuntime(
		s.AsyncTransport,
		s.SyncTransport,
		metrics.RegisterPrometheusMetrics(),
	)

	apiRouter := routes.InitializeNewRoutes(runtime)

	router := mux.NewRouter()
	apiRouter.RegisterRoutes(router)
	router.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", s.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * time.Duration(s.Timeout),
		ReadTimeout:  time.Second * time.Duration(s.Timeout),
		IdleTimeout:  time.Second * time.Duration(s.Timeout),
		Handler:      router,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

}
