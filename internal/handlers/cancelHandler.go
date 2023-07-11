package handlers

import (
	redisdb "news_telegram_bot/pkg/databases/redis"
	"news_telegram_bot/pkg/translator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cancel(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	redisdb.DoneChatState(callbackQuery.Message.Chat.ID)

	lang, err := redisdb.GetLanguage(callbackQuery.Message.Chat.ID)
	if err != nil {
		return // TODO!!!!!!!
	}
	printer := translator.GetPrinter(lang)

	responceMessage := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, printer.Sprintf("Дію скасовано..."))
	bot.Send(responceMessage)
}
