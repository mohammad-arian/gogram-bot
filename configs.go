package gogram

type KeyboardMarkup interface {
	Add(ButtonKinds)
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

func (t *TextOptionalParams) AddInlineKeyboardButton(i ...InlineKeyboardButton) {
	if t.ReplyMarkup == nil {
		t.ReplyMarkup = &InlineKeyboardMarkup{}
	}
	for _, button := range i {
		b := ButtonKinds{inlineKeyboardButton: button}
		t.ReplyMarkup.Add(b)
	}
}

func (t *TextOptionalParams) AddReplyKeyboardButton(i ...KeyboardButton) {
	if t.ReplyMarkup == nil {
		t.ReplyMarkup = &ReplyKeyboardMarkup{}
	}
	for _, button := range i {
		b := ButtonKinds{keyboardButton: button}
		t.ReplyMarkup.Add(b)
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

type MessageEntity struct {
	Type     string
	offset   int
	length   int
	url      string
	user     User
	language string
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (i *InlineKeyboardMarkup) Add(b ButtonKinds) {
	i.InlineKeyboard = append(i.InlineKeyboard, []InlineKeyboardButton{b.inlineKeyboardButton})
}

type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton `json:"keyboard"`
}

func (r *ReplyKeyboardMarkup) Add(b ButtonKinds) {
	r.Keyboard = append(r.Keyboard, []KeyboardButton{b.keyboardButton})
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

type ButtonKinds struct {
	keyboardButton       KeyboardButton
	inlineKeyboardButton InlineKeyboardButton
}
