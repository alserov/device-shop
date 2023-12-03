package consumer

import (
	"github.com/IBM/sarama"
	"log"
)

func Subscribe(topic string, c sarama.Consumer) (<-chan string, error) {
	pList, err := c.Partitions(topic)
	if err != nil {
		return nil, err
	}
	offset := sarama.OffsetNewest

	chMessages := make(chan string)
	go func() {
		defer close(chMessages)
		for _, p := range pList {
			pConsumer, err := c.ConsumePartition(topic, p, offset)
			if err != nil {
				log.Println(err)
				return
			}
			for msg := range pConsumer.Messages() {
				chMessages <- string(msg.Value)
			}
		}
	}()

	return chMessages, nil
}
