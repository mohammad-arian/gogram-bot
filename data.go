package gogram

import (
	"errors"
	"os"
)

type Method interface {
	// check() is used in Request
	check() error
	Send(b Bot) (Response, error)
}

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

func (t TextData) Send(b Bot) (response Response, err error) {
	return Request("sendMessage", b, t, &ResponseImpl{Result: &Message{}})
}

func (t TextData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Text": t.Text, "ChatId": t.ChatId})
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

func (p PhotoData) Send(b Bot) (response Response, err error) {
	return Request("sendPhoto", b, p, &ResponseImpl{Result: &Message{}})
}

func (p PhotoData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Photo": p.Photo, "ChatId": p.ChatId})
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

func (v VideoData) Send(b Bot) (response Response, err error) {
	return Request("sendVideo", b, v, &ResponseImpl{Result: &Message{}})
}

func (v VideoData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Video": v.Video, "ChatId": v.ChatId})
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

func (a AudioData) Send(b Bot) (response Response, err error) {
	return Request("sendAudio", b, a, &ResponseImpl{Result: &Message{}})
}
func (a AudioData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Audio": a.Audio, "ChatId": a.ChatId})
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

func (d DocumentData) Send(b Bot) (response Response, err error) {
	return Request("sendDocument", b, d, &ResponseImpl{Result: &Message{}})
}
func (d DocumentData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Document": d.Document, "ChatId": d.ChatId})
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

func (v VoiceData) Send(b Bot) (response Response, err error) {
	return Request("sendVoice", b, v, &ResponseImpl{Result: &Message{}})
}
func (v VoiceData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Voice": v.Voice, "ChatId": v.ChatId})
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

func (a AnimationData) Send(b Bot) (response Response, err error) {
	return Request("sendAnimation", b, a, &ResponseImpl{Result: &Message{}})
}
func (a AnimationData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Animation": a.Animation, "ChatId": a.ChatId})
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

func (p PollData) Send(b Bot) (response Response, err error) {
	return Request("sendPoll", b, p, &ResponseImpl{Result: &Message{}})
}
func (p PollData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Question": p.Question, "ChatId": p.ChatId,
		"Options": p.Options})
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

func (d DiceData) Send(b Bot) (response Response, err error) {
	return Request("sendDice", b, d, &ResponseImpl{Result: &Message{}})
}
func (d DiceData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Emoji": d.Emoji, "ChatId": d.ChatId})
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

func (v VideoNoteData) Send(b Bot) (response Response, err error) {
	return Request("sendVideoNote", b, v, &ResponseImpl{Result: &Message{}})
}
func (v VideoNoteData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"VideoNote": v.VideoNote, "ChatId": v.ChatId})
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

func (l LocationData) Send(b Bot) (response Response, err error) {
	return Request("sendLocation", b, l, &ResponseImpl{Result: &Message{}})
}
func (l LocationData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": l.ChatId})
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

func (c ContactData) Send(b Bot) (response Response, err error) {
	return Request("sendContact", b, c, &ResponseImpl{Result: &Message{}})
}
func (c ContactData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"PhoneNumber": c.PhoneNumber, "ChatId": c.ChatId,
		"FirstName": c.FirstName})
}

// MediaGroupData represents an album.
type MediaGroupData struct {
	ChatId int          `json:"chat_id"`
	Media  []InputMedia `json:"media"`
	// leave this field. it will be set automatically.
	Files                    []*os.File
	ReplyToMessageId         int  `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`
}

func (m MediaGroupData) Send(b Bot) (response Response, err error) {
	for _, j := range m.Media {
		m.Files = append(m.Files, j.returnFile())
	}
	return Request("sendMediaGroup", b, m, &ResponseImpl{Result: &[]Message{}})
}
func (m MediaGroupData) check() error {
	if len(m.Media) == 0 {
		return errors.New("media slice is empty. pass media a slice of structs of type " +
			"InputMediaPhoto, InputMediaVideo, InputMediaDocument or InputMediaAudio")
	}
	return nil
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

func (f ForwardMessageData) Send(b Bot) (response Response, err error) {
	return Request("forwardMessage", b, f, &ResponseImpl{Result: &Message{}})
}
func (f ForwardMessageData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"FromChatId": f.FromChatId, "ChatId": f.ChatId,
		"MessageId": f.MessageId})
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

func (c CopyMessageData) Send(b Bot) (response Response, err error) {
	return Request("copyMessage", b, c, &ResponseImpl{Result: &Message{}})
}
func (c CopyMessageData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"FromChatId": c.FromChatId, "ChatId": c.ChatId,
		"MessageId": c.MessageId})
}

type DeleteMessageData struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}

func (d DeleteMessageData) Send(b Bot) (response Response, err error) {
	return Request("deleteMessage", b, d, &ResponseImpl{Result: &Boolean{}})
}
func (d DeleteMessageData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": d.ChatId, "MessageId": d.MessageId})
}

type DeleteChatStickerSetData struct {
	ChatId int `json:"chat_id"`
}

func (d DeleteChatStickerSetData) Send(b Bot) (response Response, err error) {
	return Request("deleteChatStickerSet", b, d, &ResponseImpl{Result: &Boolean{}})
}
func (d DeleteChatStickerSetData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": d.ChatId})
}

type SetChatStickerSetData struct {
	ChatId         int    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

func (s SetChatStickerSetData) Send(b Bot) (response Response, err error) {
	return Request("setChatStickerSet", b, s, &ResponseImpl{Result: &Boolean{}})
}
func (s SetChatStickerSetData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "StickerSetName": s.StickerSetName})
}

type GetChatMemberData struct {
	ChatId int `json:"chat_id"`
	UserId int `json:"user_id"`
}

func (g GetChatMemberData) Send(b Bot) (response Response, err error) {
	return Request("getChatMember", b, g, &ResponseImpl{Result: &[]ChatMember{}})
}
func (g GetChatMemberData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": g.ChatId, "UserId": g.UserId})
}

type GetChatMemberCountData struct {
	ChatId int `json:"chat_id"`
}

func (g GetChatMemberCountData) Send(b Bot) (response Response, err error) {
	return Request("getChatMemberCount", b, g, &ResponseImpl{Result: &Int{}})
}
func (g GetChatMemberCountData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": g.ChatId})
}

type GetChatAdministratorsData struct {
	ChatId int `json:"chat_id"`
}

func (g GetChatAdministratorsData) Send(b Bot) (response Response, err error) {
	return Request("getChatAdministrators", b, g, &ResponseImpl{})
}
func (g GetChatAdministratorsData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": g.ChatId})
}

type GetChatData struct {
	ChatId int `json:"chat_id"`
}

func (g GetChatData) Send(b Bot) (response Response, err error) {
	return Request("getChat", b, g, &ResponseImpl{Result: &[]ChatMember{}})
}
func (g GetChatData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": g.ChatId})
}

type LeaveData struct {
	ChatId int `json:"chat_id"`
}

func (l LeaveData) Send(b Bot) (response Response, err error) {
	return Request("leaveChat", b, l, &ResponseImpl{Result: &Boolean{}})
}
func (l LeaveData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": l.ChatId})
}

type UnpinAllChatMessagesData struct {
	ChatId int `json:"chat_id"`
}

func (u UnpinAllChatMessagesData) Send(b Bot) (response Response, err error) {
	return Request("unpinAllChatMessages", b, u, &ResponseImpl{Result: &Boolean{}})
}
func (u UnpinAllChatMessagesData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": u.ChatId})
}

type SetChatDescriptionData struct {
	ChatId      int    `json:"chat_id"`
	Description string `json:"description"`
}

func (s SetChatDescriptionData) Send(b Bot) (response Response, err error) {
	return Request("setChatDescription", b, s, &ResponseImpl{Result: &Boolean{}})
}
func (s SetChatDescriptionData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "Description": s.Description})
}

