package gogram

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func sendTextLogic(b Bot, id int, text string) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token), nil)
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
	//req, err := http.NewRequest("GET", fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", b.Token), nil)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//q := req.URL.Query()
	//q.Add("chat_id", strconv.Itoa(id))
	//q.Add("photo", photo)
	//req.URL.RawQuery = q.Encode()
	//client := &http.Client{}
	//_, err = client.Do(req)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	data := url.Values{
		"chat_id": {strconv.Itoa(id)},
		"photo":   {photo},
	}
	_, err := http.PostForm(fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", b.Token), data)
	if err != nil {
		log.Println(err)
	}

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
