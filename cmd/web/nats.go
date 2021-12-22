package main

import (
	"awesomeProject1/pkg/model"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
)

var Sc stan.Conn

func Stan() {
	quit := make(chan struct{})
	ConnectStan("nice-nice")
	<-quit
}

func ConnectStan(clientID string) {
	clusterID := "test-cluster"    // nats cluster id
	url := "nats://127.0.0.1:4222" // nats url

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url),
		stan.Pings(1, 3),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can't connect: %v", err)
	}

	log.Println("Connect")

	Sc = sc
}

func TakeMessage(subject, qgroup, durable string, reg *Reg) {
	mcb := func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			log.Printf("Ошибка публикации сообщения:%v", err)
		}
		var order model.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Fatalf("Invalid data, %v", err)
			return
		}
		_, err = reg.order.Create(&order)
		if err != nil {
			log.Fatalf("Error with data, %v", err)
		}
	}

	_, err := Sc.QueueSubscribe(subject,
		qgroup, mcb,
		stan.DeliverAllAvailable(),
		stan.SetManualAckMode(),
		stan.DurableName(durable))
	if err != nil {
		log.Println(err)
	}
}
