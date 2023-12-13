package broker

type Broker struct {
	Addr   string
	Topics *Topics
}

type Topics struct {
	Request *RequestTopics
}

type RequestTopics struct {
	Total      string
	Successful string
}
