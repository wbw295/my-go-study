package main

import (
	"github.com/streadway/amqp"
	"strconv"
	"sync"
	"testing"
)

func TestQueueDeclare(t *testing.T) {
	var err error
	defer printError(&err, t)
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return
	}
	channel, err := conn.Channel()
	if err != nil {
		return
	}
	q, err := channel.QueueDeclare("test_queue", true, false, false, false, nil)
	if err != nil {
		return
	}

	ds, err := channel.Consume(q.Name, "test", false, false, false, false, nil)
	if err != nil {
		return
	}

	wg := &sync.WaitGroup{}
	times := 10
	wg.Add(times + 1)
	for i := 0; i < times; i++ {
		go func() {
			for d := range ds {
				t.Logf("received msg: %s", string(d.Body))
				err2 := d.Nack(false, true)
				if err2 != nil {
					t.Error(err2)
					break
				}
			}
			wg.Done()
		}()
	}

	go func() {
		for i := 0; i < 10000; i++ {
			err2 := channel.Publish("", q.Name, false, false, amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte("msg_" + strconv.Itoa(i+1)),
			})
			if err2 != nil {
				t.Errorf("loop %d occur error: %v", i+1, err2)
				break
			}
		}
		wg.Done()
	}()
	t.Log(q)
	wg.Wait()
}

func printError(err *error, t *testing.T) {
	if *err != nil {
		t.Error(*err)
	}
}
