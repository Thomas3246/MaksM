package handler

import (
	"github.com/Thomas3246/MaksM/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler struct {
	api     *tgbotapi.BotAPI
	service *service.BotService
}

func NewBotHandler(api *tgbotapi.BotAPI, service *service.BotService) *BotHandler {
	return &BotHandler{
		api:     api,
		service: service}
}

func (h *BotHandler) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := h.api.GetUpdatesChan(u)

	for update := range updates {
		go h.HandleMessage(&update)
	}
}

func (h *BotHandler) HandleMessage(update *tgbotapi.Update) {

	if update.Message != nil {
	}

}
