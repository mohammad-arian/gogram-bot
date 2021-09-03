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
		    It must have two parameters, one of type Message
		    the other of type Bot.
			In the below example, we have a Bot variable called bot.
			We passed a function of type func (message gogram.Message, bot gogram.Bot)
			to our bot called handle.
			When telegram server sends something, handle function is invoked.
			Then we can use bot parameter to send something back to user who sent bot message;
			or we can use another bot.

			var bot = gogram.NewBot("<Token>", handle)
			bot.Listener(<Port>)

			func handle(message gogram.Message, bot gogram.Bot) {
				message.User.SendMessageToUser(bot, message.Text)
			}
	*/
	MessageHandler func(message Message, bot Bot)
	Self           User `json:"result"`
}

// NewBot creates a Bot
func NewBot(token string, handler func(message Message, bot Bot)) Bot {
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
	var newBot = Bot{Token: token, MessageHandler: handler}
	_ = json.Unmarshal(resToByte, &newBot)
	return newBot
}

// Update from webhook
type Update struct {
	Message Message `json:"message"`
}

type Message struct {
	MessageId int       `json:"message_id"`
	User      User      `json:"from"`
	Chat      Chat      `json:"chat"`
	Text      string    `json:"text"`
	Animation Animation `json:"animation"`
}

type User struct {
	// Chat id is a unique identification number of a Telegram chat (personal or group chat).
	// However, the Telegram User id is a unique identification number of a particular Telegram user.
	// Use Chat id for groups, and User id for a specific user
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	// SupportsInlineQueries shows if Bot supports inline queries
	// This field is only for bots
	SupportsInlineQueries bool `json:"supports_inline_queries"`
}

// Chat id is a unique identification number of a Telegram chat (personal or group chat).
// However, the Telegram User id is a unique identification number of a particular Telegram user.
// Use Chat id for groups, and User id for a specific user
type Chat struct {
	Id int `json:"id"`
}

type Animation struct {
	FileId string `json:"file_id"`
}

// TypeIndicator function returns the type of message
// This make it easier to know which fields are empty and which aren't
// TypeIndicator may return "Text", "Animation" and etc
func (m Message) TypeIndicator() string {
	switch {
	case m.Text != "":
		return "Text"
	case m.Animation != Animation{}:
		return "Animation"
	default:
		return "Unknown"
	}
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { webhookHandler(w, r, b) })
	_ = http.ListenAndServe(":"+port, nil)
}

func webhookHandler(w http.ResponseWriter, r *http.Request, bot Bot) {
	res, _ := ioutil.ReadAll(r.Body)
	update := Update{}
	err := json.Unmarshal(res, &update)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", update)
	log.Println(string(res))
	if bot.MessageHandler != nil {
		bot.MessageHandler(update.Message, bot)
	} else {
		log.Println("Warning: webhook just received something, but you have not added any handler to bot")
	}
}
