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
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
	fmt.Printf("%+v\n", send)
}

// TestAnswerInlineQueryData_Send tests if AnswerInlineQueryData work correctly. since we don't have an
// InlineQueryId until user sends an inline message, we ignore errors about InlineQueryId.
func TestAnswerInlineQueryData_Send(t *testing.T) {
	prepare()
	url1 := "https://somepic.png"
	ph1 := InlineQueryResultPhoto{Type: "photo", Id: "1", PhotoUrl: url1, ThumbUrl: url1}
	d := AnswerInlineQueryData{Results: []QueryAnswer{&ph1}, InlineQueryId: "0"}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		if send.getErrorCode() != 400 {
			t.Error(send.getDescription())
		}
	}
	fmt.Printf("%+v\n", send)
}
