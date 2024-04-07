package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := "6847886589:AAH0-9Pyo4b-8Cq92LZR3sGAAHeDYZdfcpQ"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatId := tu.ID(update.Message.Chat.ID)

		uu := update.Message.Text

		mess, err := strconv.Atoi(uu)

		if err != nil {
			return
		}

		message1 := float64(mess)
		messagemph := message1 * 1.6
		messagekmh := message1 * 0.625

		fullmessage1 := fmt.Sprintln("Kmh:", messagemph)
		fullmessage2 := fmt.Sprintln("Mph:", messagekmh)
		message := tu.Message(
			chatId,
			fullmessage1,
		)

		messagee := tu.Message(
			chatId,
			fullmessage2,
		)

		bot.SendMessage(message)
		bot.SendMessage(messagee)

	}, th.AnyMessageWithText())

	bh.Start()
}
