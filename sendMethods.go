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

// SendText sends message to a User.
// b Bot parameter indicated which bot to send
// the message with. This way you can send messages with different bots
// text is the message that will be sent
// pass nil or *TextOptionalParams struct to optionalParams. It adds some optional
// parameters to request, like reply_markup, disable_notification and ...
func (r *ReplyAble) SendText(b Bot, text string, optionalParams *TextOptionalParams) (response string, err error) {
	var id = r.Id
	if id == 0 {
		return "", errors.New("id field is empty")
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token),
		nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Set("chat_id", strconv.Itoa(id))
	q.Set("text", text)
	if optionalParams != nil {
		urlValueSetter(*optionalParams, &q)
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

func (r *ReplyAble) SendPhoto(b Bot, photo interface{}, optionalParams *PhotoOptionalParams) (response string, err error) {
	var id = r.Id
	if id == 0 {
		return "", errors.New("id field is empty")
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", b.Token),
		nil)
	if err != nil {
		return "", err
	}
	switch p := photo.(type) {
	case *os.File:
		all, err := ioutil.ReadAll(p)
		if err != nil {
			return "", err
		}
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
		_, err = photoField.Write(all)
		if err != nil {
			return "", err
		}
		if err != nil {
			return "", err
		}
		if optionalParams != nil {
			formFieldSetter(*optionalParams, w)
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
		if optionalParams != nil {
			urlValueSetter(*optionalParams, &q)
		}
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

func (r *ReplyAble) SendVideo(b Bot, video interface{}) (response string, err error) {
	var id = r.Id
	if id == 0 {
		return "", errors.New("id field is empty")
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendVideo", b.Token),
		nil)
	if err != nil {
		return "", err
	}
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
