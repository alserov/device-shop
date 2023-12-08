package broker

type Broker struct {
	BrokerAddr string
	Topics     Topics
}

type Topics struct {
	Manager Topic
}

type Topic struct {
	In  string
	Out string
}
