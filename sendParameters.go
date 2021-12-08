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
	Duration                 int             `json:"duration"`
	Width                    int             `json:"width"`
	Height                   int             `json:"height"`
	Caption                  string          `json:"caption"`
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

type AudioOptionalParams struct {
	Performer                string          `json:"performer"`
	Title                    string          `json:"title"`
	Duration                 int             `json:"duration"`
	Caption                  string          `json:"caption"`
	ParseMode                string          `json:"parse_mode"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

type DocumentOptionalParams struct {
	Caption string `json:"caption"`
	// Optional. Disables automatic server-side content type
	// detection for files uploaded using multipart/form-data
	DisableContentTypeDetection bool            `json:"disable_content_type_detection"`
	ParseMode                   string          `json:"parse_mode"`
	CaptionEntities             []MessageEntity `json:"caption_entities"`
	DisableNotification         bool            `json:"disable_notification"`
	ReplyToMessageId            int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply    bool            `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

type VoiceOptionalParams struct {
	Duration                 int             `json:"duration"`
	Caption                  string          `json:"caption"`
	ParseMode                string          `json:"parse_mode"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

type AnimationOptionalParams struct {
	Duration                 int             `json:"duration"`
	Width                    int             `json:"width"`
	Height                   int             `json:"height"`
	Caption                  string          `json:"caption"`
	ParseMode                string          `json:"parse_mode"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

type PollOptionalParams struct {
	IsAnonymous bool `json:"is_anonymous"`
	// Poll type, “quiz” or “regular”, defaults to “regular”
	Type string `json:"type"`
	// True, if the poll allows multiple answers,
	// ignored for polls in quiz mode, defaults to False
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	// 0-based identifier of the correct answer option,
	// required for polls in quiz mode
	CorrectOptionId int `json:"correct_option_id"`
	// Text that is shown when a user chooses an
	// incorrect answer or taps on the lamp icon in a
	// quiz-style poll, 0-200 characters with at most 2 line
	// feeds after entities parsing
	Explanation          string          `json:"explanation"`
	ExplanationParseMode string          `json:"explanation_parse_mode"`
	ExplanationEntities  []MessageEntity `json:"explanation_entities"`
	// Amount of time in seconds the poll will be active after
	// creation, 5-600. Can't be used together with close_date.
	OpenPeriod int `json:"open_period"`
	// Point in time (Unix timestamp) when the poll will
	// be automatically closed. Must be at least 5 and no more
	// than 600 seconds in the future.
	// Can't be used together with open_period.
	CloseDate int `json:"close_date"`
	// Pass True, if the poll needs to be immediately closed.
	// This can be useful for poll preview.
	IsClosed                 bool `json:"is_closed"`
	DisableNotification      bool `json:"disable_notification"`
	ReplyToMessageId         int  `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

type DiceOptionalParams struct {
	// Emoji on which the dice throw animation is based.
	// Currently, must be one of “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”.
	// Dice can have values 1-6 for “🎲”, “🎯” and “🎳”,
	// values 1-5 for “🏀” and “⚽”, and values 1-64 for “🎰”.
	// Defaults to “🎲”
	Emoji                    string `json:"emoji"`
	DisableNotification      bool   `json:"disable_notification"`
	ReplyToMessageId         int    `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

type VideoNoteOptionalParams struct {
	Duration                 int  `json:"duration"`
	Length                   int  `json:"length"`
	DisableNotification      bool `json:"disable_notification"`
	ReplyToMessageId         int  `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

type LocationOptionalParams struct {
	HorizontalAccuracy       float64 `json:"horizontal_accuracy"`
	LivePeriod               int     `json:"live_period"`
	Heading                  int     `json:"heading"`
	ProximityAlertRadius     int     `json:"proximity_alert_radius"`
	DisableNotification      bool    `json:"disable_notification"`
	ReplyToMessageId         int     `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool    `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

type ContactOptionalParams struct {
	LastName string `json:"last_name"`
	// Additional data about the contact in the form of a vCard
	Vcard                    string `json:"vcard"`
	DisableNotification      bool   `json:"disable_notification"`
	ReplyToMessageId         int    `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

// MediaGroupOptionalParams represents an album.
type MediaGroupOptionalParams struct {
	ReplyToMessageId         int  `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`
}

type ForwardMessageOptionalParams struct {
	DisableNotification bool `json:"disable_notification"`
}

type CopyMessageOptionalParams struct {
	Caption                  string          `json:"caption"`
	ParseMode                string          `json:"parse_mode"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	InlineKeyboard
	ReplyKeyboard
	ForceReply
}

type GetUserProfilePhotosOptionalParams struct {
	// Sequential number of the first photo to be returned.
	// By default, all photos are returned.
	Offset int `json:"offset"`
	// Limits the number of photos to be retrieved.
	// Values between 1-100 are accepted. Defaults to 100.
	Limit int `json:"limit"`
}

type BanChatMemberOptionalParams struct {
	// Date when the user will be unbanned, unix time.
	// If user is banned for more than 366 days or less
	// than 30 seconds from the current time they are considered to be banned forever.
	// Applied for supergroups and channels only.
	UntilDate int `json:"until_date"`
	// Pass True to delete all messages from the chat for the user that is being removed.
	// If False, the user will be able to see messages in the group that were sent before
	// the user was removed. Always True for supergroups and channels.
	RevokeMessages bool `json:"revoke_messages"`
}

type RestrictChatMemberOptionalParams struct {
	UntilDate int `json:"until_date"`
}

type PromoteChatMemberOptionalParams struct {
	// Pass True, if the administrator's presence in the chat is hidden.
	IsAnonymous bool `json:"is_anonymous"`
	// Pass True, if the administrator can access the chat event log, chat statistics,
	// message statistics in channels, see channel members, see anonymous administrators
	// in supergroups and ignore slow mode. Implied by any other administrator privilege.
	CanManageChat bool `json:"can_manage_chat"`
	// Pass True, if the administrator can create channel posts, channels only.
	CanPostMessages bool `json:"can_post_messages"`
	// Pass True, if the administrator can edit messages of other users and can pin messages, channels only.
	CanEditMessages bool `json:"can_edit_messages"`
	// Pass True, if the administrator can delete messages of other users.
	CanDeleteMessages bool `json:"can_delete_messages"`
	// Pass True, if the administrator can manage voice chats.
	CanManageVoiceChats bool `json:"can_manage_voice_chats"`
	// Pass True, if the administrator can restrict, ban or unban chat members.
	CanRestrictMembers bool `json:"can_restrict_members"`
	// Pass True, if the administrator can add new administrators with a subset of their own privileges
	// or demote administrators that he has promoted, directly or indirectly
	// (promoted by administrators that were appointed by him)
	CanPromoteMembers bool `json:"can_promote_members"`
	// Pass True, if the administrator can change chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info"`
	// Pass True, if the administrator can invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users"`
	// Pass True, if the administrator can pin messages, supergroups only
	CanPinMessages bool `json:"can_pin_messages"`
}

type CreateChatInviteLinkOptionalParams struct {
	ExpireDate  int `json:"expire_date"`
	MemberLimit int `json:"member_limit"`
}

type EditChatInviteLinkOptionalParams struct {
	ExpireDate  int `json:"expire_date"`
	MemberLimit int `json:"member_limit"`
}

type PinChatMessageOptionalParams struct {
	DisableNotification bool `json:"disable_notification"`
}

type UnpinChatMessageOptionalParams struct {
	MessageId int `json:"message_id"`
}

type AnswerCallbackQueryOptionalParams struct {
	text      string
	showAlert bool
	url       string
	cacheTime string
}

type MyCommandsOptionalParams struct {
	// Scope describing scope of users for which the commands are relevant. Defaults to "default".
	Scope        BotCommandScope `json:"scope"`
	LanguageCode string          `json:"language_code"`
}

type EditMessageTextOptionalParams struct {
	ChatId                int             `json:"chat_id"`
	MessageId             int             `json:"message_id"`
	InlineMessageId       int             `json:"inline_message_id"`
	ParseMode             string          `json:"parse_mode"`
	Entities              []MessageEntity `json:"entities"`
	DisableWebPagePreview bool            `json:"disable_web_page_preview"`
	InlineKeyboard
}

type EditMessageCaptionOptionalParams struct {
	ChatId          int             `json:"chat_id"`
	MessageId       int             `json:"message_id"`
	InlineMessageId int             `json:"inline_message_id"`
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	InlineKeyboard
}

type StopPollOptionalParams struct {
	InlineKeyboard
}

type EditMessageMediaOptionalParams struct {
	ChatId          int `json:"chat_id"`
	MessageId       int `json:"message_id"`
	InlineMessageId int `json:"inline_message_id"`
	InlineKeyboard
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
