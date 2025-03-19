package service

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type BotService struct {
}

func NewBotService(db *sql.DB, redis *redis.Client) *BotService {

	return &BotService{}
}
