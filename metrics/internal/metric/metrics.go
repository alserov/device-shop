package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

func Setup(reg prometheus.Registerer, metrics ...[]prometheus.Collector) {
	var m []prometheus.Collector
	for _, met := range metrics {
		m = append(m, met...)
	}

	reg.MustRegister(m...)
}