type SetChatTitleData struct {
	ChatId int    `json:"chat_id"`
	Title  string `json:"title"`
}

func (s SetChatTitleData) Send(b Bot) (response Response, err error) {
	return Request("setChatTitle", b, s, &ResponseImpl{Result: &Boolean{}})
}
func (s SetChatTitleData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "Title": s.Title})
}

type DeleteChatPhotoData struct {
	ChatId int `json:"chat_id"`
}

func (d DeleteChatPhotoData) Send(b Bot) (response Response, err error) {
	return Request("deleteChatPhoto", b, d, &ResponseImpl{Result: &Boolean{}})
}
func (d DeleteChatPhotoData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": d.ChatId})
}

type SetChatPhotoData struct {
	ChatId int      `json:"chat_id"`
	Photo  *os.File `json:"photo"`
}

func (s SetChatPhotoData) Send(b Bot) (response Response, err error) {
	return Request("setChatPhoto", b, s, &ResponseImpl{Result: &Boolean{}})
}
func (s SetChatPhotoData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "Photo": s.Photo})
}

type RevokeChatInviteLinkData struct {
	ChatId     int    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

func (r RevokeChatInviteLinkData) Send(b Bot) (response Response, err error) {
	return Request("revokeChatInviteLink", b, r, &ResponseImpl{Result: &Boolean{}})
}
func (r RevokeChatInviteLinkData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": r.ChatId, "InviteLink": r.InviteLink})
}

type ExportChatInviteLinkData struct {
	ChatId int `json:"chat_id"`
}

func (e ExportChatInviteLinkData) Send(b Bot) (response Response, err error) {
	return Request("exportChatInviteLink", b, e, &ResponseImpl{Result: &Boolean{}})
}
func (e ExportChatInviteLinkData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": e.ChatId})
}

type SendChatActionData struct {
	ChatId int    `json:"chat_id"`
	Action string `json:"action"`
}

func (s SendChatActionData) Send(b Bot) (response Response, err error) {
	return Request("sendChatAction", b, s, &ResponseImpl{Result: &Boolean{}})
}
func (s SendChatActionData) check() error {
	var actions = map[string]bool{"typing": true, "upload_photo": true, "record_video": true, "upload_video": true,
		"general": true, "upload_document": true, "upload_voice": true, "record_voice": true, "find_location": true,
		"record_video_note": true, "upload_video_note": true}
	if _, ok := actions[s.Action]; ok == false {
		return errors.New("action is unknown, read the document")
	}
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId})
}

