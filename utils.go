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
	"os"
	"reflect"
	"strconv"
)

func formFieldSetter(s interface{}, w *multipart.Writer) error {
	for i := 0; i < reflect.ValueOf(s).NumField(); i++ {
		tag := reflect.TypeOf(s).Field(i).Tag.Get("json")
		value := reflect.ValueOf(s).Field(i).Interface()
		switch j := value.(type) {
		case string:
			err := w.WriteField(tag, value.(string))
			if err != nil {
				return err
			}
		case int:
			err := w.WriteField(tag, strconv.Itoa(value.(int)))
			if err != nil {
				return err
			}
		case float64:
			err := w.WriteField(tag, fmt.Sprintf("%v", value.(float64)))
			if err != nil {
				return err
			}
		case bool:
			err := w.WriteField(tag, strconv.FormatBool(value.(bool)))
			if err != nil {
				return err
			}
		case InlineKeyboard:
			if j.inlineKeyboardMarkup.InlineKeyboardButtons != nil {
				a, _ := json.Marshal(j.inlineKeyboardMarkup)
				err := w.WriteField("reply_markup", string(a))
				if err != nil {
					return err
				}
			}
		case ReplyKeyboard:
			if j.replyKeyboardMarkup.Keyboard != nil {
				a, _ := json.Marshal(j.replyKeyboardMarkup)
				err := w.WriteField("reply_markup", string(a))
				if err != nil {
					return err
				}
			} else if j.replyKeyboardRemove != (replyKeyboardRemove{}) {
				a, _ := json.Marshal(j.replyKeyboardRemove)
				err := w.WriteField("reply_markup", string(a))
				if err != nil {
					return err
				}
			}
		case ForceReply:
			a, _ := json.Marshal(j)
			err := w.WriteField("reply_markup", string(a))
			if err != nil {
				return err
			}
		case botCommandScope:
			obj, err := j.botCommandReturn()
			if err != nil {
				return err
			}
			err = formFieldSetter(obj, w)
			if err != nil {
				return err
			}
		case *os.File:
			file, _ := w.CreateFormFile(tag, j.Name())
			_, _ = io.Copy(file, j)
			_, _ = j.Seek(0, io.SeekStart)
		case []*os.File:
			for _, f := range j {
				file, _ := w.CreateFormFile(f.Name(), f.Name())
				_, _ = io.Copy(file, f)
				_, _ = f.Seek(0, io.SeekStart)
			}
		}
	}
	return nil
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

func request(id int, method string, token string, data interface{},
	optionalParams interface{}, responseType interface{}) (response interface{}, error error) {
	if id == 0 {
		return "", errors.New("id field is empty")
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, method),
		nil)
	if err != nil {
		return "", err
	}
	var body = &bytes.Buffer{}
	w := multipart.NewWriter(body)
	if data != nil {
		err = formFieldSetter(data, w)
		if err != nil {
			return nil, err
		}
	}
	if optionalParams != nil {
		err = formFieldSetter(optionalParams, w)
		if err != nil {
			return nil, err
		}
	}
	err = w.Close()
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	readRes, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(readRes, responseType)
	if err != nil {
		return responseType, err
	}
	return responseType, nil
}
