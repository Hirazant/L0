package nats

import (
	"awesomeProject1/pkg/model"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

var Sc stan.Conn

func Stan() {
	quit := make(chan struct{})
	ConnectStan("nice-nice")
	//PrintMessage("test", "test", "test-1")
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
		log.Fatalf("Не могу подключится: %v.\nПроверьте запущен ли сервер NATS на: %s", err, url)
	}

	log.Println("Подключился")

	Sc = sc
}

func TakeMessage(subject, qgroup, durable string) {
	mcb := func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			log.Printf("Ошибка публикации сообщения:%v", err)
		}
		var order model.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			fmt.Println("not good ", order)
			panic(err)
		}
		fmt.Println("all good -  ", order)
		//	fmt.Println(string(msg.Data))
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
