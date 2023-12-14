package histogram

import "github.com/prometheus/client_golang/prometheus"

type history struct {
	latency *prometheus.HistogramVec
}
type History interface {
	Metrics() []prometheus.Collector
	UpdateLatency(latency float64)
}

func NewHistory() History {
	return &history{
		latency: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "latency",
			Name:      "response_time",
			Help:      "response latency",
			Buckets:   []float64{0.1, 0.15, 0.2, 0.3},
		}, []string{}),
	}
}

func (h *history) Metrics() []prometheus.Collector {
	return []prometheus.Collector{
		h.latency,
	}
}

func (h *history) UpdateLatency(latency float64) {
	h.latency.With(prometheus.Labels{}).Observe(latency)
}
