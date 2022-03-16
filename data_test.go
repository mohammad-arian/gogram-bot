package gogram

import (
	"flag"
	"net/url"
	"testing"
)

// Some tests might need a ChatId or bot Token; set them as flags (e.g.
// go test -run TestKeyboard -ChatId=<chat id> -Token=<bot token>)
var ChatId *int = flag.Int("ChatId", 0, "chat id")
var UserId *int = flag.Int("UserId", 0, "user id")
var MessageId *int = flag.Int("MessageId", 0, "message id")
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
}

func TestTextData_Send_ParseMode(t *testing.T) {
	prepare()
	text := `<b>bold</b>, <strong>bold</strong>
		<i>italic</i>, <em>italic</em>
		<u>underline</u>, <ins>underline</ins>
		<s>strikethrough</s>, <strike>strikethrough</strike>, <del>strikethrough</del>
		<b>bold <i>italic bold <s>italic bold strikethrough</s> <u>underline italic bold</u></i> bold</b>
		<a href="https://www.example.com/">inline URL</a>
		<a href="tg://user?id=123456789">inline mention of a user</a>
		<code>inline fixed-width code</code>
		<pre>pre-formatted fixed-width code block</pre>
		<pre><code class="language-python">pre-formatted fixed-width code block written in the Python programming language</code></pre>`
	d := TextData{Text: text, ChatId: *ChatId, ParseMode: "HTML"}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
}

func TestTextData_Send_WrongParseMode(t *testing.T) {
	prepare()
	text := `<b>bold</b>, <strong>bold</strong>
		<i>italic</i>, <em>italic</em>
		<u>underline</u>, <ins>underline</ins>
		<s>strikethrough</s>, <strike>strikethrough</strike>, <del>strikethrough</del>
		<b>bold <i>italic bold <s>italic bold strikethrough</s> <u>underline italic bold</u></i> bold</b>
		<a href="https://www.example.com/">inline URL</a>
		<a href="tg://user?id=123456789">inline mention of a user</a>
		<code>inline fixed-width code</code>
		<pre>pre-formatted fixed-width code block</pre>
		<pre><code class="language-python">pre-formatted fixed-width code block written in the Python programming language</code></pre>`
	d := TextData{Text: text, ChatId: *ChatId, ParseMode: "i am wrong"}
	_, err := d.Send(*bot)
	if err == nil {
		t.Error(err)
	}
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
}

func TestKeyboard2(t *testing.T) {
	prepare()
	d := TextData{Text: "Testing Text with Keyboard", ChatId: *ChatId}
	err := d.SetReplyKeyboard(ReplyKeyboardOP{}, ReplyButton{Text: "A"},
		ReplyButton{Text: "B"})
	if err != nil {
		t.Error(err)
	}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
}

func TestSendChatActionData_Send(t *testing.T) {
	prepare()
	d := SendChatActionData{Action: "upload_photo", ChatId: *ChatId}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
}

func TestSendChatActionData_Send_WrongAction(t *testing.T) {
	prepare()
	d := SendChatActionData{Action: "WrongAction", ChatId: *ChatId}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
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
}

func TestLocationData_Send(t *testing.T) {
	prepare()
	d := LocationData{ChatId: *ChatId, Location: Location{Latitude: 51.165691, Longitude: 10.451526}}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
}

func TestPollData_Send(t *testing.T) {
	prepare()
	d := PollData{ChatId: *ChatId, Question: "This is a poll test", Options: []string{"1", "2", "3"}}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
}

func TestContactData_Send(t *testing.T) {
	prepare()
	d := ContactData{ChatId: *ChatId, Contact: Contact{PhoneNumber: "00", FirstName: "TestUser"}}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
}

func TestGetChatAdministratorsData_Send(t *testing.T) {
	prepare()
	d := GetChatAdministratorsData{ChatId: *ChatId}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
}

func TestGetGameHighScoresData_Send(t *testing.T) {
	prepare()
	d := GetGameHighScoresData{ChatId: -1001584537184, UserId: 1141922418, MessageId: 285}
	send, err := d.Send(*bot)
	if err != nil {
		t.Error(err)
	} else if send.isOk() == false {
		t.Error(send.getDescription())
	}
}
