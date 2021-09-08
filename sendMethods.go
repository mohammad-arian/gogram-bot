package gogram

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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

func sendPhotoLogic(b Bot, id int, photo interface{}) string {
	// ++++
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", b.Token),
		nil)
	client := &http.Client{}
	body := &bytes.Buffer{}
	// ++++
	switch p := photo.(type) {
	case os.File:
		w := multipart.NewWriter(body)
		field, _ := w.CreateFormField("chat_id")
		_, _ = io.Copy(field, strings.NewReader(strconv.Itoa(id)))
		file, _ := w.CreateFormFile("photo", p.Name())
		_, err = io.Copy(file, &p)
		if err != nil {
			log.Fatalln(err)
		}
		req.Body = io.NopCloser(bytes.NewReader(body.Bytes()))
		req.Header.Add("Content-Type", w.FormDataContentType())
	case string:
		q := req.URL.Query()
		q.Add("chat_id", strconv.Itoa(id))
		q.Add("photo", p)
	}
	res, err := client.Do(req)
	resToString, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
	return string(resToString)
}

func (u User) SendPhoto(b Bot, photo interface{}) {
	if u.Id == 0 {
		log.Println("User's Id field is empty")
	} else {
		sendPhotoLogic(b, u.Id, photo)
	}
}

func (c Chat) SendPhoto(b Bot, photo interface{}) {
	if c.Id == 0 {
		log.Println("Chat's Id field is empty")
	} else {
		sendPhotoLogic(b, c.Id, photo)
	}
}
