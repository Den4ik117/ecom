package main

import (
	"fmt"
	"github.com/Den4ik117/ecom/cmd/api"
	"github.com/Den4ik117/ecom/config"
	"github.com/Den4ik117/ecom/db"
	"github.com/go-sql-driver/mysql"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	database, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	} else {
		log.Println("Successfully connected to database")
	}

	rabbitmq, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.Envs.RabbitMQUser,
		config.Envs.RabbitMQPassword,
		config.Envs.RabbitMQHost,
		config.Envs.RabbitMQPort,
	))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer rabbitmq.Close()

	ch, err := rabbitmq.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel. Error: %s", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"ecom", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue. Error: %s", err)
	}

	server := api.NewApiServer(":8080", database, ch)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
