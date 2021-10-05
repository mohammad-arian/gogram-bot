package gogram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
)

func urlValueSetter(s interface{}, q *url.Values, key ...string) {
	if reflect.TypeOf(s).Kind() == reflect.Struct {
		for i := 0; i < reflect.ValueOf(s).NumField(); i++ {
			tag := reflect.TypeOf(s).Field(i).Tag.Get("json")
			value := reflect.ValueOf(s).Field(i).Interface()
			switch j := value.(type) {
			case string:
				q.Set(tag, value.(string))
			case int:
				q.Set(tag, strconv.Itoa(value.(int)))
			case bool:
				q.Set(tag, strconv.FormatBool(value.(bool)))
			case InlineKeyboard:
				if j.inlineKeyboardMarkup.InlineKeyboardButtons != nil {
					a, _ := json.Marshal(j.inlineKeyboardMarkup)
					q.Set("reply_markup", string(a))
				}
			case ReplyKeyboard:
				if j.replyKeyboardMarkup.Keyboard != nil {
					a, _ := json.Marshal(j.replyKeyboardMarkup)
					q.Set("reply_markup", string(a))
				} else if j.replyKeyboardRemove != (replyKeyboardRemove{}) {
					a, _ := json.Marshal(j.replyKeyboardRemove)
					q.Set("reply_markup", string(a))
				}
			case ForceReply:
				a, _ := json.Marshal(j)
				q.Set("reply_markup", string(a))
			}
		}
	} else if reflect.TypeOf(s).Kind() == reflect.Slice {
		if key != nil {
			a, _ := json.Marshal(s)
			q.Set(key[0], string(a))
		}
	} else if reflect.TypeOf(s).Kind() == reflect.Map {
		v := reflect.ValueOf(s)
		for _, i := range v.MapKeys() {
			q.Set(i.String(), v.MapIndex(i).String())
		}
	}
}

func formFieldSetter(s interface{}, w *multipart.Writer, key ...string) {
	if reflect.TypeOf(s).Kind() == reflect.Struct {
		for i := 0; i < reflect.ValueOf(s).NumField(); i++ {
			tag := reflect.TypeOf(s).Field(i).Tag.Get("json")
			value := reflect.ValueOf(s).Field(i).Interface()
			switch j := value.(type) {
			case string:
				_ = w.WriteField(tag, value.(string))
			case int:
				_ = w.WriteField(tag, strconv.Itoa(value.(int)))
			case bool:
				_ = w.WriteField(tag, strconv.FormatBool(value.(bool)))
			case InlineKeyboard:
				if j.inlineKeyboardMarkup.InlineKeyboardButtons != nil {
					a, _ := json.Marshal(j.inlineKeyboardMarkup)
					_ = w.WriteField("reply_markup", string(a))
				}
			case ReplyKeyboard:
				if j.replyKeyboardMarkup.Keyboard != nil {
					a, _ := json.Marshal(j.replyKeyboardMarkup)
					_ = w.WriteField("reply_markup", string(a))
				} else if j.replyKeyboardRemove != (replyKeyboardRemove{}) {
					a, _ := json.Marshal(j.replyKeyboardRemove)
					_ = w.WriteField("reply_markup", string(a))
				}
			case ForceReply:
				a, _ := json.Marshal(j)
				_ = w.WriteField("reply_markup", string(a))
			case *os.File:
				file, _ := w.CreateFormFile(tag, j.Name())
				_, _ = io.Copy(file, j)
				_, _ = j.Seek(0, io.SeekStart)
			}
		}
	}
}

// inlineKeyboardButtonColumnAdder add a InlineKeyboard in vertical orientation.
// if ReplyMarkup of TextOptionalParams is nil or of type replyKeyboardMarkup
// it will be set to inlineKeyboardMarkup, else if ReplyMarkup of TextOptionalParams
// is already of type InlineKeyboardButton, buttons will be added to it.
func inlineKeyboardButtonColumnAdder(t *InlineKeyboard, i ...InlineKeyboardButton) {
	var column [][]InlineKeyboardButton
	for _, button := range i {
		column = append(column, []InlineKeyboardButton{button})
	}
	if t.InlineKeyboardButtons == nil {
		t.InlineKeyboardButtons = column
	} else {
		t.InlineKeyboardButtons = append(t.InlineKeyboardButtons, column...)
	}
}

