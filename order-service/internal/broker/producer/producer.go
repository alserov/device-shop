package producer

import "github.com/IBM/sarama"

func NewProducer(brokers []string, clientID string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.ClientID = clientID

	p, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func NewAsyncProducer(brokers []string, clientID string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.ClientID = clientID
	config.Producer.RequiredAcks = sarama.WaitForAll

	p, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return p, nil
}
