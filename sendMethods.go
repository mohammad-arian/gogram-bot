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
		all, err := ioutil.ReadAll(p)
		if err != nil {
			return "", err
		}
		_, err = p.Seek(0, io.SeekStart)
		if err != nil {
			return "", err
		}
		_, err = io.Copy(photoField, strings.NewReader(string(all)))
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
		return "", errors.New("SendPhoto function accepts only string and *os.File types")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resToString, _ := ioutil.ReadAll(res.Body)
	return string(resToString), nil
}

func (r *ReplyAble) SendVideo(b Bot, video interface{}, optionalParams *VideoOptionalParams) (response string, err error) {
	var id = r.Id
	if id == 0 {
		return "", errors.New("id field is empty")
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendVideo", b.Token),
		nil)
	if err != nil {
		return "", err
	}
	switch v := video.(type) {
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
		videoField, err := w.CreateFormFile("video", v.Name())
		all, err := ioutil.ReadAll(v)
		if err != nil {
			return "", err
		}
		_, err = v.Seek(0, io.SeekStart)
		if err != nil {
			return "", err
		}
		_, err = io.Copy(videoField, strings.NewReader(string(all)))
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
		q.Set("video", v)
		if optionalParams != nil {
			urlValueSetter(*optionalParams, &q)
		}
		req.URL.RawQuery = q.Encode()
	default:
		return "", errors.New("SendVideo function accepts only string and *os.File types")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resToString, _ := ioutil.ReadAll(res.Body)
	return string(resToString), nil
}

// SendAudio sends audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format.
// On success, the sent Message is returned.
// Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
func (r *ReplyAble) SendAudio(b Bot, audio interface{}, optionalParams *AudioOptionalParams) (response string, err error) {
	var id = r.Id
	if id == 0 {
		return "", errors.New("id field is empty")
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendAudio", b.Token),
		nil)
	if err != nil {
		return "", err
	}
	switch v := audio.(type) {
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
		videoField, err := w.CreateFormFile("audio", v.Name())
		all, err := ioutil.ReadAll(v)
		if err != nil {
			return "", err
		}
		_, err = v.Seek(0, io.SeekStart)
		if err != nil {
			return "", err
		}
		_, err = io.Copy(videoField, strings.NewReader(string(all)))
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
		q.Set("audio", v)
		if optionalParams != nil {
			urlValueSetter(*optionalParams, &q)
		}
		req.URL.RawQuery = q.Encode()
	default:
		return "", errors.New("SendAudio function accepts only string and *os.File types")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resToString, _ := ioutil.ReadAll(res.Body)
	return string(resToString), nil
}

func (r *ReplyAble) SendDocument(b Bot, document interface{}, optionalParams *DocumentOptionalParams) (response string, err error) {
	var id = r.Id
	if id == 0 {
		return "", errors.New("id field is empty")
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", b.Token),
		nil)
	if err != nil {
		return "", err
	}
	switch v := document.(type) {
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
		videoField, err := w.CreateFormFile("document", v.Name())
		all, err := ioutil.ReadAll(v)
		if err != nil {
			return "", err
		}
		_, err = v.Seek(0, io.SeekStart)
		if err != nil {
			return "", err
		}
		_, err = io.Copy(videoField, strings.NewReader(string(all)))
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
		q.Set("document", v)
		if optionalParams != nil {
			urlValueSetter(*optionalParams, &q)
		}
		req.URL.RawQuery = q.Encode()
	default:
		return "", errors.New("SendDocument function accepts only string and *os.File types")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resToString, _ := ioutil.ReadAll(res.Body)
	return string(resToString), nil
}

// SendVoice sends audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with
// OPUS (other formats may be sent as Audio or Document).
// On success, the sent Message is returned.
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (r *ReplyAble) SendVoice(b Bot, voice interface{}, optionalParams *VoiceOptionalParams) (response string, err error) {
	var id = r.Id
	if id == 0 {
		return "", errors.New("id field is empty")
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendVoice", b.Token),
		nil)
	if err != nil {
		return "", err
	}
	switch v := voice.(type) {
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
		videoField, err := w.CreateFormFile("voice", v.Name())
		all, err := ioutil.ReadAll(v)
		if err != nil {
			return "", err
		}
		_, err = v.Seek(0, io.SeekStart)
		if err != nil {
			return "", err
		}
		_, err = io.Copy(videoField, strings.NewReader(string(all)))
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
		q.Set("voice", v)
		if optionalParams != nil {
			urlValueSetter(*optionalParams, &q)
		}
		req.URL.RawQuery = q.Encode()
	default:
		return "", errors.New("SendVoice function accepts only string and *os.File types")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resToString, _ := ioutil.ReadAll(res.Body)
	return string(resToString), nil
}

func (r *ReplyAble) SendAnimation(b Bot, animation interface{}, optionalParams *AnimationOptionalParams) (response string, err error) {
	var id = r.Id
	if id == 0 {
		return "", errors.New("id field is empty")
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendAnimation", b.Token),
		nil)
	if err != nil {
		return "", err
	}
	switch v := animation.(type) {
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
		videoField, err := w.CreateFormFile("animation", v.Name())
		all, err := ioutil.ReadAll(v)
		if err != nil {
			return "", err
		}
		_, err = v.Seek(0, io.SeekStart)
		if err != nil {
			return "", err
		}
		_, err = io.Copy(videoField, strings.NewReader(string(all)))
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
		q.Set("animation", v)
		if optionalParams != nil {
			urlValueSetter(*optionalParams, &q)
		}
		req.URL.RawQuery = q.Encode()
	default:
		return "", errors.New("SendAnimation function accepts only string and *os.File types")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resToString, _ := ioutil.ReadAll(res.Body)
	return string(resToString), nil
}

func (r *ReplyAble) SendPoll(b Bot, question string, options []string, optionalParams *PollOptionalParams) (response string, err error) {
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
	q.Set("question", question)
	urlValueSetter(options, &q)
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
