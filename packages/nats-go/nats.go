package natsGo

import (
	"errors"
	"log"
	"sync"
	"os"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

type Nats struct {
	Url         string
	ClientID    string
	ClusterID   string
	IsConnected bool
	ConnMutex   sync.Mutex
	Connection  stan.Conn
}

func NewNats(url, clientID, clusterID string) *Nats {
	return &Nats{
		Url:       url,
		ClientID:  clientID,
		ClusterID: clusterID,
	}
}

func (natClient *Nats) Disconnect() {
	natClient.ConnMutex.Lock()
	defer natClient.ConnMutex.Unlock()

	if !natClient.IsConnected {
		log.Println("Already disconnected from NATS Streaming server.")
		return
	}

	if err := natClient.Connection.Close(); err != nil {
		log.Printf("Error disconnecting from NATS Streaming server: %v", err)
	}

	natClient.IsConnected = false
	log.Println("Disconnected from NATS Streaming server.")
}

func Connect2Nats() (stan.Conn, error) {
	// Set the NATS Streaming server URI and cluster ID
	clusterID := os.Getenv("CLUSTER_ID")
	natsURL := os.Getenv("NATS_URL")
	clientID := uuid.New().String()

	if(natsURL == "" && clusterID == "") {
		log.Fatalf("Error clusterId or natUrl name is missing",)
		return nil, errors.New("Error clusterId or natUrl name is missing")
	}

	// set nats connection
	newNats := NewNats(natsURL, clientID, clusterID)
	connection, err := newNats.connect()
	return connection, err
}

func (natClient *Nats) connect() (stan.Conn, error) {
	
	natClient.ConnMutex.Lock()
	defer natClient.ConnMutex.Unlock()

	if natClient.IsConnected {
		log.Println("Already connected to NATS Streaming server.")
		return natClient.Connection, nil
	}

	sc, err := stan.Connect(natClient.ClusterID, natClient.ClientID, stan.NatsURL(natClient.Url), stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
		log.Printf("Connection lost, reason: %v", reason)
		natClient.ConnMutex.Lock()
		natClient.IsConnected = false
		natClient.ConnMutex.Unlock()
	}))

	if err != nil {
		log.Fatalf("Error connecting to NATS Streaming server: %v", err)
		return nil, err
	}

	natClient.Connection = sc
	natClient.IsConnected = true
	log.Println("Successfully connected to NATS Streaming server.")
	return natClient.Connection, nil
}
