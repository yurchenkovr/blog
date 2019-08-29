package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
)

type Logger struct {
	ID int
}

func main() {
	//Connect Options
	opts := []nats.Option{nats.Name("Logger PUB service")}

	nc, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatalf("Error.Conn PUB: %v", err)
	}
	defer nc.Close()

	testData := &Logger{
		ID: 1,
	}
	bytes, err := json.Marshal(testData)
	if err != nil {
		log.Fatalf("Error.Marshal.PUB: %v", err)
	}

	subj := "updates"
	if err := nc.Publish(subj, bytes); err != nil {
		log.Fatalf("Error.Publish: %v", err)
	}

	if err := nc.Flush(); err != nil {
		log.Fatalf("Error.PUB.Flush: %v", err)
	}

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Published [%s] : '%d'\n", subj, testData.ID)
}
