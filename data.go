package gogram

import (
	"errors"
	"os"
)

type Method interface {
	// Check is used in Request
	Check() error
	// Send Sends requests to telegram server using Request
	Send(b Bot) (Response, error)
}

// TextData sends text messages. On success, the sent Message is returned.
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

func (t TextData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Text": t.Text, "ChatId": t.ChatId, "ParseMode": t.ParseMode})
}

// PhotoData sends photos. On success, the sent Message is returned.
type PhotoData struct {
	// photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new
	// video using os.Open(<file_name>). The photo must be at most 10 MB in size.
	// The photo's width and height must not exceed 10000 in total.
	// Width and height ratio must be at most 20.
	Photo                    any             `json:"photo"`
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

func (p PhotoData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Photo": p.Photo, "ChatId": p.ChatId, "ParseMode": p.ParseMode})
}

// VideoData sends video files, Telegram clients support mp4 videos (other formats may be sent as Document).
// On success, the sent Message is returned.
// Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type VideoData struct {
	ChatId int `json:"chat_id"`
	// video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a video from the Internet, or
	// upload a new video using os.Open(<file_name>).
	Video                    any             `json:"video"`
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

func (v VideoData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Video": v.Video, "ChatId": v.ChatId,
		"ParseMode": v.ParseMode})
}

// AudioData sends audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned.
// Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
type AudioData struct {
	ChatId int `json:"chat_id"`
	// audio file to send. Pass a file_id as string to send an audio file that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a string for Telegram to get an audio file from the Internet,
	// or upload a new video using os.Open(<file_name>).
	Audio                    any             `json:"audio"`
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
func (a AudioData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Audio": a.Audio, "ChatId": a.ChatId, "ParseMode": a.ParseMode})
}

// DocumentData sends general files. On success, the sent Message is returned.
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type DocumentData struct {
	ChatId int `json:"chat_id"`
	// file to send. Pass a file_id as string to send an audio file that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a string for Telegram to get a file from the Internet,
	// or upload a new video using os.Open(<file_name>).
	Document                    any             `json:"document"`
	Caption                     string          `json:"caption"`
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
func (d DocumentData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Document": d.Document, "ChatId": d.ChatId,
		"ParseMode": d.ParseMode})
}

// VoiceData sends audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with OPUS
// (other formats may be sent as Audio or Document). On success, the sent Message is returned.
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type VoiceData struct {
	ChatId int `json:"chat_id"`
	// audio file to send. Pass a file_id as string to send an audio file that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a string for Telegram to get an audio file from the Internet,
	// or upload a new video using os.Open(<file_name>).
	Voice                    any             `json:"voice"`
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
func (v VoiceData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Voice": v.Voice, "ChatId": v.ChatId,
		"ParseMode": v.ParseMode})
}

// AnimationData sends animation files (GIF or H.264/MPEG-4 AVC video without sound).
// On success, the sent Message is returned.
// Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
type AnimationData struct {
	ChatId                   int             `json:"chat_id"`
	Animation                any             `json:"animation"`
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
func (a AnimationData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Animation": a.Animation, "ChatId": a.ChatId,
		"ParseMode": a.ParseMode})
}

// PollData sends a native poll. On success, the sent Message is returned.
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
func (p PollData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Question": p.Question, "ChatId": p.ChatId,
		"Options": p.Options})
}

// DiceData sends an animated emoji that will display a random value.
// On success, the sent Message is returned.
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

func (d DiceData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Emoji": d.Emoji, "ChatId": d.ChatId})
}

// VideoNoteData sends video messages.
// As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long.
// On success, the sent Message is returned.
type VideoNoteData struct {
	ChatId                   int  `json:"chat_id"`
	VideoNote                any  `json:"videoNote"`
	Duration                 int  `json:"duration"`
	Length                   int  `json:"length"`
	DisableNotification      bool `json:"disable_notification"`
	ReplyToMessageId         int  `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`
	Keyboard
}

func (v VideoNoteData) Send(b Bot) (response Response, err error) {
	return Request("sendVideoNote", b, v, &ResponseImpl{Result: &Message{}})
}
func (v VideoNoteData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"VideoNote": v.VideoNote, "ChatId": v.ChatId})
}

// LocationData sends point on the map.
// On success, the sent Message is returned.
type LocationData struct {
	ChatId int `json:"chat_id"`
	Location
	DisableNotification      bool `json:"disable_notification"`
	ReplyToMessageId         int  `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`
	Keyboard
}

func (l LocationData) Send(b Bot) (response Response, err error) {
	return Request("sendLocation", b, l, &ResponseImpl{Result: &Message{}})
}
func (l LocationData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": l.ChatId})
}

