package gogram

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Bot represents a bot. you can create multiple bots
// Token is required; but MessageHandler and Self are optional
type Bot struct {
	// Token of your Bot
	Token string
	/*
			MessageHandler invokes when webhook sends a new update.
		    In the below example, we have a Bot called bot.
		    We passed a function of type func (message gogram.Update, bot gogram.Bot)
			to our bot called handle.
			When telegram server sends something, this function is called.
			Then we can use update.Message.User to send something back to user;

			// create bot
			var bot, _ = gogram.NewBot("<Token>", handle)
			// start  listening to telegram
			bot.Listener(<Port>)
			// handler function
			func handle(update gogram.Update, bot gogram.Bot) {
				update.Message.User.SendText(bot, update.Message.Text, nil)
			}
	*/
	MessageHandler func(message Update, bot Bot)
	Self           User `json:"result"`
	// if Debug set to true, every time Listener receives something, it will be printed.
	Debug bool
}

// NewBot creates a Bot
func NewBot(token string, handler func(message Update, bot Bot), debug bool) (Bot, error) {
	res, err := request("getme", token, nil, nil, &UserResponse{})
	if err != nil {
		return Bot{}, err
	}
	getMeRes := res.(*UserResponse)
	if getMeRes.Ok != true {
		return Bot{}, errors.New("token is wrong")
	}
	var newBot = Bot{Token: token, MessageHandler: handler, Self: getMeRes.Result, Debug: debug}
	return newBot, nil
}

// SetWebhook specifies an url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
// containing a JSON-serialized Update.
// In case of an unsuccessful request, we will give up after a reasonable amount of attempts.
// Returns True on success.
// If you'd like to make sure that the Webhook request comes from Telegram,
// we recommend using a secret path in the URL, e.g. https://www.example.com/<token>.
// Since nobody else knows your bot's token, you can be pretty sure it's us.
func (b Bot) SetWebhook(url string, optionalParams *SetWebhookOptionalParams) (response *BooleanResponse, err error) {
	type data struct {
		Url string `json:"url"`
	}
	d := data{Url: url}
	res, err := request("setWebhook", b.Token, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

// Listener listens to upcoming webhook updates
func (b Bot) Listener(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { webhookHandler(r, b) })
	_ = http.ListenAndServe(":"+port, nil)
}

func webhookHandler(r *http.Request, bot Bot) {
	res, _ := ioutil.ReadAll(r.Body)
	if bot.Debug {
		log.Println(string(res))
	}
	update := Update{}
	err := json.Unmarshal(res, &update)
	if err != nil {
		log.Fatalln(err)
	}
	if bot.MessageHandler != nil {
		bot.MessageHandler(update, bot)
	} else {
		log.Println("Warning: Listener just received something, but you have not added any handler to bot")
	}
}
