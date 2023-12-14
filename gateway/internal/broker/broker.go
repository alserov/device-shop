package broker

type Broker struct {
	Addr   string
	Topics *Topics
}

type Topics struct {
	Metrics *Metrics
}

type Metrics struct {
	Users   string `yaml:"users"`
	Orders  string `yaml:"orders"`
	Latency string `yaml:"latency"`
}