type GetFileData struct {
	FileId string `json:"file_id"`
}

func (g GetFileData) Send(b Bot) (response Response, err error) {
	return Request("getFile", b, g, &ResponseImpl{Result: &Boolean{}})
}
func (g GetFileData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"FileId": g.FileId})
}

type UnbanChatMemberData struct {
	ChatId       int  `json:"chat_id"`
	UserId       int  `json:"user_id"`
	OnlyIfBanned bool `json:"only_if_banned"`
}

func (u UnbanChatMemberData) Send(b Bot) (response Response, err error) {
	return Request("unbanChatMember", b, u, &ResponseImpl{Result: &Boolean{}})
}
func (u UnbanChatMemberData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": u.ChatId, "UserId": u.UserId})
}

type SetChatAdministratorCustomTitleData struct {
	ChatId      int    `json:"chat_id"`
	UserId      int    `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

func (s SetChatAdministratorCustomTitleData) Send(b Bot) (response Response, err error) {
	return Request("setChatAdministratorCustomTitle", b, s, &ResponseImpl{Result: &Boolean{}})
}
func (s SetChatAdministratorCustomTitleData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "UserId": s.UserId,
		"CustomTitle": s.CustomTitle})
}

type SetChatPermissionsData struct {
	ChatId      int             `json:"chat_id"`
	Permissions ChatPermissions `json:"permissions"`
}

func (s SetChatPermissionsData) Send(b Bot) (response Response, err error) {
	return Request("setChatPermissions", b, s, &ResponseImpl{Result: &Boolean{}})
}
func (s SetChatPermissionsData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId})
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

func (u UserProfilePhotosData) Send(b Bot) (response Response, err error) {

	return Request("getUserProfilePhotos", b, u, &ResponseImpl{Result: &UserProfilePhotos{}})
}
func (u UserProfilePhotosData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"UserId": u.UserId})
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

func (ban BanChatMemberData) Send(b Bot) (response Response, err error) {
	return Request("banChatMember", b, ban, &ResponseImpl{Result: &Boolean{}})
}
func (ban BanChatMemberData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"UserId": ban.UserId, "ChatId": ban.ChatId})
}

type RestrictChatMemberData struct {
	ChatId      int             `json:"chat_id"`
	UserId      int             `json:"user_id"`
	Permissions ChatPermissions `json:"permissions"`
	UntilDate   int             `json:"until_date"`
}

func (r RestrictChatMemberData) Send(b Bot) (response Response, err error) {
	return Request("restrictChatMember", b, r, &ResponseImpl{Result: &Boolean{}})
}
func (r RestrictChatMemberData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"UserId": r.UserId, "ChatId": r.ChatId})
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

func (p PromoteChatMemberData) Send(b Bot) (response Response, err error) {
	return Request("promoteChatMember", b, p, &ResponseImpl{Result: &Boolean{}})
}
func (p PromoteChatMemberData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"UserId": p.UserId, "ChatId": p.ChatId})
}

type CreateChatInviteLinkData struct {
	ChatId      int `json:"chat_id"`
	ExpireDate  int `json:"expire_date"`
	MemberLimit int `json:"member_limit"`
}

func (c CreateChatInviteLinkData) Send(b Bot) (response Response, err error) {
	return Request("createChatInviteLink", b, c, &ResponseImpl{Result: &ChatInviteLink{}})
}
func (c CreateChatInviteLinkData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": c.ChatId})
}

type EditChatInviteLinkData struct {
	ChatId      int    `json:"chat_id"`
	InviteLink  string `json:"invite_link"`
	ExpireDate  int    `json:"expire_date"`
	MemberLimit int    `json:"member_limit"`
}

func (e EditChatInviteLinkData) Send(b Bot) (response Response, err error) {
	return Request("editChatInviteLink", b, e, &ResponseImpl{Result: &ChatInviteLink{}})
}
func (e EditChatInviteLinkData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": e.ChatId, "InviteLink": e.InviteLink})
}

type PinChatMessageData struct {
	ChatId              int  `json:"chat_id"`
	MessageId           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification"`
}

