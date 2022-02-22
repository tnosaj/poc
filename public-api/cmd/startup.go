package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tnosaj/poc/public-api/internals"
)

func evaluateInputs() (internals.Settings, error) {
	var s internals.Settings

	flag.BoolVar(&s.Debug, "v", false, "Enable verbose debugging output")

	flag.StringVar(&s.AsyncTransportSettings.Name, "async", "nullqueue", "Use this queue")
	flag.StringVar(&s.Port, "p", "8080", "Starts server on this port")
	flag.StringVar(&s.SyncTransportSettings.Name, "sync", "nullhttp", "Use this http transport")

	var tmpString string
	flag.StringVar(&tmpString, "baseurls", `{"client":{"sync":"syncClientHost"},"session":{"sync":"syncSessionHost","async":"asyncSessionHost"}}`, "Json format of urls")
	json.Unmarshal([]byte(tmpString), &s.Backends)

	flag.IntVar(&s.SyncTransportSettings.Timeout, "t", 1, "Timeout in seconds for a backend answer")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: [flags] command [command args…]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	setupLogger(s.Debug)

	acceptedQueuess := map[string]bool{
		"kafka":     true,
		"nullqueue": true,
	}
	if !acceptedQueuess[s.AsyncTransportSettings.Name] {
		return internals.Settings{}, fmt.Errorf("Unknown queue engine specified: %q", s.AsyncTransportSettings)
	}

	return s, nil
}

func setupLogger(debug bool) {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.Debug("Configured logger")
}
