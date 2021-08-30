package gogram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// bot represents a Bot
type bot struct {
	// Token of your bot
	Token string
	// MassageHandler invokes when webhook sends a new update
	MassageHandler func(message Message)
	Self           User `json:"result"`
}

// NewBot creates a bot
func NewBot(token string) bot {
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

//func SendMessage (b bot, text string) {
//	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token), nil)
//	q := req.URL.Query()
//	q.Add("chat_id", strconv.Itoa(u.Id))
//	q.Add("text", text)
//}

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

// Chat id is a unique identification number of a Telegram chat (personal or group chat).
// However, the Telegram User id is a unique identification number of a particular Telegram user.
// Use Chat id for groups, and User id for a specific user
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { webhookHandler(w, r, b) })
	_ = http.ListenAndServe(":"+port, nil)
}

func webhookHandler(w http.ResponseWriter, r *http.Request, bot bot) {
	res, _ := ioutil.ReadAll(r.Body)
	update := Update{}
	err := json.Unmarshal(res, &update)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", update)
	log.Println(string(res))
	//bot.MassageHandler(update.Message)
}
