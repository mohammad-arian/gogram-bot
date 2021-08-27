package gogram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bot struct {
	Token  string
	Result Result `json:"result"`
}
type Result struct {
	Id                    int    `json:"id"`
	FirstName             string `json:"first_name"`
	Username              string `json:"username"`
	SupportsInlineQueries bool   `json:"supports_inline_queries"`
}

func NewBot(token string) Bot {
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
	bot := Bot{Token: token}
	_ = json.Unmarshal(resToByte, &bot)
	fmt.Println(bot)
	return bot
}
func (b Bot) SetWebhook(url string) {
	_, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s", b.Token, url))
	if err != nil {
		return
	}
}