// ContactData sends phone contacts.
// On success, the sent Message is returned.
type ContactData struct {
	ChatId int `json:"chat_id"`
	Contact
	DisableNotification      bool `json:"disable_notification"`
	ReplyToMessageId         int  `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`
	Keyboard
}

func (c ContactData) Send(b Bot) (response Response, err error) {
	return Request("sendContact", b, c, &ResponseImpl{Result: &Message{}})
}
func (c ContactData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"PhoneNumber": c.PhoneNumber, "ChatId": c.ChatId,
		"FirstName": c.FirstName})
}

// MediaGroupData sends a group of photos, videos, documents or audios as an album.
// Documents and audio files can be only grouped in an album with messages of the same type.
// On success, an array of Messages that were sent is returned.
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
func (m MediaGroupData) Check() error {
	if len(m.Media) == 0 {
		return errors.New("media slice is empty. pass media a slice of structs of type " +
			"InputMediaPhoto, InputMediaVideo, InputMediaDocument or InputMediaAudio")
	}
	return nil
}

// ForwardMessageData forwards messages of any kind. Service messages can't be forwarded.
// On success, the sent Message is returned.
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
func (f ForwardMessageData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"FromChatId": f.FromChatId, "ChatId": f.ChatId,
		"MessageId": f.MessageId})
}

// CopyMessageData copies messages of any kind. Service messages and invoice messages can't be copied.
// The method is analogous to the method forwardMessage, but the copied message doesn't have a link to
// the original message. Returns the MessageId of the sent message on success.
type CopyMessageData struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Unique identifier for the chat where the original message was sent
	// (or channel username in the format @channelusername)
	FromChatId int `json:"from_chat_id"`
	// Message identifier in the chat specified in from_chat_id
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
func (c CopyMessageData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"FromChatId": c.FromChatId, "ChatId": c.ChatId,
		"MessageId": c.MessageId, "ParseMode": c.ParseMode})
}

// DeleteMessageData deletes a message, including service messages, with the following limitations:
//- A message can only be deleted if it was sent less than 48 hours ago.
//- A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
//- Bots can delete outgoing messages in private chats, groups, and supergroups.
//- Bots can delete incoming messages in private chats.
//- Bots granted can_post_messages permissions can delete outgoing messages in channels.
//- If the bot is an administrator of a group, it can delete any message there.
//- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
// Returns true on success.
type DeleteMessageData struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}

func (d DeleteMessageData) Send(b Bot) (response Response, err error) {
	return Request("deleteMessage", b, d, &ResponseImpl{})
}
func (d DeleteMessageData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": d.ChatId, "MessageId": d.MessageId})
}

// DeleteChatStickerSetData deletes a group sticker set from a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Use the field Chat.CanSetStickerSet optionally returned in GetChatData
// requests to check if the bot can use this method.
// Returns True on success.
type DeleteChatStickerSetData struct {
	ChatId int `json:"chat_id"`
}

func (d DeleteChatStickerSetData) Send(b Bot) (response Response, err error) {
	return Request("deleteChatStickerSet", b, d, &ResponseImpl{})
}
func (d DeleteChatStickerSetData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": d.ChatId})
}

