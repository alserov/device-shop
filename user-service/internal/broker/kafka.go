package broker

type Broker struct {
	Addr   string
	Topics Topics
}

type Topics struct {
	Email string

	Worker Topic
}

type Topic struct {
	In  string
	Out string
}
