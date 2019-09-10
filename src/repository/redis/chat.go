package redis

import (
	"blog/src/models"
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
)

type ChatRepository interface {
	Create(models.Chat) error
	List() (models.MessageList, error)
	SelectDB(int) error
}

type chatRepository struct {
	rds *redis.Client
	key string
}

func NewChatRepository(rds *redis.Client, key string) ChatRepository {
	return &chatRepository{rds: rds, key: key}
}

func (c chatRepository) Create(chatMsg models.Chat) error {
	bytes, err := json.Marshal(chatMsg)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	err = c.rds.ZAddNX(c.key, redis.Z{Score: float64(chatMsg.CratedAt.Unix()), Member: bytes}).Err()
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return nil
}

func (c chatRepository) List() (models.MessageList, error) {
	var msgList models.MessageList

	msg, err := c.rds.ZRange(c.key, 0, -1).Result()
	if err != nil {
		log.Printf("Error: %v", err)
		return models.MessageList{}, err
	}

	msgList.Messages = msg

	return msgList, nil
}

func (c chatRepository) SelectDB(db int) error {
	cmd := redis.NewStringCmd("select", db)

	if err := c.rds.Process(cmd); err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return nil
}
