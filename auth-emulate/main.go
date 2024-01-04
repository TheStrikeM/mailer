package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

type VerificationItem struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func main() {

	producer, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)

	item := VerificationItem{
		Email: "testthestrikem@gmail.com",
		Code:  "Hello world!",
	}

	jsonItem, _ := json.Marshal(item)
	currentMessage := &sarama.ProducerMessage{
		Topic: "verification",
		Value: sarama.StringEncoder(jsonItem),
	}

	partition, offset, _ := producer.SendMessage(currentMessage)
	fmt.Println(fmt.Sprintf("Отправлено, партиция - %d, оффсет - %d", partition, offset))
}
