package gogram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Bot represents a bot. you can create multiple bots
// Token is required; but Handler and Self are optional
type Bot struct {
	// Token of your Bot.
	// This field is mandatory.
	Token string
	/*
			Handler invokes when webhook sends a new update.
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
	Handler func(message *Update, bot Bot)
	// Simultaneous if set to true, Handler functions run Simultaneously.
	// This field is not mandatory
	Simultaneous bool
	Proxy        *url.URL
	// Debug if set to true, every time Listener receives something, it will be printed.
	// This field is not mandatory
	Debug bool
}

func (b Bot) ActivateProxy() error {
	if b.Proxy == nil {
		return errors.New("proxy field of the bot is empty")
	}
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(b.Proxy)}
	return nil
}

// VerifyBot verifies the token
func (b *Bot) VerifyBot() error {
	res, err := request("getme", *b, nil, &UserResponse{})
	if err != nil {
		return err
	}
	getMeRes := res.(*UserResponse)
	if getMeRes.Ok != true {
		return errors.New("token is wrong")
	}
	return nil
}

// SetWebhook specifies an url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url,
// containing a JSON-serialized Update.
// In case of an unsuccessful request, we will give up after a reasonable amount of attempts.
// Returns True on success.
// If you'd like to make sure that the Webhook request comes from Telegram,
// we recommend using a secret path in the URL, e.g. https://www.example.com/<token>.
// Since nobody else knows your bot's token, you can be pretty sure it's us.
func (b Bot) SetWebhook(data SetWebhookData) (response *BooleanResponse, err error) {
	return data.Send(b)
}

func (b Bot) SetMyCommands(data SetMyCommandsData) (response *BooleanResponse, err error) {
	return data.Send(b)
}

func (b Bot) DeleteMyCommands(data DeleteMyCommandsData) (response *BooleanResponse, err error) {
	return data.Send(b)
}

func (b Bot) GetMyCommands(data GetMyCommandsData) (response *BotCommandResponse, err error) {
	return data.Send(b)
}

// Listener listens to upcoming webhook updates
func (b Bot) Listener(port string, ip ...string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { webhookHandler(r, b) })
	address := ":" + port
	if len(ip) != 0 {
		address = ip[0] + address
	}
	fmt.Println(http.ListenAndServe(address, nil))
}

func webhookHandler(r *http.Request, bot Bot) {
	res, _ := ioutil.ReadAll(r.Body)
	if bot.Debug {
		log.Println(string(res))
	}
	update := &Update{}
	err := json.Unmarshal(res, update)
	if err != nil {
		log.Println(fmt.Errorf("webhookHandler error: %w\n", err))
	}
	if bot.Handler == nil {
		log.Println("Warning: Listener just received something, but you have not added a handler to bot." +
			"add handler to bot by setting bot's Handler field to a function of type func(message Update, bot Bot) ")
	} else if bot.Simultaneous {
		// start each handler in a goroutine. since http.ListenAndServe() is a blocking function,
		// we don't have to wait for goroutines to finish.
		go bot.Handler(update, bot)
		// webhookHandler returns so telegram won't wait for response. this improves
		// performance and avoids errors such as request timeout.
		return
	} else {
		bot.Handler(update, bot)
	}
}
