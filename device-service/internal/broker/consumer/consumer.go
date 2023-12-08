package consumer

import (
	"github.com/IBM/sarama"
	"log"
)

func Subscribe(topic string, c sarama.Consumer) (<-chan []byte, error) {
	pList, err := c.Partitions(topic)
	if err != nil {
		return nil, err
	}

	chMessages := make(chan []byte)
	go func() {
		defer close(chMessages)
		for _, p := range pList {
			pConsumer, err := c.ConsumePartition(topic, p, sarama.OffsetNewest)
			if err != nil {
				log.Println(err)
				return
			}
			for msg := range pConsumer.Messages() {
				chMessages <- msg.Value
			}
		}
	}()

	return chMessages, nil
}