func (p PinChatMessageData) Send(b Bot) (response Response, err error) {
	return Request("pinChatMessage", b, p, &ResponseImpl{Result: &Boolean{}})
}
func (p PinChatMessageData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": p.ChatId, "MessageId": p.MessageId})
}

type UnpinChatMessageData struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}

func (u UnpinChatMessageData) Send(b Bot) (response Response, err error) {
	return Request("unpinChatMessage", b, u, &ResponseImpl{Result: &Boolean{}})
}
func (u UnpinChatMessageData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": u.ChatId, "MessageId": u.MessageId})
}

type AnswerCallbackQueryData struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text"`
	ShowAlert       bool   `json:"show_alert"`
	Url             string `json:"url"`
	CacheTime       string `json:"cache_time"`
}

func (a AnswerCallbackQueryData) Send(b Bot) (response Response, err error) {
	return Request("answerCallbackQuery", b, a, &ResponseImpl{Result: &Boolean{}})
}
func (a AnswerCallbackQueryData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"CallbackQueryId": a.CallbackQueryId})
}

type SetMyCommandsData struct {
	Commands []BotCommand `json:"commands"`
	// Scope describing scope of users for which the commands are relevant. Defaults to "default".
	Scope        BotCommandScope `json:"scope"`
	LanguageCode string          `json:"language_code"`
}

func (s SetMyCommandsData) Send(b Bot) (response Response, err error) {
	return Request("setMyCommands", b, s, &ResponseImpl{Result: &Boolean{}})
}
func (s SetMyCommandsData) check() error {
	types := map[string]int{"default": 1, "chat_member": 1, "all_private_chats": 1,
		"all_group_chats": 1, "all_chat_administrators": 1, "chat": 1, "chat_administrators": 1}
	if _, ok := types[s.Scope.Type]; ok == false {
		s.Scope.Type = "default"
	}
	return globalEmptyFieldChecker(map[string]interface{}{"Commands": s.Commands})
}

type DeleteMyCommandsData struct {
	// Scope describing scope of users for which the commands are relevant. Defaults to "default".
	Scope        BotCommandScope `json:"scope"`
	LanguageCode string          `json:"language_code"`
}

