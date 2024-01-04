package main

import (
	"github.com/IBM/sarama"
)

func main() {
	admin, _ := sarama.NewClusterAdmin([]string{"localhost:9092"}, nil)
	admin.CreateTopic("verification", &sarama.TopicDetail{
		NumPartitions:     3,
		ReplicationFactor: 1,
	}, false)
}
