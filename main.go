package main

import (
	"flag"
	"log"
	"log/syslog"
	"os"

	"github.com/benschw/metrics-agent/metrics"
)

func main() {

	bind := flag.String("bind", "0.0.0.0:8080", "address to bind http server to")
	useSyslog := flag.Bool("syslog", false, "log to syslog")
	flag.Parse()

	if *useSyslog {
		logwriter, err := syslog.New(syslog.LOG_NOTICE, "metrics-agent")
		if err == nil {
			log.SetOutput(logwriter)
		}
	}

	log.Print("constructing service")
	svc, err := metrics.NewMetricsService(*bind)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Print("running service")
	if err := svc.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
