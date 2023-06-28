package handlers

import (
	redisdb "news_telegram_bot/pkg/databases/redis"
	"news_telegram_bot/pkg/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func receiveLanguage(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, botState *state.BotState) {
	defer bot.Send(tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID))

	language := callbackQuery.Data
	err := redisdb.SetLanguage(callbackQuery.Message.Chat.ID, language)

	responceTest := map[string]string{
		"uk-UA": "Мова 🇺🇦 налаштована!",
		"en-US": "The language 🇺🇸 is set up!",
	}
	message := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "")
	if err == nil {
		message.Text = responceTest[language]
		defer setMenu(callbackQuery.Message.Chat.ID, language, bot)
	} else {
		message.Text = "Oops, an error occurred, try again!\n\nОй, сталася помилка, спробуй ще!"
		defer sendLanguageSwitcher(callbackQuery.Message.Chat.ID, bot)

	}
	bot.Send(message)
}

func sendLanguageSwitcher(chatID int64, bot *tgbotapi.BotAPI) {
	languageKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇺🇦", "uk-UA"),
			tgbotapi.NewInlineKeyboardButtonData("🇺🇸", "en-US"),
		),
	)

	message := tgbotapi.NewMessage(
		chatID,
		"🇺🇸Please select your preferred language\n"+
			"🇺🇦Будь ласка, оберіть бажану мову",
	)

	message.ReplyMarkup = languageKeyboard

	if _, err := bot.Send(message); err != nil {
		panic(err)
	}
}