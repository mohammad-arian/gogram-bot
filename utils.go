package gogram

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
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
		case KeyboardMarkup:
			q.Set(tag, j.toString())
		case MessageEntity:
			a, _ := json.Marshal(i)
			q.Set(tag, string(a))
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
		case KeyboardMarkup:
			field, _ := w.CreateFormField(tag)
			_, _ = io.Copy(field, strings.NewReader(j.toString()))
		case MessageEntity:
			field, _ := w.CreateFormField(tag)
			a, _ := json.Marshal(i)
			_, _ = io.Copy(field, strings.NewReader(string(a)))
		}
	}
}

// inlineKeyboardButtonColumnAdder add a InlineKeyboard in vertical orientation.
// if ReplyMarkup of TextOptionalParams is nil or of type replyKeyboardMarkup
// it will be set to inlineKeyboardMarkup, else if ReplyMarkup of TextOptionalParams
// is already of type InlineKeyboardButton, buttons will be added to it.
func inlineKeyboardButtonColumnAdder(t *TextOptionalParams, i ...InlineKeyboardButton) {
	var column [][]InlineKeyboardButton
	for _, button := range i {
		column = append(column, []InlineKeyboardButton{button})
	}
	switch r := t.ReplyMarkup.(type) {
	case *inlineKeyboardMarkup:
		r.InlineKeyboard = append(r.InlineKeyboard, column...)
	default:
		t.ReplyMarkup = &inlineKeyboardMarkup{InlineKeyboard: column}
	}
}

// AddInlineKeyboardButtonRow is like AddInlineKeyboardButtonColumn but adds a
// InlineKeyboard in horizontal orientation.
func inlineKeyboardButtonRowAdder(t *TextOptionalParams, i ...InlineKeyboardButton) {
	row := [][]InlineKeyboardButton{{}}
	for _, button := range i {
		row[0] = append(row[0], button)
	}
	switch r := t.ReplyMarkup.(type) {
	case *inlineKeyboardMarkup:
		r.InlineKeyboard = append(r.InlineKeyboard, row...)
	default:
		t.ReplyMarkup = &inlineKeyboardMarkup{InlineKeyboard: row}
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
func replyKeyboardButtonColumnAdder(t *TextOptionalParams, oneTimeKeyboard bool,
	resizeKeyboard bool, inputFieldPlaceholder string, selective bool, i ...KeyboardButton) {
	var column [][]KeyboardButton
	for _, button := range i {
		column = append(column, []KeyboardButton{button})
	}
	switch r := t.ReplyMarkup.(type) {
	case *replyKeyboardMarkup:
		r.Keyboard = append(r.Keyboard, column...)
	default:
		t.ReplyMarkup = &replyKeyboardMarkup{Keyboard: column, ResizeKeyboard: resizeKeyboard,
			OneTimeKeyboard: oneTimeKeyboard, InputFieldPlaceholder: inputFieldPlaceholder, Selective: selective}
	}
}

// AddReplyKeyboardButtonRow is like AddReplyKeyboardButtonColumn but adds a
// InlineKeyboard in horizontal orientation.
func replyKeyboardButtonRowAdder(t *TextOptionalParams, oneTimeKeyboard bool,
	resizeKeyboard bool, inputFieldPlaceholder string, selective bool, i ...KeyboardButton) {
	row := [][]KeyboardButton{{}}
	for _, button := range i {
		row[0] = append(row[0], button)
	}
	switch r := t.ReplyMarkup.(type) {
	case *replyKeyboardMarkup:
		r.Keyboard = append(r.Keyboard, row...)
	default:
		t.ReplyMarkup = &replyKeyboardMarkup{Keyboard: row, ResizeKeyboard: resizeKeyboard,
			OneTimeKeyboard: oneTimeKeyboard, InputFieldPlaceholder: inputFieldPlaceholder, Selective: selective}
	}
}
