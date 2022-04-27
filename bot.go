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
	// Handler is invokes by webhookHandler when webhook sends a new update.
	Handler func(message *Update, bot Bot)
	// if set to true, each Handler will run in a seperated goroutine.
	concurrent bool
	// set Proxy for all connections. make
	Proxy *url.URL
	// Debug if set to true, every time Listener receives something, it will be printed.
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
func (b Bot) VerifyBot() (Response, error) {
	return Request("getme", b, nil, &ResponseImpl{Result: &User{}})
}

// SetWebhook specifies an url and receive incoming updates via an outgoing webhook.
// Whenever there is an update for the bot, we will send an HTTPS POST Request to the specified url,
// containing a JSON-serialized Update.
// In case of an unsuccessful Request, we will give up after a reasonable amount of attempts.
// Returns True on success.
// If you'd like to make sure that the Webhook Request comes from Telegram,
// we recommend using a secret path in the URL, e.g. https://www.example.com/<token>.
// Since nobody else knows your bot's token, you can be pretty sure it's us.
func (b Bot) SetWebhook(data SetWebhookData) (response Response, err error) {
	return data.Send(b)
}

func (b Bot) SetMyCommands(data SetMyCommandsData) (response Response, err error) {
	return data.Send(b)
}

func (b Bot) DeleteMyCommands(data DeleteMyCommandsData) (response Response, err error) {
	return data.Send(b)
}

func (b Bot) GetMyCommands(data GetMyCommandsData) (response Response, err error) {
	return data.Send(b)
}

// Listener listens to upcoming webhook updates and calls webhookHandler when telegram
// sends an update.
func (b Bot) Listener(port string, ip ...string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { webhookHandler(r, b) })
	address := ":" + port
	if len(ip) != 0 {
		address = ip[0] + address
	}
	log.Fatal(http.ListenAndServe(address, nil))
}

// webhookHandler is called by Bot.Listener when telegram sends an update.
// If Bot has a Handler, it will be called, otherwise a message will be printed.
// If Bot.concurrent set to true, each handler will be called in a separate goroutine.
// Since ListenAndServe function in Bot.Listener is a blocking function,
// we don't have to wait for goroutines to finish, however if http.ListenAndServe in Bot.Listener
// returns an error, all goroutines (handlers) will be aborted.
func webhookHandler(r *http.Request, bot Bot) {
	res, _ := ioutil.ReadAll(r.Body)
	if bot.Debug {
		log.Println(string(res))
	}
	update := &Update{}
	err := json.Unmarshal(res, update)
	if err != nil {
		log.Println(fmt.Errorf("error while unmarshaling json to Update: %w\n", err))
	}
	if bot.Handler == nil {
		log.Println("Warning: Listener just received something, but you have not added a handler to bot." +
			"add handler to bot by setting bot's Handler field to a function of type func(message Update, bot Bot)")
	} else if bot.concurrent {
		go bot.Handler(update, bot)
		return
	} else {
		bot.Handler(update, bot)
	}
}
