package gogram

import (
	"errors"
	"os"
)

// SendText sends message to a User.
// b Bot parameter indicated which bot to send
// the message with. This way you can send messages with different bots
// text is the message that will be sent
// pass nil or *TextOptionalParams struct to optionalParams to add optional
// parameters to request
func (r *ReplyAble) SendText(b Bot, text string, optionalParams *TextOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId int    `json:"chat_id"`
		Text   string `json:"text"`
	}
	d := data{ChatId: r.Id, Text: text}
	res, err := request("sendMessage", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendPhoto(b Bot, photo interface{},
	optionalParams *PhotoOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Photo  interface{} `json:"photo"`
	}
	d := data{ChatId: r.Id, Photo: photo}
	res, err := request("sendPhoto", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendVideo(b Bot, video interface{},
	optionalParams *VideoOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Video  interface{} `json:"video"`
	}
	d := data{ChatId: r.Id, Video: video}
	res, err := request("sendVideo", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

// SendAudio sends audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format.
// On success, the sent Message is returned.
// Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
func (r *ReplyAble) SendAudio(b Bot, audio interface{},
	optionalParams *AudioOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Audio  interface{} `json:"audio"`
	}
	d := data{ChatId: r.Id, Audio: audio}
	res, err := request("sendAudio", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendDocument(b Bot, document interface{},
	optionalParams *DocumentOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId   int         `json:"chat_id"`
		Document interface{} `json:"document"`
	}
	d := data{ChatId: r.Id, Document: document}
	res, err := request("sendDocument", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

// SendVoice sends audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with
// OPUS (other formats may be sent as Audio or Document).
// On success, the sent Message is returned.
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (r *ReplyAble) SendVoice(b Bot, voice interface{},
	optionalParams *VoiceOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Voice  interface{} `json:"voice"`
	}
	d := data{ChatId: r.Id, Voice: voice}
	res, err := request("sendVoice", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendAnimation(b Bot, animation interface{},
	optionalParams *AnimationOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId    int         `json:"chat_id"`
		Animation interface{} `json:"animation"`
	}
	d := data{ChatId: r.Id, Animation: animation}
	res, err := request("sendAnimation", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendDice(b Bot, optionalParams *DiceOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("sendDice", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendVideoNote(b Bot, videoNote interface{},
	optionalParams *VideoNoteOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId    int         `json:"chat_id"`
		VideoNote interface{} `json:"videoNote"`
	}
	d := data{ChatId: r.Id, VideoNote: videoNote}
	res, err := request("sendVideoNote", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendPoll(b Bot, question string, options []string,
	optionalParams *PollOP) (response *MessageResponse, err error) {
	if options == nil {
		return nil, errors.New("options slice is empty")
	}
	type data struct {
		ChatId   int      `json:"chat_id"`
		Question string   `json:"question"`
		Options  []string `json:"options"`
	}
	d := data{ChatId: r.Id, Question: question, Options: options}
	res, err := request("sendPoll", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

// SendMediaGroup sends a group of photos, videos, documents or audios as an album.
// Documents and audio files can be only grouped in an album with messages of the same type.
// On success, an array of Messages that were sent is returned.
// You can add file_ids as string to send a media that exists on the Telegram servers (recommended),
// HTTP URLs as string for Telegram to get a media from the Internet, or a file of type *os.File to
// photos, videos, documents and audios slices.
func (r *ReplyAble) SendMediaGroup(b Bot, media []InputMedia,
	optionalParams *MediaGroupOP) (response *SliceMessageResponse, err error) {
	if len(media) == 0 {
		return &SliceMessageResponse{}, errors.New("media slice is empty. pass media a slice of structs of type " +
			"InputMediaPhoto, InputMediaVideo, InputMediaDocument, InputMediaAudio or InputMediaAnimation")
	}
	for _, j := range media {
		if j.checkInputMedia() != nil {
			return &SliceMessageResponse{}, err
		}
	}
	type data struct {
		ChatId int          `json:"chat_id"`
		Media  []InputMedia `json:"media"`
	}
	d := data{ChatId: r.Id, Media: media}
	res, err := request("sendMediaGroup", b, &d, optionalParams, &SliceMessageResponse{})
	return res.(*SliceMessageResponse), err
}

func (r *ReplyAble) SendLocation(b Bot, latitude float64, longitude float64,
	optionalParams *LocationOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId    int     `json:"chat_id"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	d := data{ChatId: r.Id, Latitude: latitude, Longitude: longitude}
	res, err := request("sendLocation", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendContact(b Bot, phoneNumber string, firstName string,
	optionalParams *ContactOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId      int    `json:"chat_id"`
		PhoneNumber string `json:"phone_number"`
		FirstName   string `json:"first_name"`
	}
	d := data{ChatId: r.Id, PhoneNumber: phoneNumber, FirstName: firstName}
	res, err := request("sendContact", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

// SendChatAction tells the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot,
// Telegram clients clear its typing status). Returns True on success.
// action parameter is the type of the action. Choose one, depending on what the user is about to receive:
// "typing" for text messages, "upload_photo" for photos, "record_video" or "upload_video" for videos,
// "record_voice" or "upload_voice" for voice notes, "upload_document" for "general" files,
// "find_location" for location data, "record_video_note" or "upload_video_note" for video notes.
func (r *ReplyAble) SendChatAction(b Bot, action string) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int    `json:"chat_id"`
		Action string `json:"action"`
	}
	var actions = []string{"typing", "upload_photo", "record_video", "upload_video", "general",
		"upload_document", "upload_voice", "record_voice", "find_location", "record_video_note", "upload_video_note"}
	for _, v := range actions {
		if v == action {
			d := data{ChatId: r.Id, Action: action}
			res, err := request("sendChatAction", b, &d, nil, &BooleanResponse{})
			return res.(*BooleanResponse), err
		}
	}
	return nil, errors.New(action + " is an unknown action, read the document.")
}

func (r *ReplyAble) ForwardMessage(b Bot, targetChatId int, messageId int,
	optionalParams *ForwardMessageOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId     int `json:"chat_id"`
		FromChatId int `json:"from_chat_id"`
		MessageId  int `json:"message_id"`
	}
	d := data{ChatId: targetChatId, FromChatId: r.Id, MessageId: messageId}
	res, err := request("forwardMessage", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

// CopyMessage copies messages of any kind.
// Service messages and invoice messages can't be copied.
// The method is analogous to the method forwardMessage,
// but the copied message doesn't have a link to the original
// message. Returns the MessageId of the sent message on success.
func (r *ReplyAble) CopyMessage(b Bot, targetChatId int, messageId int,
	optionalParams *CopyMessageOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId     int `json:"chat_id"`
		FromChatId int `json:"from_chat_id"`
		MessageId  int `json:"message_id"`
	}
	d := data{ChatId: targetChatId, FromChatId: r.Id, MessageId: messageId}
	res, err := request("forwardMessage", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) GetUserProfilePhotos(b Bot,
	optionalParams *GetUserProfilePhotosOP) (response *UserProfileResponse, err error) {
	type data struct {
		UserId int `json:"user_id"`
	}
	d := data{UserId: r.Id}
	res, err := request("getUserProfilePhotos", b, &d, optionalParams, &UserProfileResponse{})
	return res.(*UserProfileResponse), err
}

// GetFile gets basic info about a file and prepare it for downloading.
// For the moment, bots can download files of up to 20 MB in size.
// On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
func (f *File) GetFile(b Bot, fileId string) (response *FileResponse, err error) {
	type data struct {
		FileId string `json:"file_id"`
	}
	d := data{FileId: fileId}
	res, err := request("getFile", b, &d, nil, &FileResponse{})
	return res.(*FileResponse), err
}

// BanChatMember bans a user in a group, a supergroup or a channel.
// In the case of supergroups and channels, the user will not be able
// to return to the chat on their own using invite links, etc.,
// unless unbanned first.
// The bot must be an administrator in the chat for this to work and must
// have the appropriate admin rights.
func (r *User) BanChatMember(b Bot, chatId int,
	optionalParams *BanChatMemberOP) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
		UserId int `json:"user_id"`
	}
	d := data{ChatId: chatId, UserId: r.Id}
	res, err := request("banChatMember", b, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

// UnbanChatMember unbans a previously banned user in a supergroup or channel.
// The user will not return to the group or channel automatically,
// but will be able to join via link, etc.
// The bot must be an administrator for this to work.
// This method guarantees that after the call the user is not a member of the chat,
// but will be able to join it. So if the user is a member of the chat they will also be removed
// from the chat. If you don't want this, set onlyIfBanned to true, otherwise set to false.
func (r *User) UnbanChatMember(b Bot, chatId int, onlyIfBanned bool) (response *BooleanResponse, err error) {
	type data struct {
		ChatId       int  `json:"chat_id"`
		UserId       int  `json:"user_id"`
		OnlyIfBanned bool `json:"only_if_banned"`
	}
	d := data{ChatId: chatId, UserId: r.Id, OnlyIfBanned: onlyIfBanned}
	res, err := request("unbanChatMember", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *User) RestrictChatMember(b Bot, chatId int, permissions ChatPermissions,
	optionalParams *RestrictChatMemberOP) (response *BooleanResponse, err error) {
	type data struct {
		ChatId      int             `json:"chat_id"`
		UserId      int             `json:"user_id"`
		Permissions ChatPermissions `json:"permissions"`
	}
	d := data{ChatId: chatId, UserId: r.Id, Permissions: permissions}
	res, err := request("restrictChatMember", b, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *User) PromoteChatMember(b Bot, chatId int,
	optionalParams *PromoteChatMemberOP) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
		UserId int `json:"user_id"`
	}
	d := data{ChatId: chatId, UserId: r.Id}
	res, err := request("promoteChatMember", b, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

// SetChatAdministratorCustomTitle sets a custom title for an administrator in a supergroup promoted by the bot.
func (r *User) SetChatAdministratorCustomTitle(b Bot, chatId int, customTitle string) (response *BooleanResponse, err error) {
	type data struct {
		ChatId      int    `json:"chat_id"`
		UserId      int    `json:"user_id"`
		CustomTitle string `json:"custom_title"`
	}
	d := data{ChatId: chatId, UserId: r.Id, CustomTitle: customTitle}
	res, err := request("setChatAdministratorCustomTitle", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *User) SetChatPermissions(b Bot, permissions ChatPermissions) (response *BooleanResponse, err error) {
	type data struct {
		ChatId      int             `json:"chat_id"`
		Permissions ChatPermissions `json:"permissions"`
	}
	d := data{ChatId: r.Id, Permissions: permissions}
	res, err := request("setChatPermissions", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) ExportChatInviteLink(b Bot) (response *MapResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("exportChatInviteLink", b, &d, nil, &MapResponse{})
	return res.(*MapResponse), err
}

func (r *Chat) CreateChatInviteLink(b Bot,
	optionalParams *CreateChatInviteLinkOP) (response *InviteLinkResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("createChatInviteLink", b, &d, optionalParams, &InviteLinkResponse{})
	return res.(*InviteLinkResponse), err
}

func (r *Chat) EditChatInviteLink(b Bot, inviteLink string,
	optionalParams *EditChatInviteLinkOP) (response *InviteLinkResponse, err error) {
	type data struct {
		ChatId     int    `json:"chat_id"`
		InviteLink string `json:"invite_link"`
	}
	d := data{ChatId: r.Id, InviteLink: inviteLink}
	res, err := request("editChatInviteLink", b, &d, optionalParams, &InviteLinkResponse{})
	return res.(*InviteLinkResponse), err
}

func (r *Chat) RevokeChatInviteLink(b Bot, inviteLink string) (response *InviteLinkResponse, err error) {
	type data struct {
		ChatId     int    `json:"chat_id"`
		InviteLink string `json:"invite_link"`
	}
	d := data{ChatId: r.Id, InviteLink: inviteLink}
	res, err := request("revokeChatInviteLink", b, &d, nil, &InviteLinkResponse{})
	return res.(*InviteLinkResponse), err
}

func (r *Chat) SetChatPhoto(b Bot, photo *os.File) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int      `json:"chat_id"`
		Photo  *os.File `json:"photo"`
	}
	d := data{ChatId: r.Id, Photo: photo}
	res, err := request("setChatPhoto", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) DeleteChatPhoto(b Bot) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("deleteChatPhoto", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) SetChatTitle(b Bot, title string) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int    `json:"chat_id"`
		Title  string `json:"title"`
	}
	d := data{ChatId: r.Id, Title: title}
	res, err := request("setChatTitle", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) SetChatDescription(b Bot, description string) (response *BooleanResponse, err error) {
	type data struct {
		ChatId      int    `json:"chat_id"`
		Description string `json:"description"`
	}
	d := data{ChatId: r.Id, Description: description}
	res, err := request("setChatDescription", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) PinChatMessage(b Bot, messageId int,
	optionalParams *PinChatMessageOP) (response *BooleanResponse, err error) {
	type data struct {
		ChatId    int `json:"chat_id"`
		MessageId int `json:"message_id"`
	}
	d := data{ChatId: r.Id, MessageId: messageId}
	res, err := request("pinChatMessage", b, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) UnpinChatMessage(b Bot,
	optionalParams *UnpinChatMessageOP) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("unpinChatMessage", b, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) UnpinAllChatMessages(b Bot) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("unpinAllChatMessages", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) LeaveChat(b Bot) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("leaveChat", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

// GetChat gets up-to-date information about the chat (current name of the user
// for one-on-one conversations, current username of a user, group or channel, etc.)
func (r *Chat) GetChat(b Bot) (response *ChatResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("getChat", b, &d, nil, &ChatResponse{})
	return res.(*ChatResponse), err
}

func (r *Chat) GetChatAdministrators(b Bot) (response *ChatMemberResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("getChatAdministrators", b, &d, nil, &ChatMemberResponse{})
	member := res.(*ChatMemberResponse)
	member.permissionSetter()
	return member, err
}

func (r *Chat) GetChatMemberCount(b Bot) (response *IntResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("getChatMemberCount", b, &d, nil, &IntResponse{})
	return res.(*IntResponse), err
}

func (r *Chat) GetChatMember(b Bot, userId int) (response *ChatMemberResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
		UserId int `json:"user_id"`
	}
	d := data{ChatId: r.Id, UserId: userId}
	res, err := request("getChatMember", b, &d, nil, &ChatMemberResponse{})
	return res.(*ChatMemberResponse), err
}

func (r *Chat) SetChatStickerSet(b Bot, stickerSetName string) (response *BooleanResponse, err error) {
	type data struct {
		ChatId         int    `json:"chat_id"`
		StickerSetName string `json:"sticker_set_name"`
	}
	d := data{ChatId: r.Id, StickerSetName: stickerSetName}
	res, err := request("setChatStickerSet", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) DeleteChatStickerSet(b Bot) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("deleteChatStickerSet", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *ReplyAble) AnswerCallbackQuery(b Bot, callbackQueryId string,
	optionalParams *AnswerCallbackQueryOP) (response *BooleanResponse, err error) {
	type data struct {
		CallbackQueryId string `json:"callback_query_id"`
	}
	d := data{CallbackQueryId: callbackQueryId}
	res, err := request("answerCallbackQuery", b, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (m *Message) EditMessageText(b Bot, text string,
	optionalParams *EditMessageTextOP) (response *MapResponse, err error) {
	type data struct {
		Text      string `json:"text"`
		ChatId    int    `json:"chat_id"`
		MessageId int    `json:"message_id"`
	}
	d := data{Text: text, ChatId: m.Chat.Id, MessageId: m.MessageId}
	res, err := request("editMessageText", b, &d, &optionalParams, &MapResponse{})
	return res.(*MapResponse), err
}

func (c *CallbackQuery) EditMessageText(b Bot, text string,
	optionalParams *EditMessageTextOP) (response *MapResponse, err error) {
	type data struct {
		Text            string `json:"text"`
		InlineMessageId string `json:"inline_message_id"`
	}
	d := data{Text: text, InlineMessageId: c.InlineMessageId}
	res, err := request("editMessageText", b, &d, &optionalParams, &MapResponse{})
	return res.(*MapResponse), err
}

// EditMessageCaption edits captions of messages.
// On success, if the edited message is not an inline message, MapResponse's Result is the
// edited Message as a string, otherwise MapResponse's Result is True as a string.
func (m *Message) EditMessageCaption(b Bot,
	optionalParams *EditMessageCaptionOP) (response *MapResponse, err error) {
	type data struct {
		ChatId    int `json:"chat_id"`
		MessageId int `json:"message_id"`
	}
	d := data{ChatId: m.Chat.Id, MessageId: m.MessageId}
	res, err := request("editMessageCaption", b, &d, &optionalParams, &MapResponse{})
	return res.(*MapResponse), err
}

// EditMessageCaption edits captions of messages.
// On success, if the edited message is not an inline message, MapResponse's Result is the
// edited Message as a string, otherwise MapResponse's Result is True as a string.
func (c *CallbackQuery) EditMessageCaption(b Bot,
	optionalParams *EditMessageCaptionOP) (response *MapResponse, err error) {
	type data struct {
		InlineMessageId string `json:"inline_message_id"`
	}
	d := data{InlineMessageId: c.InlineMessageId}
	res, err := request("editMessageCaption", b, &d, &optionalParams, &MapResponse{})
	return res.(*MapResponse), err
}

// EditMessageMedia edits animation, audio, document, photo, or video messages.
// If a message is part of a message album, then it can be edited only to an audio for audio albums,
// only to a document for document albums and to a photo or a video otherwise.
// When an inline message is edited, a new file can't be uploaded; use a previously
// uploaded file via its file_id or specify a URL.
// pass media a type of InputMediaAudio, InputMediaPhoto, InputMediaVideo or InputMediaDocument and make sure
// Media field of them is not empty. Media field can be file_id, URL or file.
// On success, if the edited message is not an inline message, MapResponse's Result is the
// edited Message as a string, otherwise MapResponse's Result is True as a string.
func (m *Message) EditMessageMedia(b Bot, media InputMedia,
	optionalParams *EditMessageMediaOP) (response *MapResponse, err error) {
	if media == nil {
		return &MapResponse{}, errors.New("media is nil. pass media a struct of type " +
			"InputMediaPhoto, InputMediaVideo, InputMediaDocument, InputMediaAudio or InputMediaAnimation")
	}
	type data struct {
		ChatId    int        `json:"chat_id"`
		MessageId int        `json:"message_id"`
		Media     InputMedia `json:"media"`
	}
	if media.checkInputMedia() != nil {
		return &MapResponse{}, err
	}
	d := data{ChatId: m.Chat.Id, MessageId: m.MessageId, Media: media}
	res, err := request("editMessageMedia", b, &d, optionalParams, &MapResponse{})
	return res.(*MapResponse), err
}

// EditMessageMedia edits animation, audio, document, photo, or video messages.
// If a message is part of a message album, then it can be edited only to an audio for audio albums,
// only to a document for document albums and to a photo or a video otherwise.
// When an inline message is edited, a new file can't be uploaded; use a previously
// uploaded file via its file_id or specify a URL.
// pass media a type of InputMediaAudio, InputMediaPhoto, InputMediaVideo or InputMediaDocument and make sure
// Media field of them is not empty. Media field can be file_id, URL or file.
// On success, if the edited message is not an inline message, MapResponse's Result is the
// edited Message as a string, otherwise MapResponse's Result is True as a string.
func (c *CallbackQuery) EditMessageMedia(b Bot, media InputMedia,
	optionalParams *EditMessageMediaOP) (response *MapResponse, err error) {
	if media == nil {
		return &MapResponse{}, errors.New("media is nil. pass media a struct of type " +
			"InputMediaPhoto, InputMediaVideo, InputMediaDocument, InputMediaAudio or InputMediaAnimation")
	}
	type data struct {
		InlineMessageId string     `json:"inline_message_id"`
		Media           InputMedia `json:"media"`
	}
	if media.checkInputMedia() != nil {
		return &MapResponse{}, err
	}
	d := data{InlineMessageId: c.InlineMessageId, Media: media}
	res, err := request("editMessageMedia", b, &d, optionalParams, &MapResponse{})
	return res.(*MapResponse), err
}

// EditMessageReplyMarkup edits only the reply markup of messages.
// On success, if the edited message is not an inline message, the edited Message is returned,
// otherwise True is returned.
func (m *Message) EditMessageReplyMarkup(b Bot,
	optionalParams *EditMessageMediaOP) (response *MapResponse, err error) {
	type data struct {
		ChatId    int `json:"chat_id"`
		MessageId int `json:"message_id"`
	}
	d := data{ChatId: m.Chat.Id, MessageId: m.MessageId}
	res, err := request("editMessageReplyMarkup", b, &d, optionalParams, &MapResponse{})
	return res.(*MapResponse), err
}

// EditMessageReplyMarkup edits only the reply markup of messages.
// On success, if the edited message is not an inline message, the edited Message is returned,
// otherwise True is returned.
func (c *CallbackQuery) EditMessageReplyMarkup(b Bot,
	optionalParams *EditMessageMediaOP) (response *MapResponse, err error) {
	type data struct {
		InlineMessageId string     `json:"inline_message_id"`
		Media           InputMedia `json:"media"`
	}
	d := data{InlineMessageId: c.InlineMessageId}
	res, err := request("editMessageReplyMarkup", b, &d, optionalParams, &MapResponse{})
	return res.(*MapResponse), err
}

func (r *Chat) StopPoll(b Bot, messageId int,
	optionalParams *StopPollOP) (response *PollResponse, err error) {
	type data struct {
		ChatId    int `json:"chat_id"`
		MessageId int `json:"message_id"`
	}
	d := data{ChatId: r.Id, MessageId: messageId}
	res, err := request("stopPoll", b, &d, optionalParams, &PollResponse{})
	return res.(*PollResponse), err
}

// DeleteMessage deletes a message, including service messages, with the following limitations:
//- A message can only be deleted if it was sent less than 48 hours ago.
//- A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
//- Bots can delete outgoing messages in private chats, groups, and supergroups.
//- Bots can delete incoming messages in private chats.
//- Bots granted can_post_messages permissions can delete outgoing messages in channels.
//- If the bot is an administrator of a group, it can delete any message there.
//- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
// Returns True on success.
func (m *Message) DeleteMessage(b Bot, chatId int) (response *BooleanResponse, err error) {
	type data struct {
		ChatId    int `json:"chat_id"`
		MessageId int `json:"message_id"`
	}
	d := data{ChatId: chatId, MessageId: m.MessageId}
	res, err := request("deleteMessage", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *ReplyAble) SendSticker(b Bot, sticker interface{},
	optionalParams *SendStickerOP) (response *MessageResponse, err error) {
	type data struct {
		ChatId  int         `json:"chat_id"`
		Sticker interface{} `json:"sticker"`
	}
	d := data{ChatId: r.Id, Sticker: sticker}
	res, err := request("sendSticker", b, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (s StickerSet) GetStickerSet(b Bot) (response *StickerSetResponse, err error) {
	type data struct {
		Name string `json:"name"`
	}
	d := data{Name: s.Name}
	res, err := request("getStickerSet", b, &d, nil, &StickerSetResponse{})
	return res.(*StickerSetResponse), err
}

func (r *User) UploadStickerFile(b Bot, pngSticker *os.File) (response *FileResponse, err error) {
	type data struct {
		UserId     int      `json:"user_id"`
		PngSticker *os.File `json:"png_sticker"`
	}
	d := data{UserId: r.Id, PngSticker: pngSticker}
	res, err := request("uploadStickerFile", b, &d, nil, &FileResponse{})
	return res.(*FileResponse), err
}

func (r *User) CreateNewStickerSet(b Bot, name string, title string,
	emojis string, optionalParams *CreateNewStickerSetOP) (response *BooleanResponse, err error) {
	type data struct {
		UserId int    `json:"user_id"`
		Name   string `json:"name"`
		Title  string `json:"title"`
		Emojis string `json:"emojis"`
	}
	d := data{UserId: r.Id, Name: name, Title: title, Emojis: emojis}
	res, err := request("createNewStickerSet", b, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

// AddStickerToSet adds a new sticker to a set created by the bot.
// You must use exactly one of the fields png_sticker or tgs_sticker of AddStickerToSetOptionalParams.
// Animated stickers can be added to animated sticker sets and only to them.
// Animated sticker sets can have up to 50 stickers.
// Static sticker sets can have up to 120 stickers.
// Returns True on success.
func (r *User) AddStickerToSet(b Bot, name string,
	emojis string, optionalParams AddStickerToSetOP) (response *BooleanResponse, err error) {
	type data struct {
		UserId int    `json:"user_id"`
		Name   string `json:"name"`
		Emojis string `json:"emojis"`
	}
	d := data{UserId: r.Id, Name: name, Emojis: emojis}
	res, err := request("addStickerToSet", b, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (s Sticker) SetStickerPositionInSet(b Bot, sticker string, position int) (response *BooleanResponse, err error) {
	type data struct {
		Sticker  string `json:"sticker"`
		Position int    `json:"position"`
	}
	d := data{Sticker: sticker, Position: position}
	res, err := request("setStickerPositionInSet", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (s Sticker) DeleteStickerFromSet(b Bot, sticker string) (response *BooleanResponse, err error) {
	type data struct {
		Sticker string `json:"sticker"`
	}
	d := data{Sticker: sticker}
	res, err := request("deleteStickerFromSet", b, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *User) SetStickerSetThumb(b Bot, name string,
	optionalParams SetStickerSetThumbOP) (response *BooleanResponse, err error) {
	type data struct {
		UserId int    `json:"user_id"`
		Name   string `json:"name"`
	}
	d := data{Name: name}
	res, err := request("setStickerSetThumb", b, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}
