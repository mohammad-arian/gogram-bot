package gogram

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func sendTextLogic(b Bot, id int, text string, optionalParams *TextOptionalParams) (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token),
		nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Set("chat_id", strconv.Itoa(id))
	q.Set("text", text)
	if optionalParams != nil {
		if optionalParams.ReplyMarkup != nil {
			q.Set("reply_markup", optionalParams.ReplyMarkup.toString())
		}
		q.Set("disable_notification", strconv.FormatBool(optionalParams.DisableNotification))
		q.Set("parse_mode", optionalParams.ParseMode)
		q.Set("disable_web_page_preview", strconv.FormatBool(optionalParams.DisableWebPagePreview))
		q.Set("reply_to_message_id", strconv.Itoa(optionalParams.ReplyToMessageId))
		q.Set("allow_sending_without_reply", strconv.FormatBool(optionalParams.AllowSendingWithoutReply))
		//q.Set("entities", optionalParams.ParseMode)     // fix this part later
	}
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resToString, _ := ioutil.ReadAll(res.Body)
	return string(resToString), nil
}

// SendText sends message to a User.
// b Bot parameter indicated which bot to send
// the message with. This way you can send messages with different bots
// text is the message that will be sent
// pass nil or *TextOptionalParams struct to optionalParams. It adds some optional
// parameters to request, like reply_markup, disable_notification and ...
func (u User) SendText(b Bot, text string, optionalParams *TextOptionalParams) (response string, err error) {
	if u.Id == 0 {
		return "", errors.New("user's Id field is empty")
	}
	return sendTextLogic(b, u.Id, text, optionalParams)
}

// SendText sends message to a Chat.
// b Bot parameter indicated which bot to send
// the message with. This way you can send messages with different bots
// text is the message that will be sent
// pass nil or *TextOptionalParams struct to optionalParams. It adds some optional
// parameters to request, like reply_markup, disable_notification and ...
func (c Chat) SendText(b Bot, text string, optionalParams *TextOptionalParams) (response string, err error) {
	if c.Id == 0 {
		return "", errors.New("chat's Id field is empty")
	}
	return sendTextLogic(b, c.Id, text, optionalParams)
}

func sendPhotoLogic(b Bot, id int, photo interface{}, optionalParams *TextOptionalParams) (string, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", b.Token),
		nil)
	switch p := photo.(type) {
	case *os.File:
		var body = &bytes.Buffer{}
		w := multipart.NewWriter(body)
		chatId, err := w.CreateFormField("chat_id")
		if err != nil {
			return "", err
		}
		_, err = io.Copy(chatId, strings.NewReader(strconv.Itoa(id)))
		if err != nil {
			return "", err
		}
		photoField, err := w.CreateFormFile("photo", p.Name())
		_, err = io.Copy(photoField, p)
		if err != nil {
			return "", err
		}
		err = w.Close()
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", w.FormDataContentType())
		req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
	case string:
		q := req.URL.Query()
		q.Set("chat_id", strconv.Itoa(id))
		q.Set("photo", p)
		req.URL.RawQuery = q.Encode()
	default:
		return "", errors.New("SendPhoto function accepts string and *os.File types")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resToString, _ := ioutil.ReadAll(res.Body)
	return string(resToString), nil
}

func (u User) SendPhoto(b Bot, photo interface{}, optionalParams *TextOptionalParams) (response string, err error) {
	if u.Id == 0 {
		return "", errors.New("user's Id field is empty")
	}
	return sendPhotoLogic(b, u.Id, photo, optionalParams)
}

func (c Chat) SendPhoto(b Bot, photo interface{}, optionalParams *TextOptionalParams) (response string, err error) {
	if c.Id == 0 {
		return "", errors.New("chat's Id field is empty")
	}
	return sendPhotoLogic(b, c.Id, photo, optionalParams)
}

func sendVideoLogic(b Bot, id int, video interface{}) (response string, err error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendVideo", b.Token),
		nil)
	switch p := video.(type) {
	case *os.File:
		var body = &bytes.Buffer{}
		w := multipart.NewWriter(body)
		chatId, err := w.CreateFormField("chat_id")
		if err != nil {
			return "", err
		}
		_, err = io.Copy(chatId, strings.NewReader(strconv.Itoa(id)))
		if err != nil {
			return "", err
		}
		photoField, err := w.CreateFormFile("video", p.Name())
		_, err = io.Copy(photoField, p)
		if err != nil {
			return "", err
		}
		err = w.Close()
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", w.FormDataContentType())
		req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
	case string:
		q := req.URL.Query()
		q.Set("chat_id", strconv.Itoa(id))
		q.Set("video", p)
		req.URL.RawQuery = q.Encode()
	default:
		return "", errors.New("SendVideo function accepts string and *os.File types")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resToString, _ := ioutil.ReadAll(res.Body)
	return string(resToString), nil
}

func (u User) SendVideo(b Bot, video interface{}) (response string, err error) {
	if u.Id == 0 {
		return "", errors.New("user's Id field is empty")
	}
	return sendVideoLogic(b, u.Id, video)
}

func (c Chat) SendVideo(b Bot, video interface{}) (response string, err error) {
	if c.Id == 0 {
		return "", errors.New("chat's Id field is empty")
	}
	return sendVideoLogic(b, c.Id, video)
}
