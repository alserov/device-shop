package broker

import (
	"github.com/IBM/sarama"
)

func Subscribe(topic string, c sarama.Consumer) (sarama.PartitionConsumer, <-chan error) {
	var (
		chErr = make(chan error)
	)

	pList, err := c.Partitions(topic)
	if err != nil {
		chErr <- err
	}

	var pConsumer sarama.PartitionConsumer
	for _, p := range pList {
		pConsumer, err = c.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			chErr <- err
		}
	}

	return pConsumer, chErr
}
