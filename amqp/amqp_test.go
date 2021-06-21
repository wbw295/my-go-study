package main

import (
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"testing"
	"time"
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

	consumersWg := &sync.WaitGroup{}
	times := 10
	consumersWg.Add(times)
	for i := 0; i < times; i++ {
		go func(i int) {
			defer func(i int) {
				t.Logf("consumer %d end!", i+1)
				consumersWg.Done()
			}(i)
			for d := range ds {
				t.Logf("consumer %d received msg: %s", i, string(d.Body))
				err2 := d.Ack(false)
				if err2 != nil {
					t.Error(err2)
					break
				}
			}
		}(i)
	}
	exit := make(chan struct{})
	pubWg := &sync.WaitGroup{}
	pubWg.Add(1)
	go func() {
		defer func() {
			t.Log("mq publish end!")
			pubWg.Done()
		}()
		for i := 0; i < 100; i++ {
			end := false
			select {
			case <-exit:
				t.Log("received exit signal")
				end = true
			default:
				err2 := channel.Publish("", q.Name, false, false, amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte("msg_" + strconv.Itoa(i+1)),
				})
				time.Sleep(1 * time.Second)
				if err2 != nil {
					t.Errorf("loop %d occur error: %v", i+1, err2)
					break
				}
			}
			if end {
				break
			}

		}
	}()
	//err := <-conn.NotifyClose(make(chan *amqp.Error))
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, os.Interrupt)
	s := <-signals
	t.Log(s)
	exit <- struct{}{}
	pubWg.Wait()
	_ = conn.Close()
	consumersWg.Wait()
}

func printError(err *error, t *testing.T) {
	if *err != nil {
		t.Error(*err)
	}
}
