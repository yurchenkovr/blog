package transport

import (
	"blog/src/infrastructure/config"
	"blog/src/infrastructure/middleware/jwt"
	"blog/src/repository/chat"
	"blog/src/repository/postgres"
	"blog/src/repository/redis"
	us "blog/src/usecases"
	"blog/src/usecases/rbac"
	"github.com/labstack/echo"
	"github.com/nats-io/nats.go"
	"log"
)

func New(cfgApi *config.APIms, cfgNats *config.NATSms) *echo.Echo {
	e := echo.New()

	dbHandler := postgres.New(cfgApi)

	jwtService := jwt.New(cfgApi)

	rdsHandler, err := redis.New()
	if err != nil {
		log.Printf("error: %v", err)
		return nil
	}

	hub := chat.NewHub()
	go hub.Run()

	rbac := rbac.Service{}

	nc, err := nats.Connect(cfgNats.NS.Url)
	if err != nil {
		log.Fatalf("error when connect to nats")
	}

	NewService(*e, us.NewArtService(postgres.NewArticleRepository(dbHandler), &rbac, nc), jwtService.MWFunc())
	NewUserService(*e, us.NewUserService(postgres.NewUserRepository(dbHandler), jwtService, &rbac, nc), jwtService.MWFunc())
	NewChatService(*e, us.NewChatService(redis.NewChatRepository(rdsHandler, "chat"), &rbac), jwtService.MWFunc(), hub)

	return e
}
