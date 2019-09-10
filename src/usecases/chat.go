package usecases

import (
	"blog/src/models"
	"blog/src/repository/redis"
	"blog/src/usecases/rbac"
	"errors"
	"github.com/labstack/echo"
	"log"
	"time"
)

type ChatService interface {
	Create(echo.Context, CreateReqChat) error
	List() (models.MessageList, error)
	SelectDB(int) error
}

type chatService struct {
	chatRep redis.ChatRepository
	rbac    rbac.RBAC
}

type CreateReqChat struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func NewChatService(chatRep redis.ChatRepository, rbac rbac.RBAC) ChatService {
	return &chatService{chatRep: chatRep, rbac: rbac}
}

func (c chatService) Create(ctx echo.Context, req CreateReqChat) error {
	if c.rbac.IsBlocked(ctx) {
		return errors.New("Your user is blocked.\n")
	}

	chatMsg := models.Chat{
		Username: req.Username,
		Message:  req.Message,
		CratedAt: time.Time{},
	}

	if err := c.chatRep.Create(chatMsg); err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return nil
}

func (c chatService) List() (models.MessageList, error) {
	msgList, err := c.chatRep.List()
	if err != nil {
		log.Printf("Error: %v", err)
		return models.MessageList{}, err
	}

	return msgList, nil
}

func (c chatService) SelectDB(db int) error {
	if err := c.chatRep.SelectDB(db); err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return nil
}
