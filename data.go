package gogram

import (
	"errors"
	"os"
)

type TextData struct {
	Text                     string          `json:"text"`
	ChatId                   int             `json:"chat_id"`
	ParseMode                string          `json:"parse_mode"`
	Entities                 []MessageEntity `json:"entities"`
	DisableWebPagePreview    bool            `json:"disable_web_page_preview"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	Keyboard
}

func (t TextData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Text": t.Text, "ChatId": t.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("sendMessage", b, t, &MessageResponse{})
	return res.(*MessageResponse), err
}

type PhotoData struct {
	Photo                    interface{}     `json:"photo"`
	ChatId                   int             `json:"chat_id"`
	ParseMode                string          `json:"parse_mode"`
	Caption                  string          `json:"caption"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	Keyboard
}

func (p PhotoData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Photo": p.Photo, "ChatId": p.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("sendPhoto", b, p, &MessageResponse{})
	return res.(*MessageResponse), err
}

type VideoData struct {
	ChatId                   int             `json:"chat_id"`
	Video                    interface{}     `json:"video"`
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
	Keyboard
}

func (v VideoData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Video": v.Video, "ChatId": v.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("sendVideo", b, v, &MessageResponse{})
	return res.(*MessageResponse), err
}

type AudioData struct {
	ChatId                   int             `json:"chat_id"`
	Audio                    interface{}     `json:"audio"`
	Performer                string          `json:"performer"`
	Title                    string          `json:"title"`
	Duration                 int             `json:"duration"`
	Caption                  string          `json:"caption"`
	ParseMode                string          `json:"parse_mode"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	Keyboard
}

func (a AudioData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Audio": a.Audio, "ChatId": a.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("sendAudio", b, a, &MessageResponse{})
	return res.(*MessageResponse), err
}

type DocumentData struct {
	ChatId int `json:"chat_id"`
	// file to send. Pass a file_id as string to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a string for Telegram to get a file from the Internet, or pass a *os.File
	Document interface{} `json:"document"`
	Caption  string      `json:"caption"`
	// Optional. Disables automatic server-side content type
	// detection for files uploaded using multipart/form-data
	DisableContentTypeDetection bool            `json:"disable_content_type_detection"`
	ParseMode                   string          `json:"parse_mode"`
	CaptionEntities             []MessageEntity `json:"caption_entities"`
	DisableNotification         bool            `json:"disable_notification"`
	ReplyToMessageId            int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply    bool            `json:"allow_sending_without_reply"`
	Keyboard
}

func (d DocumentData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Document": d.Document, "ChatId": d.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("sendDocument", b, d, &MessageResponse{})
	return res.(*MessageResponse), err
}

type VoiceData struct {
	ChatId                   int             `json:"chat_id"`
	Voice                    interface{}     `json:"voice"`
	Duration                 int             `json:"duration"`
	Caption                  string          `json:"caption"`
	ParseMode                string          `json:"parse_mode"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	Keyboard
}

func (v VoiceData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Voice": v.Voice, "ChatId": v.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("sendVoice", b, v, &MessageResponse{})
	return res.(*MessageResponse), err
}

type AnimationData struct {
	ChatId                   int             `json:"chat_id"`
	Animation                interface{}     `json:"animation"`
	Duration                 int             `json:"duration"`
	Width                    int             `json:"width"`
	Height                   int             `json:"height"`
	Caption                  string          `json:"caption"`
	ParseMode                string          `json:"parse_mode"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	Keyboard
}

func (a AnimationData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Animation": a.Animation, "ChatId": a.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("sendAnimation", b, a, &MessageResponse{})
	return res.(*MessageResponse), err
}

type PollData struct {
	ChatId      int      `json:"chat_id"`
	Question    string   `json:"question"`
	Options     []string `json:"options"`
	IsAnonymous bool     `json:"is_anonymous"`
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
	Keyboard
}

func (p PollData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Question": p.Question, "ChatId": p.ChatId, "Options": p.Options}); err != nil {
		return nil, err
	}
	res, err := request("sendPoll", b, p, &MessageResponse{})
	return res.(*MessageResponse), err
}

type DiceData struct {
	ChatId int `json:"chat_id"`
	// Emoji on which the dice throw animation is based.
	// Currently, must be one of “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”.
	// Dice can have values 1-6 for “🎲”, “🎯” and “🎳”,
	// values 1-5 for “🏀” and “⚽”, and values 1-64 for “🎰”.
	// Defaults to “🎲”
	Emoji                    string `json:"emoji"`
	DisableNotification      bool   `json:"disable_notification"`
	ReplyToMessageId         int    `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply"`
	Keyboard
}