// SetChatStickerSetData sets a new group sticker set for a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Use the field Chat.CanSetStickerSet optionally returned in GetChatData
// requests to check if the bot can use this method.
// Returns True on success.
type SetChatStickerSetData struct {
	ChatId         int    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

func (s SetChatStickerSetData) Send(b Bot) (response Response, err error) {
	return Request("setChatStickerSet", b, s, &ResponseImpl{})
}
func (s SetChatStickerSetData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": s.ChatId, "StickerSetName": s.StickerSetName})
}

// GetChatMemberData gets information about a member of a chat.
// Returns a ChatMember object on success.
type GetChatMemberData struct {
	ChatId int `json:"chat_id"`
	UserId int `json:"user_id"`
}

func (g GetChatMemberData) Send(b Bot) (response Response, err error) {
	return Request("getChatMember", b, g, &ResponseImpl{})
}

func (g GetChatMemberData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": g.ChatId, "UserId": g.UserId})
}

// GetChatMemberCountData gets the number of members in a chat. Returns Int on success.
type GetChatMemberCountData struct {
	ChatId int `json:"chat_id"`
}

func (g GetChatMemberCountData) Send(b Bot) (response Response, err error) {
	return Request("getChatMemberCount", b, g, &ResponseImpl{})
}
func (g GetChatMemberCountData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": g.ChatId})
}

// GetChatAdministratorsData gets a list of administrators in a chat.
// On success, returns an Array of ChatMember objects that contains information about
// all chat administrators except other bots.
// If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
type GetChatAdministratorsData struct {
	ChatId int `json:"chat_id"`
}

func (g GetChatAdministratorsData) Send(b Bot) (response Response, err error) {
	return Request("getChatAdministrators", b, g, &ResponseImpl{})
}
func (g GetChatAdministratorsData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": g.ChatId})
}

// GetChatData gets up-to-date information about the chat (current name of the user for one-on-one
// conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
type GetChatData struct {
	ChatId int `json:"chat_id"`
}

func (g GetChatData) Send(b Bot) (response Response, err error) {
	return Request("getChat", b, g, &ResponseImpl{Result: &Chat{}})
}
func (g GetChatData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": g.ChatId})
}

// LeaveChatData leaves a group, supergroup or channel for your bot.
// Returns True on success.
type LeaveChatData struct {
	ChatId int `json:"chat_id"`
}

func (l LeaveChatData) Send(b Bot) (response Response, err error) {
	return Request("leaveChat", b, l, &ResponseImpl{})
}
func (l LeaveChatData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": l.ChatId})
}

// UnpinAllChatMessagesData clears the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work and
// must have the 'can_pin_messages' administrator right in a supergroup or
// 'can_edit_messages' administrator right in a channel.
// Returns True on success.
type UnpinAllChatMessagesData struct {
	ChatId int `json:"chat_id"`
}

func (u UnpinAllChatMessagesData) Send(b Bot) (response Response, err error) {
	return Request("unpinAllChatMessages", b, u, &ResponseImpl{})
}
func (u UnpinAllChatMessagesData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": u.ChatId})
}

// SetChatDescriptionData changes the description of a group, a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type SetChatDescriptionData struct {
	ChatId      int    `json:"chat_id"`
	Description string `json:"description"`
}

func (s SetChatDescriptionData) Send(b Bot) (response Response, err error) {
	return Request("setChatDescription", b, s, &ResponseImpl{})
}
func (s SetChatDescriptionData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": s.ChatId, "Description": s.Description})
}

// SetChatTitleData changes the title of a chat. Titles can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type SetChatTitleData struct {
	ChatId int    `json:"chat_id"`
	Title  string `json:"title"`
}

func (s SetChatTitleData) Send(b Bot) (response Response, err error) {
	return Request("setChatTitle", b, s, &ResponseImpl{})
}
func (s SetChatTitleData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": s.ChatId, "Title": s.Title})
}

// DeleteChatPhotoData deletes a chat photo. Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type DeleteChatPhotoData struct {
	ChatId int `json:"chat_id"`
}

func (d DeleteChatPhotoData) Send(b Bot) (response Response, err error) {
	return Request("deleteChatPhoto", b, d, &ResponseImpl{})
}
func (d DeleteChatPhotoData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": d.ChatId})
}

// SetChatPhotoData sets a new profile photo for the chat. Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type SetChatPhotoData struct {
	ChatId int      `json:"chat_id"`
	Photo  *os.File `json:"photo"`
}

func (s SetChatPhotoData) Send(b Bot) (response Response, err error) {
	return Request("setChatPhoto", b, s, &ResponseImpl{})
}
func (s SetChatPhotoData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": s.ChatId, "Photo": s.Photo})
}

// RevokeChatInviteLinkData revokes an invitation link created by the bot.
// If the primary link is revoked, a new link is automatically generated.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns the revoked invite link as ChatInviteLink object.
type RevokeChatInviteLinkData struct {
	ChatId     int    `json:"chat_id"`
	InviteLink string `json:"invite_link"`
}

func (r RevokeChatInviteLinkData) Send(b Bot) (response Response, err error) {
	return Request("revokeChatInviteLink", b, r, &ResponseImpl{Result: &ChatInviteLink{}})
}
func (r RevokeChatInviteLinkData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": r.ChatId, "InviteLink": r.InviteLink})
}

// ExportChatInviteLinkData generates a new primary invite link for a chat; any previously generated primary
// link is revoked. The bot must be an administrator in the chat for this to work and must have the
// appropriate administrator rights.
// Returns the new invite link as String on success.
type ExportChatInviteLinkData struct {
	ChatId int `json:"chat_id"`
}

func (e ExportChatInviteLinkData) Send(b Bot) (response Response, err error) {
	return Request("exportChatInviteLink", b, e, &ResponseImpl{})
}
func (e ExportChatInviteLinkData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": e.ChatId})
}

// SendChatActionData tells the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot, Telegram
// clients clear its typing status). for more info visit https://core.telegram.org/bots/api#sendchataction
// Returns True on success.
type SendChatActionData struct {
	ChatId int    `json:"chat_id"`
	Action string `json:"action"`
}

