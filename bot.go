package gogram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Bot represents a bot
// All fields are required
type Bot struct {
	// Token of your bot
	Token string
	// Port which server will listen to
	Port string
	// MassageHandler invokes when webhook sends a new update
	MassageHandler func(message Message)

	Self User `json:"result"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	// SupportsInlineQueries indicates if bot supports inline queries
	// This field is only for bot itself
	SupportsInlineQueries bool `json:"supports_inline_queries"`
}

// Update from webhook
type Update struct {
	Message Message `json:"message"`
}

type Message struct {
	MessageId int  `json:"message_id"`
	User      User `json:"from"`
	Chat      Chat `json:"chat"`
	Type      MassageType
	Text      string    `json:"text"`
	Animation Animation `json:"animation"`
}

type Animation struct {
	FileId string `json:"file_id"`
}

type MassageType struct {
	Text string
}

type Chat struct {
	Id int `json:"id"`
}

// NewBot creates a Bot
func NewBot(token string, port string) Bot {
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
	bot := Bot{Token: token, Port: port}
	_ = json.Unmarshal(resToByte, &bot)
	return bot
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
func (b Bot) Listener() {
	http.HandleFunc("/", handle)
	_ = http.ListenAndServe(":"+b.Port, nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	res, _ := ioutil.ReadAll(r.Body)
	update := Update{}
	err := json.Unmarshal(res, &update)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", update)
}
