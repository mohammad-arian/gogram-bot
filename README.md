# gogram-bot

**A super light-weight fast library for [Telegram bot API](https://core.telegram.org/bots/api).**


You have totally 7 files. They are very easy to pick up:

* **bot.go**: Contains Bot struct, which represents a bot and methods related to it.
* **data.go**: Contains the majority of structs that will be used to send something to telegram
(text, pic...), *I call those structs data*. Each data has two methods; Send and Check. 
Send sends the data to Request function (in utils.go) and Check will be called in Request to check
if required fields are empty or not (also might check validation of data fields).
* **utils.go**: You don't really need it. In general, it adds data fields to request, and contains a very 
important function called Request (which you also won't use directly).
* **data_test.go**: All test are here. tests must be done on a real bot token and most tests need more than
just bot token. you must provide them by flags.
* **types.go**: A long and boring file which you barely use, but it has some cool stuff like 
[Reply Keyboard](https://core.telegram.org/bots#keyboards) and 
[Inline Keyboard](https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating). 
Instead of using them directly, you will add a reply keyboard or inline keyboard to your message
through the data itself.
* **inlineMode.go**: All methods related to handling and answering 
[Inline Messages](https://core.telegram.org/bots/inline) are here.
* **passport.go**: An interface for [Telegram Passport](https://core.telegram.org/bots/api#telegram-passport).
***

## An Example:

We are going to make an echo bot.

```go
package main

import (
	"github.com/gcoder-dev/gogram-bot"
	"log"
)

var bot = gogram.Bot{Token: "Your Bot Token", Handler: handle, Concurrent: true}

func main() {
	response, err := gogram.SetWebhookData{Url: "Your webhook url"}.Send(bot)
	if err != nil {
		log.Fatalf("%+v---%+v\n", response, err)
	}
	bot.Listener("Port")
}

func handle(update gogram.Update, bot gogram.Bot) {
	response, err := gogram.TextData{Text: update.Message.Text, ChatId: update.Message.Chat.Id}.Send(bot)
	if err != nil {
		log.Fatalf("%+v---%+v\n", response, err)
	}
}
```
<br />
Brief and simple. Now Let's go through it step by step:

```go
var bot = gogram.Bot{Token: "Your Bot Token", Handler: handle, Concurrent: true}
```
First we create a bot. These are the fields:<br />
**Token:** the token you got from telegram. it is the only mandatory field. <br />
**Handler:** Every time telegram sends something to your bot, Handler will be called.
Handler is a function with handle(update gogram.Update, bot gogram.Bot) signature. <br />
**Concurrent:** When set to true, every handler will be called in a separate goroutine. you don't 
have to wait for goroutines or anything. just set Concurrent to true and let the magic happens. 
<br /><br />
```go
response, err := gogram.SetWebhookData{Url: "Your webhook url"}.Send(bot)
```
SetWebhookData is a data in data.go. Telegram sends new updates to Url. If you're using
heroku just head over to settings and use your domain for Url.
<br /><br />
```go
bot.Listener("Port")
```
When telegram send an update to your webhook Url that we sat before, you need to listen
for it. Listener is a method of Bot. It listens for upcoming updates and when received
something, it will call your Handler. Pass Listener a port (optionally you can add an IP). If 
you're using heroku simply pass Listener `os.Getenv("PORT")`.<br />
For example if your webhook url is 230.59.33.219:8004 (this is a random ip and port),
you might pass Listener "8004".
<br /><br />
Now lets see what our handler does:
```go
response, err := gogram.TextData{Text: update.Message.Text, ChatId: update.Message.Chat.Id}.Send(bot)
```
when a user sends something to your bot, it will be delivered to your
handler as an Update struct. Later we use Update to get the Text message, id of sender or many other things.
Update might contain a Message, InlineQuery, CallbackQuery or Poll struct. 
Go ahead and head over to types.go and take a look at Update and Message structs.<br >
In our handler, we create a TextData; use Update and pass Text the text user sent and id of sender to ChatId, and 
finally send it with Send method.
***
How to add [Reply Keyboard](https://core.telegram.org/bots#keyboards) and
[Inline Keyboard](https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating)
to message?
```go
func handle(update gogram.Update, bot gogram.Bot) {
    d := PhotoData{Photo: "pass a url, file_id or a file", ChatId: "a chat id"}
    err := d.SetInlineKeyboard(false, InlineButton{CallbackData: "hi", Text: "1"},
    InlineButton{Text: "Bye", CallbackData: "2"})
    if err != nil {
        return
    }
    d.Send(bot)
}
```
That was pretty much it! All data structs work the same.