func (s SendChatActionData) Send(b Bot) (response Response, err error) {
	return Request("sendChatAction", b, s, &ResponseImpl{})
}
func (s SendChatActionData) Check() error {
	var actions = map[string]bool{"typing": true, "upload_photo": true, "record_video": true, "upload_video": true,
		"general": true, "upload_document": true, "upload_voice": true, "record_voice": true, "find_location": true,
		"record_video_note": true, "upload_video_note": true}
	if _, ok := actions[s.Action]; ok == false {
		return errors.New(s.Action + " is an unknown action, read the document")
	}
	return globalEmptyFieldChecker(map[string]any{"ChatId": s.ChatId})
}

// GetFileData gets basic info about a file and prepare it for downloading.
// For the moment, bots can download files of up to 20MB in size.
// On success, a File object is returned.
// The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>,
// where <file_path> is taken from the response.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile again.
type GetFileData struct {
	FileId string `json:"file_id"`
}

func (g GetFileData) Send(b Bot) (response Response, err error) {
	return Request("getFile", b, g, &ResponseImpl{Result: &File{}})
}
func (g GetFileData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"FileId": g.FileId})
}

// UnbanChatMemberData unbans a previously banned user in a supergroup or channel.
// The user will not return to the group or channel automatically, but will be able to join via link, etc.
// The bot must be an administrator for this to work. By default, this method guarantees that after the
// call the user is not a member of the chat, but will be able to join it. So if the user is a member of
// the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned.
// Returns True on success.
type UnbanChatMemberData struct {
	ChatId       int  `json:"chat_id"`
	UserId       int  `json:"user_id"`
	OnlyIfBanned bool `json:"only_if_banned"`
}

func (u UnbanChatMemberData) Send(b Bot) (response Response, err error) {
	return Request("unbanChatMember", b, u, &ResponseImpl{})
}
func (u UnbanChatMemberData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": u.ChatId, "UserId": u.UserId})
}

// SetChatAdministratorCustomTitleData sets a custom title for an administrator in a supergroup promoted by the bot.
// Returns True on success.
type SetChatAdministratorCustomTitleData struct {
	ChatId      int    `json:"chat_id"`
	UserId      int    `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

func (s SetChatAdministratorCustomTitleData) Send(b Bot) (response Response, err error) {
	return Request("setChatAdministratorCustomTitle", b, s, &ResponseImpl{})
}

func (s SetChatAdministratorCustomTitleData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": s.ChatId, "UserId": s.UserId,
		"CustomTitle": s.CustomTitle})
}

// SetChatPermissionsData sets default chat permissions for all members.
// The bot must be an administrator in the group or a supergroup for this to work and must
// have the can_restrict_members administrator rights. Returns True on success.
type SetChatPermissionsData struct {
	ChatId      int             `json:"chat_id"`
	Permissions ChatPermissions `json:"permissions"`
}

func (s SetChatPermissionsData) Send(b Bot) (response Response, err error) {
	return Request("setChatPermissions", b, s, &ResponseImpl{})
}
func (s SetChatPermissionsData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": s.ChatId})
}

// GetUserProfilePhotosData gets a list of profile pictures for a user. Returns a UserProfilePhotos object.
type GetUserProfilePhotosData struct {
	UserId int `json:"user_id"`
	// Sequential number of the first photo to be returned.
	// By default, all photos are returned.
	Offset int `json:"offset"`
	// Limits the number of photos to be retrieved.
	// Values between 1-100 are accepted. Defaults to 100.
	Limit int `json:"limit"`
}

func (u GetUserProfilePhotosData) Send(b Bot) (response Response, err error) {
	return Request("getUserProfilePhotos", b, u, &ResponseImpl{Result: &UserProfilePhotos{}})
}
func (u GetUserProfilePhotosData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"UserId": u.UserId})
}

// BanChatMemberData bans a user in a group, a supergroup or a channel.
// In the case of supergroups and channels, the user will not be able to return to the chat on
// their own using invite links, etc., unless unbanned first.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
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
	return Request("banChatMember", b, ban, &ResponseImpl{})
}
func (ban BanChatMemberData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"UserId": ban.UserId, "ChatId": ban.ChatId})
}

// RestrictChatMemberData restricts a user in a supergroup. The bot must be an administrator in the
// supergroup for this to work and must have the appropriate administrator rights.
// Pass True for all permissions to lift restrictions from a user. Returns True on success.
type RestrictChatMemberData struct {
	ChatId      int             `json:"chat_id"`
	UserId      int             `json:"user_id"`
	Permissions ChatPermissions `json:"permissions"`
	UntilDate   int             `json:"until_date"`
}

func (r RestrictChatMemberData) Send(b Bot) (response Response, err error) {
	return Request("restrictChatMember", b, r, &ResponseImpl{})
}
func (r RestrictChatMemberData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"UserId": r.UserId, "ChatId": r.ChatId})
}

// PromoteChatMemberData promotes or demotes a user in a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Pass False for all boolean parameters to demote a user. Returns True on success.
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
	return Request("promoteChatMember", b, p, &ResponseImpl{})
}
func (p PromoteChatMemberData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"UserId": p.UserId, "ChatId": p.ChatId})
}

// CreateChatInviteLinkData creates an additional invite link for a chat.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
type CreateChatInviteLinkData struct {
	ChatId      int `json:"chat_id"`
	ExpireDate  int `json:"expire_date"`
	MemberLimit int `json:"member_limit"`
}

func (c CreateChatInviteLinkData) Send(b Bot) (response Response, err error) {
	return Request("createChatInviteLink", b, c, &ResponseImpl{Result: &ChatInviteLink{}})
}
func (c CreateChatInviteLinkData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": c.ChatId})
}

// EditChatInviteLinkData edits a non-primary invite link created by the bot.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns the edited invite link as a ChatInviteLink object.
type EditChatInviteLinkData struct {
	ChatId      int    `json:"chat_id"`
	InviteLink  string `json:"invite_link"`
	ExpireDate  int    `json:"expire_date"`
	MemberLimit int    `json:"member_limit"`
}

func (e EditChatInviteLinkData) Send(b Bot) (response Response, err error) {
	return Request("editChatInviteLink", b, e, &ResponseImpl{Result: &ChatInviteLink{}})
}
func (e EditChatInviteLinkData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": e.ChatId, "InviteLink": e.InviteLink})
}

// PinChatMessageData adds a message to the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for
// this to work and must have the 'can_pin_messages' administrator right in a supergroup or
// 'can_edit_messages' administrator right in a channel. Returns True on success.
type PinChatMessageData struct {
	ChatId              int  `json:"chat_id"`
	MessageId           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification"`
}

