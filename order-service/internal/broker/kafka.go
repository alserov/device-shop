package broker

type Broker struct {
	BrokerAddr string
	Topics     Topics
}

type Topics struct {
	Email string

	User Topic

	Device Topic

	Collection Topic
}

type Topic struct {
	In  string
	Out string
}
