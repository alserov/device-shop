package counter

import (
	"github.com/prometheus/client_golang/prometheus"
)

type counter struct {
	users *prometheus.CounterVec
}

type Counter interface {
	Metrics() []prometheus.Collector
	IncUsers()
}

func NewCounter() Counter {
	return &counter{
		users: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "users",
			Name:      "users_amount",
			Help:      "total amount of requests",
		}, []string{}),
	}
}

func (c *counter) Metrics() []prometheus.Collector {
	return []prometheus.Collector{c.users}
}

func (c *counter) IncUsers() {
	c.users.With(prometheus.Labels{}).Inc()
}