func (p PinChatMessageData) Send(b Bot) (response Response, err error) {
	return Request("pinChatMessage", b, p, &ResponseImpl{})
}
func (p PinChatMessageData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": p.ChatId, "MessageId": p.MessageId})
}

// UnpinChatMessageData removes a message from the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this
// to work and must have the 'can_pin_messages' administrator right in a supergroup or
// 'can_edit_messages' administrator right in a channel. Returns True on success.
type UnpinChatMessageData struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}

func (u UnpinChatMessageData) Send(b Bot) (response Response, err error) {
	return Request("unpinChatMessage", b, u, &ResponseImpl{})
}
func (u UnpinChatMessageData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": u.ChatId, "MessageId": u.MessageId})
}

// AnswerCallbackQueryData sends answers to callback queries sent from inline keyboards.
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
// On success, True is returned. more info in https://core.telegram.org/bots/api#answercallbackquery
type AnswerCallbackQueryData struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text"`
	ShowAlert       bool   `json:"show_alert"`
	Url             string `json:"url"`
	CacheTime       string `json:"cache_time"`
}

func (a AnswerCallbackQueryData) Send(b Bot) (response Response, err error) {
	return Request("answerCallbackQuery", b, a, &ResponseImpl{})
}
func (a AnswerCallbackQueryData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"CallbackQueryId": a.CallbackQueryId})
}

// SetMyCommandsData changes the list of the bot's commands.
// See https://core.telegram.org/bots#commands for more details about bot commands.
// Returns True on success.
type SetMyCommandsData struct {
	Commands []BotCommand `json:"commands"`
	// Scope describing scope of users for which the commands are relevant. Defaults to "default".
	Scope        BotCommandScope `json:"scope"`
	LanguageCode string          `json:"language_code"`
}

func (s SetMyCommandsData) Send(b Bot) (response Response, err error) {
	return Request("setMyCommands", b, s, &ResponseImpl{})
}
func (s SetMyCommandsData) Check() error {
	types := map[string]int{"default": 1, "chat_member": 1, "all_private_chats": 1,
		"all_group_chats": 1, "all_chat_administrators": 1, "chat": 1, "chat_administrators": 1}
	if _, ok := types[s.Scope.Type]; ok == false {
		s.Scope.Type = "default"
	}
	return globalEmptyFieldChecker(map[string]any{"Commands": s.Commands})
}

// DeleteMyCommandsData deletes the list of the bot's commands for the given scope and user language.
// After deletion, higher level commands will be shown to affected users.
// Returns True on success.
type DeleteMyCommandsData struct {
	// Scope describing scope of users for which the commands are relevant. Defaults to "default".
	Scope        BotCommandScope `json:"scope"`
	LanguageCode string          `json:"language_code"`
}

