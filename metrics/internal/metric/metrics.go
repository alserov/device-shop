package metric

import "github.com/prometheus/client_golang/prometheus"

func Setup(reg prometheus.Registerer, metrics ...prometheus.Collector) {
	reg.MustRegister(metrics...)
}
