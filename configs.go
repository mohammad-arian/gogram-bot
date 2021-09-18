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

// AddInlineKeyboardButtonColumn add a InlineKeyboard in vertical orientation.
// if ReplyMarkup of TextOptionalParams is nil or of type ReplyKeyboardMarkup
// it will be set to InlineKeyboardMarkup, else if ReplyMarkup of TextOptionalParams
// is already of type InlineKeyboardButton, buttons will be added to it.
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

// AddInlineKeyboardButtonRow add a InlineKeyboard in horizontal orientation.
// if ReplyMarkup of TextOptionalParams is nil or of type ReplyKeyboardMarkup
// it will be set to InlineKeyboardMarkup, else if ReplyMarkup of TextOptionalParams
// is already of type InlineKeyboardButton, buttons will be added to it.
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

// AddReplyKeyboardButtonColumn add a InlineKeyboard in vertical orientation.
// if ReplyMarkup of TextOptionalParams is nil or of type InlineKeyboardMarkup
// it will be set to ReplyKeyboardMarkup, else if ReplyMarkup of TextOptionalParams
// is already of type ReplyKeyboardMarkup, buttons will be added to it.
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

// AddReplyKeyboardButtonRow add a InlineKeyboard in horizontal orientation.
// if ReplyMarkup of TextOptionalParams is nil or of type InlineKeyboardMarkup
// it will be set to ReplyKeyboardMarkup, else if ReplyMarkup of TextOptionalParams
// is already of type ReplyKeyboardMarkup, buttons will be added to it.
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
	// Optional. Requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available,
	// but clients will automatically display the usual letter-keyboard in the chat
	// the user can press a special button in the input field to see the custom keyboard again.
	OneTimeKeyboard bool `json:"one_time_keyboard"`
	// Optional. Requests clients to resize the keyboard vertically for optimal fit.
	// Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard"`
}

func (i *ReplyKeyboardMarkup) toString() string {
	a, _ := json.Marshal(i)
	return string(a)
}

// ReplyKeyboardRemove represents an object that if Telegram clients receive,
// they will remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
type ReplyKeyboardRemove struct {
	// Requests clients to remove the custom keyboard
	// (user will not be able to summon this keyboard;
	// if you want to hide the keyboard from sight but keep it accessible,
	// use OneTimeKeyboard in ReplyKeyboardMarkup).
	// Always set to true.
	RemoveKeyboard bool `json:"remove_keyboard"`
	// Set Selective to true if you want to remove the keyboard for specific users only.
	// Targets: 1) users that are @mentioned in the text of the Message object;
	//          2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	// Example: A user votes in a poll, bot returns confirmation message in reply
	// to the vote and removes the keyboard for that user,
	// while still showing the keyboard with poll options to users who haven't voted yet.
	Selective bool `json:"selective"`
}

func (i *ReplyKeyboardRemove) toString() string {
	a, _ := json.Marshal(i)
	return string(a)
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

// KeyboardButton represents one button of a reply keyboard.
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used,
	// it will be sent as a message when the button is pressed
	Text string `json:"text"`
	// Optional. If True, the user's phone number will be sent as a contact when the button is pressed.
	// Available in private chats only
	RequestContact bool `json:"request_contact"`
	// Optional. If True, the user's current location will be sent when the button is pressed.
	// Available in private chats only
	RequestLocation bool `json:"request_location"`
}

type MessageEntity struct {
	Type     string
	offset   int
	length   int
	url      string
	user     User
	language string
}
