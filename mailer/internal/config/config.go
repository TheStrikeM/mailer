package config

type Config struct {
	Env   string `yaml:"env" env-default:"local"`
	Kafka Kafka  `yaml:"kafka"`
	Email Email  `yaml:"email"`
}

type Kafka struct {
	Addr                      string `yaml:"addr" env-default:"localhost:9092"`
	VerificationTopic         string `yaml:"verification_topic" env-default:"verification"`
	VerificationConsumerGroup string `yaml:"verification_consumer_group" env-default:"verification-group"`
}

type Email struct {
	From     string `yaml:"from" env-default:"thestrikem@gmail.com"`
	Password string `yaml:"password" env-default:"uhjftxclfhgwtonu"`
	Host     string `yaml:"host" env-default:"smtp.gmail.com"`
	Port     string `yaml:"port" env-default:"587"`
}