func (d DiceData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Emoji": d.Emoji, "ChatId": d.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("sendDice", b, d, &MessageResponse{})
	return res.(*MessageResponse), err
}

type VideoNoteData struct {
	ChatId                   int         `json:"chat_id"`
	VideoNote                interface{} `json:"videoNote"`
	Duration                 int         `json:"duration"`
	Length                   int         `json:"length"`
	DisableNotification      bool        `json:"disable_notification"`
	ReplyToMessageId         int         `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"`
	Keyboard
}

func (v VideoNoteData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"VideoNote": v.VideoNote, "ChatId": v.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("sendVideoNote", b, v, &MessageResponse{})
	return res.(*MessageResponse), err
}

type LocationData struct {
	ChatId                   int     `json:"chat_id"`
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	HorizontalAccuracy       float64 `json:"horizontal_accuracy"`
	LivePeriod               int     `json:"live_period"`
	Heading                  int     `json:"heading"`
	ProximityAlertRadius     int     `json:"proximity_alert_radius"`
	DisableNotification      bool    `json:"disable_notification"`
	ReplyToMessageId         int     `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool    `json:"allow_sending_without_reply"`
	Keyboard
}

func (l LocationData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": l.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("sendLocation", b, l, &MessageResponse{})
	return res.(*MessageResponse), err
}

type ContactData struct {
	ChatId      int    `json:"chat_id"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	// Additional data about the contact in the form of a vCard
	Vcard                    string `json:"vcard"`
	DisableNotification      bool   `json:"disable_notification"`
	ReplyToMessageId         int    `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply"`
	Keyboard
}

func (c ContactData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"PhoneNumber": c.PhoneNumber, "ChatId": c.ChatId,
		"FirstName": c.FirstName}); err != nil {
		return nil, err
	}
	res, err := request("sendContact", b, c, &MessageResponse{})
	return res.(*MessageResponse), err
}

// MediaGroupData represents an album.
type MediaGroupData struct {
	ChatId int          `json:"chat_id"`
	Media  []InputMedia `json:"media"`
	// leave this field. it will be set automatically.
	File                     []*os.File
	ReplyToMessageId         int  `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`
}

func (m MediaGroupData) Send(b Bot) (response *SliceMessageResponse, err error) {
	if len(m.Media) == 0 {
		return &SliceMessageResponse{}, errors.New("media slice is empty. pass media a slice of structs of type " +
			"InputMediaPhoto, InputMediaVideo, InputMediaDocument or InputMediaAudio")
	}
	for _, j := range m.Media {
		if j.checkInputMedia(&m.File) != nil {
			return &SliceMessageResponse{}, err
		}
	}
	res, err := request("sendMediaGroup", b, m, &SliceMessageResponse{})
	return res.(*SliceMessageResponse), err
}

type ForwardMessageData struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Unique identifier for the chat where the original message was
	// sent (or channel username in the format @channelusername)
	FromChatId int `json:"from_chat_id"`
	// message identifier in the chat specified in from_chat_id
	MessageId           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification"`
	ProtectContent      bool `json:"protect_content"`
}

func (f ForwardMessageData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"FromChatId": f.FromChatId, "ChatId": f.ChatId,
		"MessageId": f.MessageId}); err != nil {
		return nil, err
	}
	res, err := request("forwardMessage", b, f, &MessageResponse{})
	return res.(*MessageResponse), err
}

