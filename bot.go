package gogram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bot struct {
	Token                 string
	Port                  int
	Id                    int    `json:"id"`
	FirstName             string `json:"first_name"`
	UserName              string `json:"username"`
	SupportsInlineQueries bool   `json:"supports_inline_queries"`
}

func NewBot(token string) Bot {
	res, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getme", token))
	if err != nil {
		log.Fatalln(err)
	}
	resToMap := map[string]interface{}{}
	resToString, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(resToString, &resToMap)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resToMap)
	bot := Bot{}
	return bot
}
func (b Bot) SetWebhook(url string) {
	_, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook?url=%s", b.Token, url))
	if err != nil {
		return
	}
}
