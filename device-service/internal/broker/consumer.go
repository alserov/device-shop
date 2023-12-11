package broker

import (
	"github.com/IBM/sarama"
	"log"
)

func Subscribe(topic string, c sarama.Consumer) (sarama.PartitionConsumer, error) {
	pList, err := c.Partitions(topic)
	if err != nil {
		return nil, err
	}

	var pConsumer sarama.PartitionConsumer
	for _, p := range pList {
		pConsumer, err = c.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			log.Println(err)
		}
	}

	return pConsumer, nil
}