type CopyMessageData struct {
	ChatId                   int             `json:"chat_id"`
	FromChatId               int             `json:"from_chat_id"`
	MessageId                int             `json:"message_id"`
	Caption                  string          `json:"caption"`
	ParseMode                string          `json:"parse_mode"`
	CaptionEntities          []MessageEntity `json:"caption_entities"`
	DisableNotification      bool            `json:"disable_notification"`
	ReplyToMessageId         int             `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	Keyboard
}

func (c CopyMessageData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"FromChatId": c.FromChatId, "ChatId": c.ChatId,
		"MessageId": c.MessageId}); err != nil {
		return nil, err
	}
	res, err := request("copyMessage", b, c, &MessageResponse{})
	return res.(*MessageResponse), err
}

type DeleteMessageData struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}

func (d DeleteMessageData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": d.ChatId,
		"MessageId": d.MessageId}); err != nil {
		return nil, err
	}
	res, err := request("deleteMessage", b, d, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type DeleteChatStickerSetData struct {
	ChatId int `json:"chat_id"`
}

func (d DeleteChatStickerSetData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": d.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("deleteChatStickerSet", b, d, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type SetChatStickerSetData struct {
	ChatId         int    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

func (s SetChatStickerSetData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId,
		"StickerSetName": s.StickerSetName}); err != nil {
		return nil, err
	}
	res, err := request("setChatStickerSet", b, s, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type GetChatMemberData struct {
	ChatId int `json:"chat_id"`
	UserId int `json:"user_id"`
}

func (g GetChatMemberData) Send(b Bot) (response *ChatMemberResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": g.ChatId,
		"UserId": g.UserId}); err != nil {
		return nil, err
	}
	res, err := request("getChatMember", b, g, &ChatMemberResponse{})
	return res.(*ChatMemberResponse), err
}

type GetChatMemberCountData struct {
	ChatId int `json:"chat_id"`
}

func (g GetChatMemberCountData) Send(b Bot) (response *IntResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": g.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("getChatMemberCount", b, g, &IntResponse{})
	return res.(*IntResponse), err
}

type GetChatAdministratorsData struct {
	ChatId int `json:"chat_id"`
}

func (g GetChatAdministratorsData) Send(b Bot) (response *ChatMemberResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": g.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("getChatAdministrators", b, g, &ChatMemberResponse{})
	member := res.(*ChatMemberResponse)
	member.permissionSetter()
	return member, err
}

type GetChatData struct {
	ChatId int `json:"chat_id"`
}

func (g GetChatData) Send(b Bot) (response *ChatResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": g.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("getChat", b, g, &ChatResponse{})
	return res.(*ChatResponse), err
}

type LeaveData struct {
	ChatId int `json:"chat_id"`
}

func (l LeaveData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": l.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("leaveChat", b, l, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type UnpinAllChatMessagesData struct {
	ChatId int `json:"chat_id"`
}

func (u UnpinAllChatMessagesData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": u.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("unpinAllChatMessages", b, u, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type SetChatDescriptionData struct {
	ChatId      int    `json:"chat_id"`
	Description string `json:"description"`
}

func (s SetChatDescriptionData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId,
		"Description": s.Description}); err != nil {
		return nil, err
	}
	res, err := request("setChatDescription", b, s, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type SetChatTitleData struct {
	ChatId int    `json:"chat_id"`
	Title  string `json:"title"`
}

func (s SetChatTitleData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId,
		"Title": s.Title}); err != nil {
		return nil, err
	}
	res, err := request("setChatTitle", b, s, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type DeleteChatPhotoData struct {
	ChatId int `json:"chat_id"`
}

func (d DeleteChatPhotoData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": d.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("deleteChatPhoto", b, d, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type SetChatPhotoData struct {
	ChatId int      `json:"chat_id"`
	Photo  *os.File `json:"photo"`
}

func (s SetChatPhotoData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "Photo": s.Photo}); err != nil {
		return nil, err
	}
	res, err := request("setChatPhoto", b, s, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type RevokeChatInviteLinkData struct {
	ChatId     int    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

func (r RevokeChatInviteLinkData) Send(b Bot) (response *InviteLinkResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": r.ChatId,
		"InviteLink": r.InviteLink}); err != nil {
		return nil, err
	}
	res, err := request("revokeChatInviteLink", b, r, &InviteLinkResponse{})
	return res.(*InviteLinkResponse), err
}

type ExportChatInviteLinkData struct {
	ChatId int `json:"chat_id"`
}

func (e ExportChatInviteLinkData) Send(b Bot) (response *MapResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": e.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("exportChatInviteLink", b, e, &MapResponse{})
	return res.(*MapResponse), err
}

type SendChatActionData struct {
	ChatId int    `json:"chat_id"`
	Action string `json:"action"`
}

func (s SendChatActionData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId}); err != nil {
		return nil, err
	}
	var actions = []string{"typing", "upload_photo", "record_video", "upload_video", "general",
		"upload_document", "upload_voice", "record_voice", "find_location", "record_video_note", "upload_video_note"}
	for _, a := range actions {
		if a == s.Action {
			res, err := request("sendChatAction", b, s, &BooleanResponse{})
			return res.(*BooleanResponse), err
		}
	}
	return nil, errors.New(s.Action + " is an unknown action, read the document.")
}

type GetFileData struct {
	FileId string `json:"file_id"`
}

func (g GetFileData) Send(b Bot) (response *FileResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"FileId": g.FileId}); err != nil {
		return nil, err
	}
	res, err := request("getFile", b, g, &FileResponse{})
	return res.(*FileResponse), err
}

type UnbanChatMemberData struct {
	ChatId       int  `json:"chat_id"`
	UserId       int  `json:"user_id"`
	OnlyIfBanned bool `json:"only_if_banned"`
}

func (u UnbanChatMemberData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": u.ChatId, "UserId": u.UserId}); err != nil {
		return nil, err
	}
	res, err := request("unbanChatMember", b, u, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type SetChatAdministratorCustomTitleData struct {
	ChatId      int    `json:"chat_id"`
	UserId      int    `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

func (s SetChatAdministratorCustomTitleData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "UserId": s.UserId,
		"CustomTitle": s.CustomTitle}); err != nil {
		return nil, err
	}
	res, err := request("setChatAdministratorCustomTitle", b, s, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type SetChatPermissionsData struct {
	ChatId      int             `json:"chat_id"`
	Permissions ChatPermissions `json:"permissions"`
}

func (s SetChatPermissionsData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("setChatPermissions", b, s, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type UserProfilePhotosData struct {
	UserId int `json:"user_id"`
	// Sequential number of the first photo to be returned.
	// By default, all photos are returned.
	Offset int `json:"offset"`
	// Limits the number of photos to be retrieved.
	// Values between 1-100 are accepted. Defaults to 100.
	Limit int `json:"limit"`
}

func (u UserProfilePhotosData) Send(b Bot) (response *UserProfileResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"UserId": u.UserId}); err != nil {
		return nil, err
	}
	res, err := request("getUserProfilePhotos", b, u, &UserProfileResponse{})
	return res.(*UserProfileResponse), err
}

type BanChatMemberData struct {
	ChatId int `json:"chat_id"`
	UserId int `json:"user_id"`
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

func (ban BanChatMemberData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"UserId": ban.UserId, "ChatId": ban.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("banChatMember", b, ban, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type RestrictChatMemberData struct {
	ChatId      int             `json:"chat_id"`
	UserId      int             `json:"user_id"`
	Permissions ChatPermissions `json:"permissions"`
	UntilDate   int             `json:"until_date"`
}

func (r RestrictChatMemberData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"UserId": r.UserId, "ChatId": r.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("restrictChatMember", b, r, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type PromoteChatMemberData struct {
	ChatId int `json:"chat_id"`
	UserId int `json:"user_id"`
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

func (p PromoteChatMemberData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"UserId": p.UserId, "ChatId": p.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("promoteChatMember", b, p, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type CreateChatInviteLinkData struct {
	ChatId      int `json:"chat_id"`
	ExpireDate  int `json:"expire_date"`
	MemberLimit int `json:"member_limit"`
}

func (c CreateChatInviteLinkData) Send(b Bot) (response *InviteLinkResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": c.ChatId}); err != nil {
		return nil, err
	}
	res, err := request("createChatInviteLink", b, c, &InviteLinkResponse{})
	return res.(*InviteLinkResponse), err
}

type EditChatInviteLinkData struct {
	ChatId      int    `json:"chat_id"`
	InviteLink  string `json:"invite_link"`
	ExpireDate  int    `json:"expire_date"`
	MemberLimit int    `json:"member_limit"`
}

func (e EditChatInviteLinkData) Send(b Bot) (response *InviteLinkResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": e.ChatId,
		"InviteLink": e.InviteLink}); err != nil {
		return nil, err
	}
	res, err := request("editChatInviteLink", b, e, &InviteLinkResponse{})
	return res.(*InviteLinkResponse), err
}

type PinChatMessageData struct {
	ChatId              int  `json:"chat_id"`
	MessageId           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification"`
}

func (p PinChatMessageData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": p.ChatId,
		"MessageId": p.MessageId}); err != nil {
		return nil, err
	}
	res, err := request("pinChatMessage", b, p, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type UnpinChatMessageData struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}

func (u UnpinChatMessageData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": u.ChatId,
		"MessageId": u.MessageId}); err != nil {
		return nil, err
	}
	res, err := request("unpinChatMessage", b, u, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type AnswerCallbackQueryData struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text"`
	ShowAlert       bool   `json:"show_alert"`
	Url             string `json:"url"`
	CacheTime       string `json:"cache_time"`
}

func (a AnswerCallbackQueryData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"CallbackQueryId": a.CallbackQueryId}); err != nil {
		return nil, err
	}
	res, err := request("answerCallbackQuery", b, a, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type SetMyCommandsData struct {
	Commands []BotCommand `json:"commands"`
	// Scope describing scope of users for which the commands are relevant. Defaults to "default".
	Scope        BotCommandScope `json:"scope"`
	LanguageCode string          `json:"language_code"`
}

func (s SetMyCommandsData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Commands": s.Commands}); err != nil {
		return nil, err
	}
	types := map[string]bool{"default": true, "chat_member": true, "all_private_chats": true,
		"all_group_chats": true, "all_chat_administrators": true, "chat": true, "chat_administrators": true}
	if _, ok := types[s.Scope.Type]; ok == false {
		s.Scope.Type = "default"
	}
	res, err := request("setMyCommands", b, s, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type DeleteMyCommandsData struct {
	// Scope describing scope of users for which the commands are relevant. Defaults to "default".
	Scope        BotCommandScope `json:"scope"`
	LanguageCode string          `json:"language_code"`
}

func (d DeleteMyCommandsData) Send(b Bot) (response *BooleanResponse, err error) {
	types := map[string]bool{"default": true, "chat_member": true, "all_private_chats": true,
		"all_group_chats": true, "all_chat_administrators": true, "chat": true, "chat_administrators": true}
	if _, ok := types[d.Scope.Type]; ok == false {
		d.Scope.Type = "default"
	}
	res, err := request("deleteMyCommands", b, d, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type GetMyCommandsData struct {
	// Scope describing scope of users for which the commands are relevant. Defaults to "default".
	Scope        BotCommandScope `json:"scope"`
	LanguageCode string          `json:"language_code"`
}

func (m GetMyCommandsData) Send(b Bot) (response *BotCommandResponse, err error) {
	res, err := request("getMyCommands", b, m, &BooleanResponse{})
	return res.(*BotCommandResponse), err
}

type EditMessageTextData struct {
	Text                  string          `json:"text"`
	InlineMessageId       string          `json:"inline_message_id"`
	ChatId                int             `json:"chat_id"`
	MessageId             int             `json:"message_id"`
	ParseMode             string          `json:"parse_mode"`
	Entities              []MessageEntity `json:"entities"`
	DisableWebPagePreview bool            `json:"disable_web_page_preview"`
	InlineKeyboard
}

func (e EditMessageTextData) Send(b Bot) (response *MapResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Text": e.Text}); err != nil {
		return nil, err
	}
	if e.InlineMessageId == "" {
		if e.ChatId == 0 && e.MessageId == 0 {
			return nil, errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	res, err := request("editMessageText", b, e, &MapResponse{})
	return res.(*MapResponse), err
}

type EditMessageCaptionData struct {
	ChatId          int             `json:"chat_id"`
	MessageId       int             `json:"message_id"`
	InlineMessageId string          `json:"inline_message_id"`
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	InlineKeyboard
}

func (e EditMessageCaptionData) Send(b Bot) (response *MapResponse, err error) {
	if e.InlineMessageId == "" {
		if e.ChatId == 0 && e.MessageId == 0 {
			return nil, errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	res, err := request("editMessageCaption", b, e, &MapResponse{})
	return res.(*MapResponse), err
}

type EditMessageReplyMarkupData struct {
	ChatId          int    `json:"chat_id"`
	MessageId       int    `json:"message_id"`
	InlineMessageId string `json:"inline_message_id"`
	InlineKeyboard
}

func (e EditMessageReplyMarkupData) Send(b Bot) (response *MapResponse, err error) {
	if e.InlineMessageId == "" {
		if e.ChatId == 0 && e.MessageId == 0 {
			return nil, errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	res, err := request("editMessageReplyMarkup", b, e, &MapResponse{})
	return res.(*MapResponse), err
}

type StopPollData struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
	InlineKeyboard
}

func (s StopPollData) Send(b Bot) (response *PollResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "MessageId": s.MessageId}); err != nil {
		return nil, err
	}
	res, err := request("stopPoll", b, s, &PollResponse{})
	return res.(*PollResponse), err
}

type EditMessageMediaData struct {
	InlineMessageId string     `json:"inline_message_id"`
	Media           InputMedia `json:"media"`
	ChatId          int        `json:"chat_id"`
	MessageId       int        `json:"message_id"`
	File            []*os.File
	InlineKeyboard
}

func (e EditMessageMediaData) Send(b Bot) (response *MapResponse, err error) {
	if e.InlineMessageId == "" {
		if e.ChatId == 0 && e.MessageId == 0 {
			return nil, errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	if e.Media == nil {
		return nil, errors.New("media is nil. pass media a struct of type " +
			"InputMediaPhoto, InputMediaVideo, InputMediaDocument or InputMediaAudio")
	}
	if e.Media.checkInputMedia(&e.File) != nil {
		return nil, err
	}
	res, err := request("editMessageMedia", b, e, &MapResponse{})
	return res.(*MapResponse), err
}

type SetWebhookData struct {
	Url                string   `json:"url"`
	Certificate        *os.File `json:"certificate"`
	IpAddress          string   `json:"ip_address"`
	MaxConnections     int      `json:"max_connections"`
	AllowedUpdates     []string `json:"allowed_updates"`
	DropPendingUpdates bool     `json:"drop_pending_updates"`
}

func (s SetWebhookData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Url": s.Url}); err != nil {
		return nil, err
	}
	res, err := request("setWebhook", b, s, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type SendStickerData struct {
	ChatId                   int         `json:"chat_id"`
	Sticker                  interface{} `json:"sticker"`
	DisableNotification      bool        `json:"disable_notification"`
	ReplyToMessageId         int         `json:"reply_To_Message_Id"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"`
	Keyboard
}

func (s SendStickerData) Send(b Bot) (response *MessageResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "Sticker": s.Sticker}); err != nil {
		return nil, err
	}
	res, err := request("sendSticker", b, s, &MessageResponse{})
	return res.(*MessageResponse), err
}

type DeleteStickerFromSetData struct {
	Sticker string `json:"sticker"`
}

func (d DeleteStickerFromSetData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Sticker": d.Sticker}); err != nil {
		return nil, err
	}
	res, err := request("deleteStickerFromSet", b, d, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type SetStickerPositionInSetData struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

func (s SetStickerPositionInSetData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Sticker": s.Sticker, "Position": s.Position}); err != nil {
		return nil, err
	}
	res, err := request("setStickerPositionInSet", b, s, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type UploadStickerFileData struct {
	UserId     int      `json:"user_id"`
	PngSticker *os.File `json:"png_sticker"`
}

func (u UploadStickerFileData) Send(b Bot) (response *FileResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"UserId": u.UserId, "PngSticker": u.PngSticker}); err != nil {
		return nil, err
	}
	res, err := request("uploadStickerFile", b, u, &FileResponse{})
	return res.(*FileResponse), err
}

type GetStickerSetData struct {
	Name string `json:"name"`
}

func (g GetStickerSetData) Send(b Bot) (response *StickerSetResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"Name": g.Name}); err != nil {
		return nil, err
	}
	res, err := request("getStickerSet", b, g, &StickerSetResponse{})
	return res.(*StickerSetResponse), err
}

type CreateNewStickerSetData struct {
	UserId        int          `json:"user_id"`
	Name          string       `json:"name"`
	Title         string       `json:"title"`
	Emojis        string       `json:"emojis"`
	PngSticker    interface{}  `json:"png_sticker"`
	TgsSticker    *os.File     `json:"tgs_sticker"`
	ContainsMasks bool         `json:"contains_masks"`
	MaskPosition  MaskPosition `json:"mask_position"`
}

func (c CreateNewStickerSetData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"UserId": c.UserId, "Name": c.Name,
		"Title": c.Title, "Emojis": c.Emojis}); err != nil {
		return nil, err
	}
	res, err := request("createNewStickerSet", b, c, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type AddStickerToSetData struct {
	UserId       int          `json:"user_id"`
	Name         string       `json:"name"`
	Emojis       string       `json:"emojis"`
	PngSticker   interface{}  `json:"png_sticker"`
	TgsSticker   *os.File     `json:"tgs_sticker"`
	MaskPosition MaskPosition `json:"mask_position"`
}

func (a AddStickerToSetData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"UserId": a.UserId, "Name": a.Name,
		"Emojis": a.Emojis}); err != nil {
		return nil, err
	}
	res, err := request("addStickerToSet", b, a, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type SetStickerSetThumbData struct {
	UserId int         `json:"user_id"`
	Name   string      `json:"name"`
	Thumb  interface{} `json:"thumb"`
}

func (s SetStickerSetThumbData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"UserId": s.UserId, "Name": s.Name}); err != nil {
		return nil, err
	}
	res, err := request("setStickerSetThumb", b, s, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

type AnswerInlineQueryData struct {
	InlineQueryId     string        `json:"inline_query_id"`
	Results           []QueryAnswer `json:"results"`
	CacheTime         int           `json:"cache_time"`
	IsPersonal        bool          `json:"is_personal"`
	NextOffset        string        `json:"next_offset"`
	SwitchPmText      string        `json:"switch_pm_text"`
	SwitchPmParameter string        `json:"switch_pm_parameter"`
}

func (a AnswerInlineQueryData) Send(b Bot) (response *BooleanResponse, err error) {
	if err = globalEmptyFieldChecker(map[string]interface{}{"InlineQueryId": a.InlineQueryId}); err != nil {
		return nil, err
	}
	if len(a.Results) == 0 {
		return &BooleanResponse{}, errors.New("results slice is empty. pass QueryAnswer structs such as " +
			"InlineQueryResultArticle, InlineQueryResultPhoto and etc")
	}
	for _, j := range a.Results {
		if e := j.checkQueryAnswer(); e != nil {
			return &BooleanResponse{}, e
		}
	}
	res, err := request("answerInlineQuery", b, a, &BooleanResponse{})
	return res.(*BooleanResponse), err
}
