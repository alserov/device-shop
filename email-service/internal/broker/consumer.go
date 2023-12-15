package broker

import (
	"github.com/IBM/sarama"
	"log"
)

func Subscribe(topic string, c sarama.Consumer) (<-chan []byte, error) {
	pList, err := c.Partitions(topic)
	if err != nil {
		return nil, err
	}
	offset := sarama.OffsetNewest

	chMessages := make(chan []byte, 5)
	go func() {
		defer close(chMessages)
		for _, p := range pList {
			pConsumer, err := c.ConsumePartition(topic, p, offset)
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
