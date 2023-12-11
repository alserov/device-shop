package request

import "github.com/prometheus/client_golang/prometheus"

type counter struct {
	m *prometheus.CounterVec
}

type Counter interface {
	Metric() prometheus.Collector
	Inc()
}

func NewCounter() Counter {
	return &counter{}
}

func (c *counter) Metric() prometheus.Collector {
	return c.m
}

func (c *counter) Inc() {
	c.m.With(nil).Inc()
}
