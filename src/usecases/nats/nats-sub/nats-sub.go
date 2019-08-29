package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
	"time"
)

type Logger struct {
	ID int
}

func main() {

	//Connect options
	opts := []nats.Option{nats.Name("Logger SUB service")}
	opts = setupConnOptions(opts)

	//Connect to Nats
	nc, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	//	defer nc.Close()

	subj, i := "updates", 0

	if _, err := nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		var m Logger

		if err := json.Unmarshal(msg.Data, &m); err != nil {
			log.Fatalf("Error.Unmarshal.SUB: %v", err)
		}

		printMsg(i, msg, &m)
	}); err != nil {
		log.Fatal(err)
	}

	if err := nc.Flush(); err != nil {
		log.Fatalf("Error.Flush: %v", err)
	}

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", subj)

	runtime.Goexit()
}

func printMsg(i int, n *nats.Msg, m *Logger) {
	log.Printf("[#%d] Received on [%s]: '%d'\n", i, n.Subject, m.ID)
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.Timeout(5*time.Second))
	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))

	return opts
}