func (d DeleteMyCommandsData) Send(b Bot) (response Response, err error) {
	return Request("deleteMyCommands", b, d, &ResponseImpl{Result: &Boolean{}})
}
func (d DeleteMyCommandsData) check() error {
	types := map[string]int{"default": 1, "chat_member": 1, "all_private_chats": 1,
		"all_group_chats": 1, "all_chat_administrators": 1, "chat": 1, "chat_administrators": 1}
	if _, ok := types[d.Scope.Type]; ok == false {
		d.Scope.Type = "default"
	}
	return nil
}

type GetMyCommandsData struct {
	// Scope describing scope of users for which the commands are relevant. Defaults to "default".
	Scope        BotCommandScope `json:"scope"`
	LanguageCode string          `json:"language_code"`
}

func (g GetMyCommandsData) Send(b Bot) (response Response, err error) {
	return Request("getMyCommands", b, g, &ResponseImpl{Result: &Boolean{}})
}
func (g GetMyCommandsData) check() error {
	return nil
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

func (e EditMessageTextData) Send(b Bot) (response Response, err error) {
	return Request("editMessageText", b, e, &ResponseImpl{})
}
func (e EditMessageTextData) check() error {
	if e.InlineMessageId == "" {
		if e.ChatId == 0 || e.MessageId == 0 {
			return errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	return globalEmptyFieldChecker(map[string]interface{}{"Text": e.Text})
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

func (e EditMessageCaptionData) Send(b Bot) (response Response, err error) {
	return Request("editMessageCaption", b, e, &ResponseImpl{})
}
func (e EditMessageCaptionData) check() error {
	if e.InlineMessageId == "" {
		if e.ChatId == 0 || e.MessageId == 0 {
			return errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	return nil
}

type EditMessageReplyMarkupData struct {
	ChatId          int    `json:"chat_id"`
	MessageId       int    `json:"message_id"`
	InlineMessageId string `json:"inline_message_id"`
	InlineKeyboard
}

func (e EditMessageReplyMarkupData) Send(b Bot) (response Response, err error) {
	return Request("editMessageReplyMarkup", b, e, &ResponseImpl{})
}
func (e EditMessageReplyMarkupData) check() error {
	if e.InlineMessageId == "" {
		if e.ChatId == 0 || e.MessageId == 0 {
			return errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	return nil
}

type StopPollData struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
	InlineKeyboard
}

func (s StopPollData) Send(b Bot) (response Response, err error) {
	return Request("stopPoll", b, s, &ResponseImpl{Result: &Poll{}})
}
func (s StopPollData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "MessageId": s.MessageId})
}

type EditMessageMediaData struct {
	InlineMessageId string     `json:"inline_message_id"`
	Media           InputMedia `json:"media"`
	ChatId          int        `json:"chat_id"`
	MessageId       int        `json:"message_id"`
	Files           []*os.File
	InlineKeyboard
}

func (e EditMessageMediaData) Send(b Bot) (response Response, err error) {
	e.Files = append(e.Files, e.Media.returnFile())
	return Request("editMessageMedia", b, e, &ResponseImpl{})
}
func (e EditMessageMediaData) check() error {
	if e.InlineMessageId == "" {
		if e.ChatId == 0 || e.MessageId == 0 {
			return errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	if e.Media == nil {
		return errors.New("media is nil. pass media a struct of type " +
			"InputMediaPhoto, InputMediaVideo, InputMediaDocument or InputMediaAudio")
	}
	return nil
}

type SetWebhookData struct {
	Url                string   `json:"url"`
	Certificate        *os.File `json:"certificate"`
	IpAddress          string   `json:"ip_address"`
	MaxConnections     int      `json:"max_connections"`
	AllowedUpdates     []string `json:"allowed_updates"`
	DropPendingUpdates bool     `json:"drop_pending_updates"`
}

func (s SetWebhookData) Send(b Bot) (response Response, err error) {
	return Request("setWebhook", b, s, &ResponseImpl{Result: &Boolean{}})
}
func (s SetWebhookData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Url": s.Url})
}

type SendStickerData struct {
	ChatId                   int         `json:"chat_id"`
	Sticker                  interface{} `json:"sticker"`
	DisableNotification      bool        `json:"disable_notification"`
	ReplyToMessageId         int         `json:"reply_To_Message_Id"`
	AllowSendingWithoutReply bool        `json:"allow_sending_without_reply"`
	Keyboard
}

func (s SendStickerData) Send(b Bot) (response Response, err error) {
	return Request("sendSticker", b, s, &ResponseImpl{Result: &Message{}})
}
func (s SendStickerData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"ChatId": s.ChatId, "Sticker": s.Sticker})
}

type DeleteStickerFromSetData struct {
	Sticker string `json:"sticker"`
}

func (d DeleteStickerFromSetData) Send(b Bot) (response Response, err error) {
	return Request("deleteStickerFromSet", b, d, &ResponseImpl{Result: &Boolean{}})
}
func (d DeleteStickerFromSetData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Sticker": d.Sticker})
}

type SetStickerPositionInSetData struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

func (s SetStickerPositionInSetData) Send(b Bot) (response Response, err error) {
	return Request("setStickerPositionInSet", b, s, &ResponseImpl{Result: &Boolean{}})
}
func (s SetStickerPositionInSetData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Sticker": s.Sticker, "Position": s.Position})
}

type UploadStickerFileData struct {
	UserId     int      `json:"user_id"`
	PngSticker *os.File `json:"png_sticker"`
}

func (u UploadStickerFileData) Send(b Bot) (response Response, err error) {
	return Request("uploadStickerFile", b, u, &ResponseImpl{Result: &Boolean{}})
}
func (u UploadStickerFileData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"UserId": u.UserId, "PngSticker": u.PngSticker})
}

type GetStickerSetData struct {
	Name string `json:"name"`
}

func (g GetStickerSetData) Send(b Bot) (response Response, err error) {
	return Request("getStickerSet", b, g, &ResponseImpl{Result: &Boolean{}})
}
func (g GetStickerSetData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"Name": g.Name})
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

func (c CreateNewStickerSetData) Send(b Bot) (response Response, err error) {
	return Request("createNewStickerSet", b, c, &ResponseImpl{Result: &Boolean{}})
}
func (c CreateNewStickerSetData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"UserId": c.UserId, "Name": c.Name,
		"Title": c.Title, "Emojis": c.Emojis})
}

