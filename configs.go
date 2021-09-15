package gogram

// TextOptionalParams represents optional parameters
// that SendText function can use.
// All fields are optional.
type TextOptionalParams struct {
	ParseMode                string                   `json:"parse_mode"`
	Entities                 []MessageEntity          `json:"entities"`
	DisableWebPagePreview    bool                     `json:"disable_web_page_preview"`
	DisableNotification      bool                     `json:"disable_notification"`
	ReplyToMessageId         int                      `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool                     `json:"allow_sending_without_reply"`
	InlineKeyboardButtons    [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (t *TextOptionalParams) AddInlineKeyboardButton(i ...InlineKeyboardButton) {
	for _, button := range i {
		t.InlineKeyboardButtons = append(t.InlineKeyboardButtons, []InlineKeyboardButton{button})
	}
}

type MessageEntity struct {
	Type     string
	offset   int
	length   int
	url      string
	user     User
	language string
}

type InlineKeyboardMarkup struct {
	InlineKeyboardButtons [][]InlineKeyboardButton `json:"inline_keyboard"`
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
