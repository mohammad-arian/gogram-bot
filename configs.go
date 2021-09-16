package gogram

import (
	"encoding/json"
)

type KeyboardMarkup interface {
	toString() string
}

// TextOptionalParams represents optional parameters
// that SendText function can use.
// All fields are optional.
type TextOptionalParams struct {
	ParseMode                string          `json:"parse_mode"`
	Entities                 []MessageEntity `json:"entities"`
	DisableWebPagePreview    bool            `json:"disable_web_page_preview"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	ReplyMarkup              KeyboardMarkup  `json:"reply_markup"`
}

func (t *TextOptionalParams) AddInlineKeyboardButtonColumn(i ...InlineKeyboardButton) {
	var column [][]InlineKeyboardButton
	for _, button := range i {
		column = append(column, []InlineKeyboardButton{button})
	}
	switch r := t.ReplyMarkup.(type) {
	case *InlineKeyboardMarkup:
		r.InlineKeyboard = append(r.InlineKeyboard, column...)
	default:
		t.ReplyMarkup = &InlineKeyboardMarkup{InlineKeyboard: column}
	}
}
func (t *TextOptionalParams) AddInlineKeyboardButtonRow(i ...InlineKeyboardButton) {
	row := [][]InlineKeyboardButton{{}}
	for _, button := range i {
		row[0] = append(row[0], button)
	}
	switch r := t.ReplyMarkup.(type) {
	case *InlineKeyboardMarkup:
		r.InlineKeyboard = append(r.InlineKeyboard, row...)
	default:
		t.ReplyMarkup = &InlineKeyboardMarkup{InlineKeyboard: row}
	}
}

func (t *TextOptionalParams) AddReplyKeyboardButtonColumn(i ...KeyboardButton) {
	var column [][]KeyboardButton
	for _, button := range i {
		column = append(column, []KeyboardButton{button})
	}
	switch r := t.ReplyMarkup.(type) {
	case *ReplyKeyboardMarkup:
		r.Keyboard = append(r.Keyboard, column...)
	default:
		t.ReplyMarkup = &ReplyKeyboardMarkup{Keyboard: column}
	}
}

func (t *TextOptionalParams) AddReplyKeyboardButtonRow(i ...KeyboardButton) {
	row := [][]KeyboardButton{{}}
	for _, button := range i {
		row[0] = append(row[0], button)
	}
	switch r := t.ReplyMarkup.(type) {
	case *ReplyKeyboardMarkup:
		r.Keyboard = append(r.Keyboard, row...)
	default:
		t.ReplyMarkup = &ReplyKeyboardMarkup{Keyboard: row}
	}
}

type PhotoOptionalParams struct {
	ParseMode                string          `json:"parse_mode"`
	Caption                  string          `json:"caption"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	ReplyMarkup              KeyboardMarkup  `json:"reply_markup"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (i *InlineKeyboardMarkup) toString() string {
	a, _ := json.Marshal(i)
	return string(a)
}

type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton `json:"keyboard"`
}

func (i *ReplyKeyboardMarkup) toString() string {
	a, _ := json.Marshal(i)
	return string(a)
}

func (i InlineKeyboardMarkup) ReplyKeyboardRemove() {

}

// InlineKeyboardButton represents one button of an inline keyboard.
// You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	// Label text on the button
	Text string `json:"text"`
	// Optional. HTTP or tg:// url to be opened
	// when button is pressed
	Url string `json:"url"`
	// Optional. Data to be sent in a callback query
	// to the bot when button is pressed
	CallbackData string `json:"callback_data"`
}

type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used,
	// it will be sent as a message when the button is pressed
	Text string `json:"text"`
	// Optional.
	RequestContact bool `json:"request_contact"`
}

type MessageEntity struct {
	Type     string
	offset   int
	length   int
	url      string
	user     User
	language string
}
