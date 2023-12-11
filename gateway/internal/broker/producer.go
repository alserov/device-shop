package broker

import "github.com/IBM/sarama"

func NewProducer(addr []string, clientID string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.ClientID = clientID

	p, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		return nil, err
	}

	return p, nil
}
