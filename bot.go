package gogram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bot struct {
	Token string
	Port  string
	User  User `json:"result"`
}

type User struct {
	Id                    int    `json:"id"`
	FirstName             string `json:"first_name"`
	Username              string `json:"username"`
	SupportsInlineQueries bool   `json:"supports_inline_queries"`
}

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

func (b Bot) Listener() {
	http.HandleFunc("/", handle)
	_ = http.ListenAndServe(":"+b.Port, nil)
}
func handle(w http.ResponseWriter, r *http.Request) {
	res, _ := ioutil.ReadAll(r.Body)
	log.Println(string(res))
}
