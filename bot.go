package gogram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Bot represents a bot. you can create multiple bots
// Token is required; but MessageHandler and Self are optional
type Bot struct {
	// Token of your Bot
	Token string
	/*
			MessageHandler invokes when webhook sends a new update.
		    In the below example, we have a Bot variable called bot.
		    We passed a function of type func (message gogram.Update, bot gogram.Bot)
			to our bot called handle.
			When telegram server sends something, handle function is called.
			Then we can use update parameter to send something back to user who sent bot a message;
			or we can use another bot.

			var bot = gogram.NewBot("<Token>", handle)
			bot.Listener(<Port>)

			func handle(update gogram.Update, bot gogram.Bot) {
				update.Message.User.SendText(bot, message.Text)
			}
	*/
	MessageHandler func(message Update, bot Bot)
	Self           User `json:"result"`
	// if Debug set to true, every time Listener receives something, it will be printed.
	Debug bool
}

// NewBot creates a Bot
func NewBot(token string, handler func(message Update, bot Bot), debug bool) Bot {
	res, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getme", token))
	if err != nil {
		log.Fatalln(err)
	}
	resToMap := map[string]interface{}{}
	resToByte, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(resToByte, &resToMap)
	if resToMap["ok"] == false {
		log.Fatalln("Your token is wrong")
	}
	var newBot = Bot{Token: token, MessageHandler: handler, Debug: debug}
	_ = json.Unmarshal(resToByte, &newBot)
	return newBot
}

// SetWebhook sets the webhook url
// Telegram server sends updates to url
func (b Bot) SetWebhook(url string) {
	_, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s", b.Token, url))
	if err != nil {
		return
	}
}

// Listener listens to upcoming webhook updates
func (b Bot) Listener(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { webhookHandler(r, b) })
	_ = http.ListenAndServe(":"+port, nil)
	return
}

func webhookHandler(r *http.Request, bot Bot) {
	res, _ := ioutil.ReadAll(r.Body)
	if bot.Debug {
		log.Println(string(res))
	}
	update := Update{}
	err := json.Unmarshal(res, &update)
	if err != nil {
		log.Fatalln(err)
	}
	if bot.MessageHandler != nil {
		bot.MessageHandler(update, bot)
	} else {
		log.Println("Warning: webhook just received something, but you have not added any handler to bot")
	}
}
