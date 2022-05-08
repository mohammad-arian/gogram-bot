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
	Handler func(Update, Bot)
	// if set to true, each Handler will run in a seperated goroutine.
	Concurrent bool
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

// Listener listens to upcoming webhook updates and calls webhookHandler when telegram
// sends an update.
func (b Bot) Listener(port string, ip ...string) {
	http.HandleFunc("/", b.webhookHandler)
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
func (b Bot) webhookHandler(_ http.ResponseWriter, r *http.Request) {
	res, _ := ioutil.ReadAll(r.Body)
	if b.Debug {
		log.Println(string(res))
	}
	update := &Update{}
	err := json.Unmarshal(res, update)
	if err != nil {
		log.Println(fmt.Errorf("error while unmarshaling json to Update: %w\n", err))
	}
	if b.Handler == nil {
		log.Println("Warning: Listener just received something, but you have not added a handler to bot." +
			"add handler to bot by setting bot's Handler field to a function of type func(message Update, bot Bot)")
	} else if b.Concurrent {
		go b.Handler(*update, b)
		return
	} else {
		b.Handler(*update, b)
	}
}