func (d DeleteMyCommandsData) Send(b Bot) (response Response, err error) {
	return Request("deleteMyCommands", b, d, &ResponseImpl{})
}
func (d DeleteMyCommandsData) Check() error {
	types := map[string]int{"default": 1, "chat_member": 1, "all_private_chats": 1,
		"all_group_chats": 1, "all_chat_administrators": 1, "chat": 1, "chat_administrators": 1}
	if _, ok := types[d.Scope.Type]; ok == false {
		d.Scope.Type = "default"
	}
	return nil
}

// GetMyCommandsData gets the current list of the bot's commands for the given scope and user language.
// Returns Array of BotCommand on success. If commands aren't set, an empty list is returned.
type GetMyCommandsData struct {
	// Scope describing scope of users for which the commands are relevant. Defaults to "default".
	Scope        BotCommandScope `json:"scope"`
	LanguageCode string          `json:"language_code"`
}

func (g GetMyCommandsData) Send(b Bot) (response Response, err error) {
	return Request("getMyCommands", b, g, &ResponseImpl{Result: &BotCommand{}})
}
func (g GetMyCommandsData) Check() error {
	return nil
}

// EditMessageTextData edit text and Game messages.
// On success, if the edited message is not an inline message, the edited Message is returned,
// otherwise True is returned.
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
func (e EditMessageTextData) Check() error {
	if e.InlineMessageId == "" {
		if e.ChatId == 0 || e.MessageId == 0 {
			return errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	return globalEmptyFieldChecker(map[string]any{"Text": e.Text, "ParseMode": e.ParseMode})
}

// EditMessageCaptionData edit captions of messages.
// On success, if the edited message is not an inline message, the edited Message is returned,
// otherwise True is returned.
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
func (e EditMessageCaptionData) Check() error {
	if e.InlineMessageId == "" {
		if e.ChatId == 0 || e.MessageId == 0 {
			return errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	return globalEmptyFieldChecker(map[string]any{"ParseMode": e.ParseMode})
}

// EditMessageReplyMarkupData edit only the reply markup of messages.
// On success, if the edited message is not an inline message, the edited Message is returned,
// otherwise True is returned.
type EditMessageReplyMarkupData struct {
	ChatId          int    `json:"chat_id"`
	MessageId       int    `json:"message_id"`
	InlineMessageId string `json:"inline_message_id"`
	InlineKeyboard
}

func (e EditMessageReplyMarkupData) Send(b Bot) (response Response, err error) {
	return Request("editMessageReplyMarkup", b, e, &ResponseImpl{})
}
func (e EditMessageReplyMarkupData) Check() error {
	if e.InlineMessageId == "" {
		if e.ChatId == 0 || e.MessageId == 0 {
			return errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	return nil
}

// StopPollData stop a poll which was sent by the bot. On success, the stopped Poll is returned.
type StopPollData struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
	InlineKeyboard
}

func (s StopPollData) Send(b Bot) (response Response, err error) {
	return Request("stopPoll", b, s, &ResponseImpl{Result: &Poll{}})
}
func (s StopPollData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": s.ChatId, "MessageId": s.MessageId})
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
func (e EditMessageMediaData) Check() error {
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
	return Request("setWebhook", b, s, &ResponseImpl{})
}
func (s SetWebhookData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Url": s.Url})
}

type SendStickerData struct {
	ChatId                   int `json:"chat_id"`
	Sticker                  `json:"sticker"`
	DisableNotification      bool `json:"disable_notification"`
	ReplyToMessageId         int  `json:"reply_To_Message_Id"`
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`
	Keyboard
}

func (s SendStickerData) Send(b Bot) (response Response, err error) {
	return Request("sendSticker", b, s, &ResponseImpl{Result: &Message{}})
}
func (s SendStickerData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": s.ChatId, "Sticker": s.Sticker})
}

type DeleteStickerFromSetData struct {
	Sticker string `json:"sticker"`
}

func (d DeleteStickerFromSetData) Send(b Bot) (response Response, err error) {
	return Request("deleteStickerFromSet", b, d, &ResponseImpl{})
}
func (d DeleteStickerFromSetData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Sticker": d.Sticker})
}

type SetStickerPositionInSetData struct {
	Sticker  string `json:"sticker"`
	Position int    `json:"position"`
}

func (s SetStickerPositionInSetData) Send(b Bot) (response Response, err error) {
	return Request("setStickerPositionInSet", b, s, &ResponseImpl{})
}
func (s SetStickerPositionInSetData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Sticker": s.Sticker, "Position": s.Position})
}

type UploadStickerFileData struct {
	UserId     int      `json:"user_id"`
	PngSticker *os.File `json:"png_sticker"`
}

func (u UploadStickerFileData) Send(b Bot) (response Response, err error) {
	return Request("uploadStickerFile", b, u, &ResponseImpl{})
}
func (u UploadStickerFileData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"UserId": u.UserId, "PngSticker": u.PngSticker})
}

type GetStickerSetData struct {
	Name string `json:"name"`
}

func (g GetStickerSetData) Send(b Bot) (response Response, err error) {
	return Request("getStickerSet", b, g, &ResponseImpl{})
}
func (g GetStickerSetData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"Name": g.Name})
}

type CreateNewStickerSetData struct {
	UserId        int          `json:"user_id"`
	Name          string       `json:"name"`
	Title         string       `json:"title"`
	Emojis        string       `json:"emojis"`
	PngSticker    any          `json:"png_sticker"`
	TgsSticker    *os.File     `json:"tgs_sticker"`
	ContainsMasks bool         `json:"contains_masks"`
	MaskPosition  MaskPosition `json:"mask_position"`
}

func (c CreateNewStickerSetData) Send(b Bot) (response Response, err error) {
	return Request("createNewStickerSet", b, c, &ResponseImpl{})
}
func (c CreateNewStickerSetData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"UserId": c.UserId, "Name": c.Name,
		"Title": c.Title, "Emojis": c.Emojis})
}

type AddStickerToSetData struct {
	UserId       int          `json:"user_id"`
	Name         string       `json:"name"`
	Emojis       string       `json:"emojis"`
	PngSticker   any          `json:"png_sticker"`
	TgsSticker   *os.File     `json:"tgs_sticker"`
	MaskPosition MaskPosition `json:"mask_position"`
}

func (a AddStickerToSetData) Send(b Bot) (response Response, err error) {
	return Request("addStickerToSet", b, a, &ResponseImpl{})
}
func (a AddStickerToSetData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"UserId": a.UserId, "Name": a.Name, "Emojis": a.Emojis})
}

type SetStickerSetThumbData struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Thumb  any    `json:"thumb"`
}

func (s SetStickerSetThumbData) Send(b Bot) (response Response, err error) {
	return Request("setStickerSetThumb", b, s, &ResponseImpl{})
}
func (s SetStickerSetThumbData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"UserId": s.UserId, "Name": s.Name})
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
	return Request("answerInlineQuery", b, a, &ResponseImpl{})
}
func (a AnswerInlineQueryData) Check() error {
	if len(a.Results) == 0 {
		return errors.New("results slice is empty. pass QueryAnswer structs such as " +
			"InlineQueryResultArticle, InlineQueryResultPhoto and etc")
	}
	for _, j := range a.Results {
		if e := j.checkQueryAnswer(); e != nil {
			return e
		}
	}
	return globalEmptyFieldChecker(map[string]any{"InlineQueryId": a.InlineQueryId})
}

type SendGameData struct {
	ChatId                   int    `json:"chat_id"`
	GameShortName            string `json:"game_short_name"`
	DisableNotification      bool   `json:"disable_notification"`
	ProtectContent           bool   `json:"protect_content"`
	ReplyToMessageId         int    `json:"reply_to_message_id"`
	AllowSendingWithoutReply bool   `json:"allow_sending_without_reply"`
	InlineKeyboard
}

func (s SendGameData) Send(b Bot) (response Response, err error) {
	return Request("sendGame", b, s, &ResponseImpl{Result: &Message{}})
}
func (s SendGameData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": s.ChatId, "GameShortName": s.GameShortName})
}

// SetGameScoreData sets the score of the specified user in a game message.
// On success, if the message is not an inline message, the Message is returned,
// otherwise True is returned. Returns an error, if the new score is not greater
// than the user's current score in the chat and force is False.
type SetGameScoreData struct {
	UserId             int    `json:"user_id"`
	Score              int    `json:"score"`
	Force              bool   `json:"force"`
	DisableEditMessage bool   `json:"disable_edit_message"`
	ChatId             int    `json:"chat_id"`
	MessageId          int    `json:"message_id"`
	InlineMessageId    string `json:"inline_message_id"`
}

func (s SetGameScoreData) Send(b Bot) (response Response, err error) {
	return Request("setGameScore", b, s, &ResponseImpl{})
}

func (s SetGameScoreData) Check() error {
	if s.InlineMessageId == "" {
		if s.ChatId == 0 || s.MessageId == 0 {
			return errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	return globalEmptyFieldChecker(map[string]any{"UserId": s.UserId, "Score": s.Score})
}

// GetGameHighScoresData Use this method to get data for high score tables.
// Will return the score of the specified user and several of their neighbors in a game.
// On success, returns an Array of GameHighScore objects.
// This method will currently return scores for the target user, plus two of their closest
// neighbors on each side. Will also return the top three users if the user and his neighbors are not among them.
// Please note that this behavior is subject to change.
type GetGameHighScoresData struct {
	UserId          int    `json:"user_id"`
	ChatId          int    `json:"chat_id"`
	MessageId       int    `json:"message_id"`
	InlineMessageId string `json:"inline_message_id"`
}

func (g GetGameHighScoresData) Send(b Bot) (response Response, err error) {
	return Request("getGameHighScores", b, g, &ResponseImpl{})
}

func (g GetGameHighScoresData) Check() error {
	if g.InlineMessageId == "" {
		if g.ChatId == 0 || g.MessageId == 0 {
			return errors.New("you need to set both MessageId and " +
				"ChatId, otherwise set InlineMessageId")
		}
	}
	return globalEmptyFieldChecker(map[string]any{"UserId": g.UserId})
}

type SendInvoiceData struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId int `json:"chat_id"`
	// Product name, 1-32 characters
	Title string `json:"title"`
	// Product description, 1-255 characters
	Description string `json:"description"`
	// bot-defined invoice payload, 1-128 bytes.
	// This will not be displayed to the user, use for your internal processes.
	Payload string `json:"payload"`
	// Payments provider token, obtained via Botfather
	ProviderToken string `json:"provider_token"`
	// Three-letter ISO 4217 currency code, see more on https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency"`
	// Price breakdown, a JSON-serialized list of components
	// (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	Prices []LabeledPrice `json:"prices"`
	// The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double).
	// For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145.
	// See the exp parameter in https://core.telegram.org/bots/payments/currencies.json, it shows the
	// number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
	MaxTipAmount int `json:"max_tip_amount"`
	// A JSON-serialized array of suggested amounts of tips in the smallest units of the
	// currency (integer, not float/double). At most 4 suggested tip amounts can be specified.
	// The suggested tip amounts must be positive, passed in a strictly increased order
	// and must not exceed max_tip_amount.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts"`
	// Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button,
	// allowing multiple users to pay directly from the forwarded message, using the same invoice.
	// If non-empty, forwarded copies of the sent message will have a URL button with a deep link
	// to the bot (instead of a Pay button), with the value used as the start parameter
	StartParameter string `json:"start_parameter"`
	// A JSON-serialized data about the invoice, which will be shared with the payment provider.
	// A detailed description of required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data"`
	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	// People like it better when they see what they are paying for.
	PhotoUrl    string `json:"photo_url"`
	PhotoSize   int    `json:"photo_size"`
	PhotoWidth  int    `json:"photo_width"`
	PhotoHeight int    `json:"photo_height"`
	// Pass true, if you require the user's full name to complete the order
	NeedName bool `json:"need_name"`
	// Pass true, if you require the user's phone number to complete the order
	NeedPhoneNumber bool `json:"need_phone_number"`
	// Pass true, if you require the user's email address to complete the order
	NeedEmail bool `json:"need_email"`
	// Pass true, if you require the user's shipping address to complete the order
	NeedShippingAddress bool `json:"need_shipping_address"`
	// Pass true, if user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider"`
	// Pass true, if user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider"`
	// Pass true, if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content"`
	// If the message is a reply, ID of the original message
	ReplyToMessageId int `json:"reply_to_message_id"`
	// Pass true, if the message should be sent even if the specified replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply"`
	InlineKeyboard
}

func (s SendInvoiceData) Send(b Bot) (response Response, err error) {
	return Request("sendInvoice", b, s, &ResponseImpl{Result: &Message{}})
}

func (s SendInvoiceData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ChatId": s.ChatId, "Title": s.Title,
		"Description": s.Description, "Payload": s.Payload, "ProviderToken": s.ProviderToken,
		"Currency": s.Currency, "Prices": s.Prices})
}

type AnswerShippingQueryData struct {
	ShippingQueryId string            `json:"shipping_query_id"`
	Ok              bool              `json:"ok"`
	ShippingOptions []ShippingOptions `json:"shipping_options"`
	ErrorMessage    string            `json:"error_message"`
}

func (a AnswerShippingQueryData) Send(b Bot) (response Response, err error) {
	return Request("answerShippingQuery", b, a, &ResponseImpl{})
}

func (a AnswerShippingQueryData) Check() error {
	return globalEmptyFieldChecker(map[string]any{"ShippingQueryId": a.ShippingQueryId, "Ok": a.Ok})
}

type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`
	Ok                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message"`
}

func (a AnswerPreCheckoutQuery) Send(b Bot) (response Response, err error) {
	return Request("answerPreCheckoutQuery", b, a, &ResponseImpl{})
}

func (a AnswerPreCheckoutQuery) Check() error {
	return globalEmptyFieldChecker(map[string]any{"PreCheckoutQueryId": a.PreCheckoutQueryId, "Ok": a.Ok})
}
