package broker

type Broker struct {
	Addr   string
	Topics *Topics
}

type Topics struct {
	Email string

	Worker WorkerTopic
}

type WorkerTopic struct {
	In  string
	Out string
}
