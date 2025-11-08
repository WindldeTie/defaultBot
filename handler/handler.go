package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type Handler struct {
	bot *tgbotapi.BotAPI
}

func NewHandler(bot *tgbotapi.BotAPI) *Handler {
	return &Handler{
		bot: bot,
	}
}

func (h *Handler) Start(debug bool) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	h.bot.Debug = debug
	updates := h.bot.GetUpdatesChan(u)
	go h.console()

	for update := range updates {
		h.HandleUpdate(update)
	}
}

// Обработка команд --------------------------------------------------------------------------------------------------

func (h *Handler) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		command := strings.TrimSpace(update.Message.Text)
		msgArr := strings.Split(command, " ")
		switch msgArr[0] {
		case "/start":
			h.handleStart(update)
			return
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			h.bot.Send(msg)
		}
	}
}

func (h *Handler) handleStart(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!")
	h.bot.Send(msg)
}

// Вспомогательные функции --------------------------------------------------------------------------------------------
