package gogram

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
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
	Handler func(message Update, bot Bot)
	// Self is some info about bot itself.
	// This field is not mandatory
	Self User `json:"result"`
	// Simultaneous if set to true, Handler functions run Simultaneously.
	// This field is not mandatory
	Simultaneous bool
	// Debug if set to true, every time Listener receives something, it will be printed.
	// This field is not mandatory
	Debug bool
}

// NewBot creates a Bot
func NewBot(token string, handler func(message Update, bot Bot), simultaneous bool, debug bool) (Bot, error) {
	res, err := request("getme", token, nil, nil, &UserResponse{})
	if err != nil {
		return Bot{}, err
	}
	getMeRes := res.(*UserResponse)
	if getMeRes.Ok != true {
		return Bot{}, errors.New("token is wrong")
	}
	var newBot = Bot{Token: token, Handler: handler, Self: getMeRes.Result,
		Simultaneous: simultaneous, Debug: debug}
	return newBot, nil
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
	log.Printf("%+v\n", r)
	err := json.Unmarshal(res, &update)
	if err != nil {
		log.Fatalln(err)
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
