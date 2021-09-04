package gogram

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func sendTextLogic(b Bot, id int, text string) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token), nil)
	if err != nil {
		log.Fatalln(err)
	}
	q := req.URL.Query()
	q.Add("chat_id", strconv.Itoa(id))
	q.Add("text", text)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
}

// SendText sends message to a User.
// b Bot parameter indicated which bot to send
// the message with. This way you can send messages
// with different bots
func (u User) SendText(b Bot, text string) {
	if u.Id == 0 {
		log.Println("User's Id field is empty")
	} else {
		sendTextLogic(b, u.Id, text)
	}
}

// SendText sends message to a Chat.
// b Bot parameter indicated which bot to send
// the message with. This way you can send messages
// with different bots
func (c Chat) SendText(b Bot, text string) {
	if c.Id == 0 {
		log.Println("User's Id field is empty")
	} else {
		sendTextLogic(b, c.Id, text)
	}
}

func sendPhotoLogic(b Bot, id int, photo string) {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(id))
	q.Add("photo", photo)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", b.Token),
		strings.NewReader(q.Encode()))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "multipart/form-data")
	//q := req.URL.Query()

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}

func (u User) SendPhoto(b Bot, photo string) {
	if u.Id == 0 {
		log.Println("User's Id field is empty")
	} else {
		sendPhotoLogic(b, u.Id, photo)
	}
}

func (c Chat) SendPhoto(b Bot, photo string) {
	if c.Id == 0 {
		log.Println("Chat's Id field is empty")
	} else {
		sendPhotoLogic(b, c.Id, photo)
	}
}
