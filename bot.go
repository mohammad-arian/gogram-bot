package gogram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// bot represents a bot
type bot struct {
	// Token of your bot
	Token string
	// MassageHandler invokes when webhook sends a new update
	MassageHandler func(message Message)
	Self           User `json:"result"`
}

// NewBot creates a bot
func NewBot(token string, port string) bot {
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
	var newBot = bot{Token: token}
	_ = json.Unmarshal(resToByte, &newBot)
	return newBot
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	// SupportsInlineQueries shows if bot supports inline queries
	// This field is only for bots
	SupportsInlineQueries bool `json:"supports_inline_queries"`
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

type Animation struct {
	FileId string `json:"file_id"`
}

type Chat struct {
	Id int `json:"id"`
}

// SetWebhook sets the webhook url
// Telegram server sends updates to url
func (b bot) SetWebhook(url string) {
	_, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s", b.Token, url))
	if err != nil {
		return
	}
}

// Listener listens to upcoming webhook updates
func (b bot) Listener(port string) {
	http.HandleFunc("/", handle)
	_ = http.ListenAndServe(":"+port, nil)
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
