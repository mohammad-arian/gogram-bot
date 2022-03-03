package gogram

import (
	"flag"
	"fmt"
	"net/url"
	"testing"
)

// Some tests might need a ChatId or bot Token; set them as flags (e.g.
// go test -run TestKeyboard -ChatId=<a chat id> -Token=<a bot token>)
var ChatId *int = flag.Int("ChatId", 0, "chat id")
var Token *string = flag.String("Token", "", "token of the bot you want to use to test methods")
var Host *string = flag.String("Host", "", "ip and port to use for bot in ip:port format")
var bot *Bot

// prepare creates a bot by passed flags and activated proxy if there is a -Host flag in ip:port format
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

// TestTextData_Send_EmptyText tests TextData.Check in case Text is empty
func TestTextData_Send_EmptyText(t *testing.T) {
	prepare()
	d := TextData{ChatId: *ChatId}
	_, err := d.Send(*bot)
	if err == nil {
		t.Error("check is broken")
	}
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

func TestKeyboard(t *testing.T) {
	prepare()
	d := TextData{Text: "Testing Text with Keyboard", ChatId: *ChatId}
	err := d.SetInlineKeyboard(false, InlineButton{CallbackData: "hi", Text: "1"},
		InlineButton{Text: "Bye", CallbackData: "2"})
	if err != nil {
		t.Error(err)
	}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
	fmt.Printf("%+v\n", send)
}

func TestDiceData_Send(t *testing.T) {
	prepare()
	d := DiceData{Emoji: "🎲", ChatId: *ChatId}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
	fmt.Printf("%+v\n", send)
}
