package nsqmate

import (
	"github.com/nsqio/go-nsq"

	"fmt"
	"testing"
	"time"
)

func TestConf(t *testing.T) {
	config := nsq.NewConfig()
	config.MaxInFlight = 100
	config.LookupdPollInterval = time.Second * 3

	consumer, _ := NewConsumer("app1_t1", "c1", config)
	consumer.AddLookupd("127.0.0.1:4161").
		SetMsgProcessor(new(GzipMessageProcessor)).
		SetHandleFunc(func(message *nsq.Message) error {
			fmt.Println("Recieve", string(message.Body))
			return nil
		})
	consumer.Run()
}
