package gogram

import (
	"bytes"
	"errors"
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

func sendPhotoLogic(b Bot, id int, photo interface{}) (string, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", b.Token),
		nil)
	switch p := photo.(type) {
	case *os.File:
		var body = &bytes.Buffer{}
		w := multipart.NewWriter(body)
		chatId, err := w.CreateFormField("chat_id")
		if err != nil {
			log.Println(err)
		}
		io.Copy(chatId, strings.NewReader(strconv.Itoa(id)))
		photoField, err := w.CreateFormFile("photo", p.Name())
		io.Copy(photoField, p)
		w.Close()
		req.Header.Set("Content-Type", w.FormDataContentType())
		req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
	case string:
		q := req.URL.Query()
		q.Add("chat_id", strconv.Itoa(id))
		q.Add("photo", p)
		req.URL.RawQuery = q.Encode()
	default:
		return "", errors.New("SendPhoto function accepts string and *os.File types")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	resToString, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(resToString), nil
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