// inlineKeyboardButtonRowAdder is like inlineKeyboardButtonColumnAdder but adds a
// inline keyboard in horizontal orientation.
func inlineKeyboardButtonRowAdder(t *InlineKeyboard, i ...InlineKeyboardButton) {
	row := [][]InlineKeyboardButton{{}}
	for _, button := range i {
		row[0] = append(row[0], button)
	}
	if t.InlineKeyboardButtons == nil {
		t.InlineKeyboardButtons = row
	} else {
		t.InlineKeyboardButtons = append(t.InlineKeyboardButtons, row...)
	}
}

// AddReplyKeyboardButtonColumn add a InlineKeyboard in vertical orientation.
// if ReplyMarkup of TextOptionalParams is nil or of type inlineKeyboardMarkup
// it will be set to replyKeyboardMarkup, else if ReplyMarkup of TextOptionalParams
// is already of type replyKeyboardMarkup, buttons will be added to it.
// set oneTimeKeyboard to true to request clients to hide the keyboard as soon as it's been used.
// The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat -
// the user can press a special button in the input field to see the custom keyboard again. otherwise, set to false.
// set selective to true if you want to show the keyboard to specific users only. otherwise, set to false
// Targets: 1) users that are @mentioned in the text of the Message object;
//          2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
// Example: A user requests to change the bot's language, bot replies to the request with a
// keyboard to select the new language. Other users in the group don't see the keyboard.
// inputFieldPlaceholder is the placeholder to be shown in the input field when the keyboard is active; pass
// empty or any string
func replyKeyboardButtonColumnAdder(t *ReplyKeyboard, oneTimeKeyboard bool,
	resizeKeyboard bool, inputFieldPlaceholder string, selective bool, i ...ReplyKeyboardButton) {
	var column [][]ReplyKeyboardButton
	for _, button := range i {
		column = append(column, []ReplyKeyboardButton{button})
	}
	if t.Keyboard == nil {
		t.replyKeyboardMarkup = replyKeyboardMarkup{Keyboard: column, ResizeKeyboard: resizeKeyboard,
			OneTimeKeyboard: oneTimeKeyboard, InputFieldPlaceholder: inputFieldPlaceholder, Selective: selective}
	} else {
		t.Keyboard = append(t.Keyboard, column...)
	}
}

// AddReplyKeyboardButtonRow is like AddReplyKeyboardButtonColumn but adds a
// InlineKeyboard in horizontal orientation.
func replyKeyboardButtonRowAdder(t *ReplyKeyboard, oneTimeKeyboard bool,
	resizeKeyboard bool, inputFieldPlaceholder string, selective bool, i ...ReplyKeyboardButton) {
	row := [][]ReplyKeyboardButton{{}}
	for _, button := range i {
		row[0] = append(row[0], button)
	}
	if t.Keyboard == nil {
		t.replyKeyboardMarkup = replyKeyboardMarkup{Keyboard: row, ResizeKeyboard: resizeKeyboard,
			OneTimeKeyboard: oneTimeKeyboard, InputFieldPlaceholder: inputFieldPlaceholder, Selective: selective}
	} else {
		t.Keyboard = append(t.Keyboard, row...)
	}
}

func request(id int, method string, token string, containsFile bool, data interface{},
	optionalParams interface{}) (response string, error error) {
	if id == 0 {
		return "", errors.New("id field is empty")
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.telegram.org/bot%s/send%s", token, method),
		nil)
	if err != nil {
		return "", err
	}
	if !containsFile {
		q := req.URL.Query()
		urlValueSetter(data, &q)
		if optionalParams != nil {
			urlValueSetter(optionalParams, &q)
		}
		req.URL.RawQuery = q.Encode()
	} else {
		var body = &bytes.Buffer{}
		w := multipart.NewWriter(body)
		formFieldSetter(data, w)
		if optionalParams != nil {
			formFieldSetter(optionalParams, w)
		}
		err = w.Close()
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", w.FormDataContentType())
		req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resToString, _ := ioutil.ReadAll(res.Body)
	return string(resToString), nil
}
