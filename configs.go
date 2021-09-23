package gogram

type InlineKeyboard struct {
	inlineKeyboardMarkup
}

type ReplyKeyboard struct {
	replyKeyboardMarkup
	replyKeyboardRemove
}

type ForceReply struct {
	IsForceReply          bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder"`
	Selective             bool   `json:"selective"`
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
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

func (i *InlineKeyboard) AddInlineKeyboardButton(horizontal bool, a ...InlineKeyboardButton) {
	if horizontal {
		inlineKeyboardButtonRowAdder(i, a...)
	} else {
		inlineKeyboardButtonColumnAdder(i, a...)
	}
}

func (r *ReplyKeyboard) AddReplyKeyboardButton(horizontal bool, oneTimeKeyboard bool,
	resizeKeyboard bool, inputFieldPlaceholder string, selective bool, i ...ReplyKeyboardButton) {
	if horizontal {
		replyKeyboardButtonRowAdder(r, oneTimeKeyboard, resizeKeyboard, inputFieldPlaceholder, selective, i...)
	} else {
		replyKeyboardButtonColumnAdder(r, oneTimeKeyboard, resizeKeyboard, inputFieldPlaceholder, selective, i...)
	}
}

// RemoveReplyKeyboard Removes the reply keyboard.
// Set selective to true if you want to remove the keyboard for specific users only.
// Targets: 1) users that are @mentioned in the text of the Message object;
//          2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
// Example: A user votes in a poll, bot returns confirmation message in reply
// to the vote and removes the keyboard for that user,
// while still showing the keyboard with poll options to users who haven't voted yet.
func (r *ReplyKeyboard) RemoveReplyKeyboard(selective bool) {
	r.replyKeyboardRemove = replyKeyboardRemove{RemoveKeyboard: true, Selective: selective}
}

// SetForceReply sends a request to telegram clients to display
// a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply').
// This can be extremely useful if you want to create user-friendly
// step-by-step interfaces without having to sacrifice privacy mode.
func (t *ForceReply) SetForceReply(selective bool, inputFieldPlaceholder string) {
	t.IsForceReply = true
	t.Selective = selective
	t.InputFieldPlaceholder = inputFieldPlaceholder
}

type PhotoOptionalParams struct {
	ParseMode                string          `json:"parse_mode"`
	Caption                  string          `json:"caption"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

type VideoOptionalParams struct {
	Duration int    `json:"duration"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Caption  string `json:"caption"`
	// Set to a file of type *os.File or string.
	// Thumbnail of the file sent;
	// can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file.
	ParseMode                string          `json:"parse_mode"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	SupportsStreaming        bool            `json:"supports_streaming"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

type inlineKeyboardMarkup struct {
	InlineKeyboardButtons [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type replyKeyboardMarkup struct {
	Keyboard              [][]ReplyKeyboardButton `json:"keyboard"`
	OneTimeKeyboard       bool                    `json:"one_time_keyboard"`
	ResizeKeyboard        bool                    `json:"resize_keyboard"`
	InputFieldPlaceholder string                  `json:"input_field_placeholder"`
	Selective             bool                    `json:"selective"`
}

// replyKeyboardRemove represents an object that if Telegram clients receive,
// they will remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
type replyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
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
	// Optional. Specify True, to send a Pay button.
	// NOTE: This type of button must always be the first button in the first row.
	Pay bool `json:"pay"`
}

// ReplyKeyboardButton represents one button of a reply keyboard.
type ReplyKeyboardButton struct {
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