type AddStickerToSetData struct {
	UserId       int          `json:"user_id"`
	Name         string       `json:"name"`
	Emojis       string       `json:"emojis"`
	PngSticker   interface{}  `json:"png_sticker"`
	TgsSticker   *os.File     `json:"tgs_sticker"`
	MaskPosition MaskPosition `json:"mask_position"`
}

func (a AddStickerToSetData) Send(b Bot) (response Response, err error) {
	return Request("addStickerToSet", b, a, &ResponseImpl{Result: &Boolean{}})
}
func (a AddStickerToSetData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"UserId": a.UserId, "Name": a.Name, "Emojis": a.Emojis})
}

type SetStickerSetThumbData struct {
	UserId int         `json:"user_id"`
	Name   string      `json:"name"`
	Thumb  interface{} `json:"thumb"`
}

func (s SetStickerSetThumbData) Send(b Bot) (response Response, err error) {
	return Request("setStickerSetThumb", b, s, &ResponseImpl{Result: &Boolean{}})
}
func (s SetStickerSetThumbData) check() error {
	return globalEmptyFieldChecker(map[string]interface{}{"UserId": s.UserId, "Name": s.Name})
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

func (a AnswerInlineQueryData) Send(b Bot) (response Response, err error) {
	return Request("answerInlineQuery", b, a, &ResponseImpl{Result: &Boolean{}})
}
func (a AnswerInlineQueryData) check() error {
	if len(a.Results) == 0 {
		return errors.New("results slice is empty. pass QueryAnswer structs such as " +
			"InlineQueryResultArticle, InlineQueryResultPhoto and etc")
	}
	for _, j := range a.Results {
		if e := j.checkQueryAnswer(); e != nil {
			return e
		}
	}
	return globalEmptyFieldChecker(map[string]interface{}{"InlineQueryId": a.InlineQueryId})
}
