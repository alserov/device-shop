package broker

type Broker struct {
	Addr   string
	Topics *Topics
}

type Topics struct {
	Request string
}
