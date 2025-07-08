package handlers

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	statusGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "uptime_status",
			Help: "UP = 1, DOWN = 0",
		},
		[]string{"url"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(statusGauge)
}

func UpdateStatusMetric(url string, status string) {
	if status == "UP" {
		statusGauge.WithLabelValues(url).Set(1)
	} else {
		statusGauge.WithLabelValues(url).Set(0)
	}
}

func MetricsHandler() http.Handler {
	return promhttp.Handler()
}
