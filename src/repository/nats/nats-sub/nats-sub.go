package nats_sub

import (
	"blog/src/infrastructure/config"
	m "blog/src/models"
	"blog/src/repository/postgres"
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"runtime"
	"time"
)

func StartServer(db *pg.DB, cfg *config.NATSms) {
	artRep := postgres.NewArticleRepository(db)

	opts := []nats.Option{nats.Name("Logger Sub Service")}
	opts = setupConnOptions(opts)

	nc, err := nats.Connect(cfg.NS.Url, opts...)
	if err != nil {
		log.Fatalf("Error.Conn PUB: %v", err)
	}

	if _, err := nc.Subscribe(cfg.NS.Subj, func(msg *nats.Msg) {
		var l m.Logger

		if err := json.Unmarshal(msg.Data, &l); err != nil {
			log.Fatalf("Error.Unmarshal.SUB: %v", err)
		}

		var loggs string

		if l.Method == "DELETE" {
			loggs = fmt.Sprintf("[%s] on [%s]: \n\tID:'%d'\n", l.Method, msg.Subject, l.ID)
		} else {
			article, err := GetArticle(artRep, l.ID)
			if err != nil {
				log.Fatalf("Error.SUB.DB: %v", err)
			}

			loggs = fmt.Sprintf("[%s] on [%s]: \n\tID:'%d'\n\tTitle: '%s'\n\tAuthor: '%s'\n", l.Method, msg.Subject, article.ID, article.Title, article.Username)
		}

		writeFileArt(loggs)

	}); err != nil {
		log.Fatal(err)
	}
	if err := nc.Flush(); err != nil {
		log.Fatalf("Error.SUB.Flush: %v", err)
	}

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	runtime.Goexit()
}

func writeFileArt(loggs string) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	if _, err := file.WriteString(loggs); err != nil {
		log.Fatalf("Error.SUB.WriteString: %v", err)
	}
}

func GetArticle(db postgres.ArticleRepository, id int) (m.Article, error) {
	article, err := db.View(id)
	if err != nil {
		return m.Article{}, err
	}

	return article, nil
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
