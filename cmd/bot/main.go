package main

import (
	"log"

	"github.com/Thomas3246/MaksM/internal/bot"
	"github.com/Thomas3246/MaksM/internal/config"
	"github.com/Thomas3246/MaksM/internal/repository/sqlite"
)

func main() {

	cfg, err := config.NewConfig("../../config/config.json")
	if err != nil {
		log.Fatalf("Ошибка при инициализации конфиг-файла: %v", err)
	}
	if cfg.BotToken == "" {
		log.Fatalf("API-токен не введён")
	}

	db, err := sqlite.InitDB()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	bot, err := bot.NewBot(cfg, db)
	if err != nil {
		log.Fatalf("Ошибка инициализации приложения: %v", err)
	}

	err = bot.Start()
	if err != nil {
		log.Fatalf("Ошибка запуска бота: %v", err)
	}
}
