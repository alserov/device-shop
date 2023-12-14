package broker

type Broker struct {
	Addr   string
	Topics *Topics
}

type Topics struct {
	User    string
	Order   string
	Latency string
}
