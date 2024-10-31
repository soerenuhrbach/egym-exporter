package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soerenuhrbach/egym-exporter/config"
	"github.com/soerenuhrbach/egym-exporter/internal/egym"
	"github.com/soerenuhrbach/egym-exporter/internal/exporter"
)

func main() {
	cfg := config.Load()

	client, err := egym.NewEgymClient(cfg.Brand, cfg.Username, cfg.Password)
	if err != nil {
		log.Fatal("Could not create egym client!")
		return
	}

	http.Handle(cfg.MetricsPath, promhttp.Handler())

	exporter := exporter.NewEgymExporter(client)
	prometheus.MustRegister(exporter)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>EGYM Exporter</title></head>
             <body>
             <h1>EGYM Exporter</h1>
             <p><a href='` + cfg.MetricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	listenAddress := fmt.Sprintf("%s:%d", cfg.BindAddress, cfg.Port)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
