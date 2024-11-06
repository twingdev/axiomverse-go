package networking

import (
	"github.com/nats-io/nats.go"
	"log"
)

type NATSClient struct {
	Conn *nats.Conn
	Js   nats.JetStreamContext
}

func NewNATSClient(url string) (*NATSClient, error) {
	nc, err := nats.Connect(url, nats.RetryOnFailedConnect(true))
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	client := &NATSClient{
		Conn: nc,
		Js:   js,
	}
	return client, nil
}

func (c *NATSClient) Close() {
	c.Conn.Close()
	log.Println("NATS connection closed.")
}
