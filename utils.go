package gogram

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func urlValueSetter(s interface{}, q *url.Values) {
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
			if j.forceReply != (forceReply{}) {
				a, _ := json.Marshal(j.forceReply)
				q.Set("reply_markup", string(a))
			}
		}
	}
}

func formFieldSetter(s interface{}, w *multipart.Writer) {
	for i := 0; i < reflect.ValueOf(s).NumField(); i++ {
		tag := reflect.TypeOf(s).Field(i).Tag.Get("json")
		value := reflect.ValueOf(s).Field(i).Interface()
		switch j := value.(type) {
		case string:
			field, _ := w.CreateFormField(tag)
			_, _ = io.Copy(field, strings.NewReader(value.(string)))
		case int:
			field, _ := w.CreateFormField(tag)
			_, _ = io.Copy(field, strings.NewReader(strconv.Itoa(value.(int))))
		case bool:
			field, _ := w.CreateFormField(tag)
			_, _ = io.Copy(field, strings.NewReader(strconv.FormatBool(value.(bool))))
		case InlineKeyboard:
			if j.inlineKeyboardMarkup.InlineKeyboardButtons != nil {
				a, _ := json.Marshal(j.inlineKeyboardMarkup)
				field, _ := w.CreateFormField("reply_markup")
				_, _ = io.Copy(field, strings.NewReader(string(a)))
			}
		case ReplyKeyboard:
			if j.replyKeyboardMarkup.Keyboard != nil {
				a, _ := json.Marshal(j.replyKeyboardMarkup)
				field, _ := w.CreateFormField("reply_markup")
				_, _ = io.Copy(field, strings.NewReader(string(a)))
			} else if j.replyKeyboardRemove != (replyKeyboardRemove{}) {
				a, _ := json.Marshal(j.replyKeyboardRemove)
				field, _ := w.CreateFormField("reply_markup")
				_, _ = io.Copy(field, strings.NewReader(string(a)))
			}
		case ForceReply:
			if j.forceReply != (forceReply{}) {
				a, _ := json.Marshal(j.forceReply)
				field, _ := w.CreateFormField("reply_markup")
				_, _ = io.Copy(field, strings.NewReader(string(a)))
			}
		case *os.File:
			field, _ := w.CreateFormField(tag)
			all, _ := ioutil.ReadAll(j)
			_, _ = j.Seek(0, io.SeekStart)
			_, _ = io.Copy(field, strings.NewReader(string(all)))
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
