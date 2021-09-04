package gogram

import (
	"bytes"
	"fmt"
	"io"
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

func sendPhotoLogic(b Bot, id int, photo os.File) {
	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	field, err := w.CreateFormField("chat_id")
	if err != nil {
		return
	}
	_, err = io.Copy(field, strings.NewReader(strconv.Itoa(id)))
	if err != nil {
		return
	}

	file, err := w.CreateFormFile("photo", photo.Name())
	if err != nil {
		return
	}
	_, err = io.Copy(file, &photo)
	if err != nil {
		return
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", b.Token),
		bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", w.FormDataContentType())

	//q := url.Values{}
	//q.Add("chat_id", strconv.Itoa(id))
	//q.Add("photo", photo)
	//req.Header.Add("Content-Type", "multipart/form-data")
	//q := req.URL.Query()

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}

func (u User) SendPhoto(b Bot, photo os.File) {
	if u.Id == 0 {
		log.Println("User's Id field is empty")
	} else {
		sendPhotoLogic(b, u.Id, photo)
	}
}

func (c Chat) SendPhoto(b Bot, photo os.File) {
	if c.Id == 0 {
		log.Println("Chat's Id field is empty")
	} else {
		sendPhotoLogic(b, c.Id, photo)
	}
}
