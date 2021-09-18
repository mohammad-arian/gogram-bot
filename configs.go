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
	inlineKeyboardButtonColumnAdder(t, i...)
}

func (t *TextOptionalParams) AddInlineKeyboardRowColumn(i ...InlineKeyboardButton) {
	inlineKeyboardButtonRowAdder(t, i...)
}

func (t *TextOptionalParams) AddReplyKeyboardButtonRow(oneTimeKeyboard bool,
	resizeKeyboard bool, inputFieldPlaceholder string, selective bool, i ...KeyboardButton) {
	replyKeyboardButtonRowAdder(t, oneTimeKeyboard, resizeKeyboard, inputFieldPlaceholder, selective, i...)
}

func (t *TextOptionalParams) AddReplyKeyboardButtonColumn(oneTimeKeyboard bool,
	resizeKeyboard bool, inputFieldPlaceholder string, selective bool, i ...KeyboardButton) {
	replyKeyboardButtonColumnAdder(t, oneTimeKeyboard, resizeKeyboard, inputFieldPlaceholder, selective, i...)
}

// RemoveReplyKeyboard Removes the reply keyboard.
// Set selective to true if you want to remove the keyboard for specific users only.
// Targets: 1) users that are @mentioned in the text of the Message object;
//          2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
// Example: A user votes in a poll, bot returns confirmation message in reply
// to the vote and removes the keyboard for that user,
// while still showing the keyboard with poll options to users who haven't voted yet.
func (t *TextOptionalParams) RemoveReplyKeyboard(selective bool) {
	t.ReplyMarkup = &replyKeyboardRemove{RemoveKeyboard: true, Selective: selective}
}

// ForceReply sends a request to telegram clients to display
// a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply').
// This can be extremely useful if you want to create user-friendly
// step-by-step interfaces without having to sacrifice privacy mode.
func (t *TextOptionalParams) ForceReply(selective bool, inputFieldPlaceholder string) {
	t.ReplyMarkup = &forceReply{ForceReply: true, Selective: selective, InputFieldPlaceholder: inputFieldPlaceholder}
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

type inlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (i *inlineKeyboardMarkup) toString() string {
	a, _ := json.Marshal(i)
	return string(a)
}

type replyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	OneTimeKeyboard       bool               `json:"one_time_keyboard"`
	ResizeKeyboard        bool               `json:"resize_keyboard"`
	InputFieldPlaceholder string             `json:"input_field_placeholder"`
	Selective             bool               `json:"selective"`
}

func (i *replyKeyboardMarkup) toString() string {
	a, _ := json.Marshal(i)
	return string(a)
}

// replyKeyboardRemove represents an object that if Telegram clients receive,
// they will remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
type replyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

func (i *replyKeyboardRemove) toString() string {
	a, _ := json.Marshal(i)
	return string(a)
}

type forceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder"`
	Selective             bool   `json:"selective"`
}

func (i *forceReply) toString() string {
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
	// Optional. If set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	// Can be empty, in which case just the bot's username will be inserted.
	// Note: This offers an easy way for users to start using your bot in inline mode
	// when they are currently in a private chat with it.
	// Especially useful when combined with switch_pm… actions – in this case the user will be
	// automatically returned to the chat they switched from, skipping the chat selection screen.
	SwitchInlineQuery string `json:"switch_inline_query"`
	// Optional. If set, pressing the button will insert the bot's username and the specified
	// inline query in the current chat's input field.
	// Can be empty, in which case only the bot's username will be inserted.
	// This offers a quick way for the user to open your bot in inline mode
	// in the same chat – good for selecting something from multiple options.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat"`
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

func (i *MessageEntity) toString() string {
	a, _ := json.Marshal(i)
	return string(a)
}
