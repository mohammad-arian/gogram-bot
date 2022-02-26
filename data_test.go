package gogram

import (
	"flag"
	"fmt"
	"net/url"
	"testing"
)

var ChatId *int = flag.Int("ChatId", 0, "chat id")
var Token *string = flag.String("Token", "", "token of the bot you want to use to test methods")
var Host *string = flag.String("Host", "", "ip and port to use for bot in ip:port format")
var bot *Bot

func prepare() {
	bot = &Bot{Token: *Token, Proxy: &url.URL{Host: *Host}}
	if *Host != "" {
		err := bot.ActivateProxy()
		if err != nil {
			panic(err)
		}
	}
}

func TestTextData_Send(t *testing.T) {
	prepare()
	d := TextData{Text: "Testing Text", ChatId: *ChatId}
	send, err := d.Send(*bot)
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.Error(err)
	}
	fmt.Printf("%+v\n", send)
}

func TestAnswerInlineQueryData_Send(t *testing.T) {
	prepare()
	url1 := "https://i.ibb.co/DVbYQck/Avatar-Maker.png"
	url2 := "https://i.ibb.co/NTKgQnZ/Screenshot-from-2022-02-07-11-11-58.png"
	ph1 := InlineQueryResultPhoto{Type: "photo", Id: "1", PhotoUrl: url1, ThumbUrl: url1}
	ph2 := InlineQueryResultPhoto{Type: "photo", Id: "2", PhotoUrl: url2, ThumbUrl: url2}
	d := AnswerInlineQueryData{Results: []QueryAnswer{ph1, ph2}, InlineQueryId: "21121"}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", send)
}
