package natsGo

import (
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

// Define a type for a function that takes two integers and returns an integer
type OnMessage func(msg *stan.Msg)

type ListenerNats struct {
	Subject        string
	QueueGroupName string
	AckWait        int
	connection     stan.Conn
	GroupName      string
}

// NewListenerNats creates a new instance of ListenerNats
func NewListenerNats(connection stan.Conn, subject, groupName string) *ListenerNats {
	return &ListenerNats{
		connection: connection,
		Subject:    subject,
		GroupName:  groupName,
	}
}

func (service *ListenerNats) Listen(connection stan.Conn, subject, queueGroupName string, onMessage OnMessage) {
	fmt.Printf("listening on subject: %s / groupName: %s\n", subject, queueGroupName)

	if subject == "" {
		log.Fatalf("Error no subject to channel")
	}

	// Set subscription options
	durableName := stan.DurableName(subject)
	manualAckMode := stan.SetManualAckMode()
	deliverAllAvailable := stan.DeliverAllAvailable()

	subscriptionOptions := []stan.SubscriptionOption{
		durableName,
		manualAckMode,
		deliverAllAvailable,
	}

	_, err := connection.QueueSubscribe(subject, queueGroupName, func(msg *stan.Msg) {
		onMessage(msg)
	}, subscriptionOptions...)
	if err != nil {
		log.Fatalf("Error subscribing to channel: %v", err)
	}
}


