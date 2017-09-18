package nsqmate

import (
	"fmt"
	"testing"
	"time"
)

func TestConf(t *testing.T) {
	config := NewConfig()
	config.MaxInFlight = 100
	config.LookupdPollInterval = time.Second * 3

	consumer, _ := NewConsumer("app3_t3", "c1", config)
	consumer.AddLookupd("127.0.0.1:4161").
		SetMsgProcessor(NewGzipMessageProcessor(nil)).
		SetHandleFunc(func(message *Message) error {
			fmt.Println("Recieve", string(message.Body))
			return nil
		})
	consumer.Run()
}
