package broker

import "github.com/IBM/sarama"

func NewProducer(addrs []string, clientID string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.ClientID = clientID

	p, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return nil, err
	}

	return p, nil
}
