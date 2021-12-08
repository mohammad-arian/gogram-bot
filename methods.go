package gogram

import (
	"errors"
	"os"
	"reflect"
)

// SendText sends message to a User.
// b Bot parameter indicated which bot to send
// the message with. This way you can send messages with different bots
// text is the message that will be sent
// pass nil or *TextOptionalParams struct to optionalParams to add optional
// parameters to request
func (r *ReplyAble) SendText(b Bot, text string, optionalParams *TextOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId int    `json:"chat_id"`
		Text   string `json:"text"`
	}
	d := data{ChatId: r.Id, Text: text}
	res, err := request("sendMessage", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendPhoto(b Bot, photo interface{},
	optionalParams *PhotoOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Photo  interface{} `json:"photo"`
	}
	d := data{ChatId: r.Id, Photo: photo}
	res, err := request("sendPhoto", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendVideo(b Bot, video interface{},
	optionalParams *VideoOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Video  interface{} `json:"video"`
	}
	d := data{ChatId: r.Id, Video: video}
	res, err := request("sendVideo", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

// SendAudio sends audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format.
// On success, the sent Message is returned.
// Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
func (r *ReplyAble) SendAudio(b Bot, audio interface{},
	optionalParams *AudioOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Audio  interface{} `json:"audio"`
	}
	d := data{ChatId: r.Id, Audio: audio}
	res, err := request("sendAudio", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendDocument(b Bot, document interface{},
	optionalParams *DocumentOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId   int         `json:"chat_id"`
		Document interface{} `json:"document"`
	}
	d := data{ChatId: r.Id, Document: document}
	res, err := request("sendDocument", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

// SendVoice sends audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with
// OPUS (other formats may be sent as Audio or Document).
// On success, the sent Message is returned.
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (r *ReplyAble) SendVoice(b Bot, voice interface{},
	optionalParams *VoiceOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Voice  interface{} `json:"voice"`
	}
	d := data{ChatId: r.Id, Voice: voice}
	res, err := request("sendVoice", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendAnimation(b Bot, animation interface{},
	optionalParams *AnimationOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId    int         `json:"chat_id"`
		Animation interface{} `json:"animation"`
	}
	d := data{ChatId: r.Id, Animation: animation}
	res, err := request("sendAnimation", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendDice(b Bot, optionalParams *DiceOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("sendDice", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendVideoNote(b Bot, videoNote interface{},
	optionalParams *VideoNoteOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId    int         `json:"chat_id"`
		VideoNote interface{} `json:"videoNote"`
	}
	d := data{ChatId: r.Id, VideoNote: videoNote}
	res, err := request("sendVideoNote", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendPoll(b Bot, question string, options []string,
	optionalParams *PollOptionalParams) (response *MessageResponse, err error) {
	if options == nil {
		return nil, errors.New("options slice is empty")
	}
	type data struct {
		ChatId   int         `json:"chat_id"`
		Question interface{} `json:"question"`
		Options  []string    `json:"options"`
	}
	d := data{ChatId: r.Id, Question: question, Options: options}
	res, err := request("sendPoll", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

// SendMediaGroup sends a group of photos, videos, documents or audios as an album.
// Documents and audio files can be only grouped in an album with messages of the same type.
// On success, an array of Messages that were sent is returned.
// You can add file_ids as string to send a media that exists on the Telegram servers (recommended),
// HTTP URLs as string for Telegram to get a media from the Internet, or a file of type *os.File to
// photos, videos, documents and audios slices.
func (r *ReplyAble) SendMediaGroup(b Bot, optionalParams *MediaGroupOptionalParams, photos []InputMediaPhoto,
	videos []InputMediaVideo, documents []InputMediaDocument, audios []InputMediaAudio) (response *[]MessageResponse, err error) {
	type data struct {
		ChatId int           `json:"chat_id"`
		Media  []interface{} `json:"media"`
		Files  []*os.File
	}
	d := data{ChatId: r.Id}
	var files []*os.File
	for _, i := range photos {
		err := i.setMediaAndType(&files)
		if err != nil {
			return nil, err
		}
		d.Media = append(d.Media, i)
	}
	for _, i := range videos {
		err := i.setMediaAndType(&files)
		if err != nil {
			return nil, err
		}
		d.Media = append(d.Media, i)
	}
	for _, i := range documents {
		err := i.setMediaAndType(&files)
		if err != nil {
			return nil, err
		}
		d.Media = append(d.Media, i)
	}
	for _, i := range audios {
		err := i.setMediaAndType(&files)
		if err != nil {
			return nil, err
		}
		d.Media = append(d.Media, i)
	}
	if d.Media == nil {
		return nil, errors.New("you did not pass any media")
	}
	d.Files = files
	res, err := request("sendMediaGroup", b.Token, &d, optionalParams, &[]MessageResponse{})
	return res.(*[]MessageResponse), err
}

func (r *ReplyAble) SendLocation(b Bot, latitude float64, longitude float64,
	optionalParams *LocationOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId    int     `json:"chat_id"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	d := data{ChatId: r.Id, Latitude: latitude, Longitude: longitude}
	res, err := request("sendLocation", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) SendContact(b Bot, phoneNumber string, firstName string,
	optionalParams *ContactOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId      int    `json:"chat_id"`
		PhoneNumber string `json:"phone_number"`
		FirstName   string `json:"first_name"`
	}
	d := data{ChatId: r.Id, PhoneNumber: phoneNumber, FirstName: firstName}
	res, err := request("sendContact", b.Token, &d, optionalParams, &MessageResponse{})
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
			res, err := request("sendChatAction", b.Token, &d, nil, &BooleanResponse{})
			return res.(*BooleanResponse), err
		}
	}
	return nil, errors.New(action + " is an unknown action, read the document.")
}

func (r *ReplyAble) ForwardMessage(b Bot, targetChatId int, messageId int,
	optionalParams *ForwardMessageOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId     int `json:"chat_id"`
		FromChatId int `json:"from_chat_id"`
		MessageId  int `json:"message_id"`
	}
	d := data{ChatId: targetChatId, FromChatId: r.Id, MessageId: messageId}
	res, err := request("forwardMessage", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

// CopyMessage copies messages of any kind.
// Service messages and invoice messages can't be copied.
// The method is analogous to the method forwardMessage,
// but the copied message doesn't have a link to the original
// message. Returns the MessageId of the sent message on success.
func (r *ReplyAble) CopyMessage(b Bot, targetChatId int, messageId int,
	optionalParams *CopyMessageOptionalParams) (response *MessageResponse, err error) {
	type data struct {
		ChatId     int `json:"chat_id"`
		FromChatId int `json:"from_chat_id"`
		MessageId  int `json:"message_id"`
	}
	d := data{ChatId: targetChatId, FromChatId: r.Id, MessageId: messageId}
	res, err := request("forwardMessage", b.Token, &d, optionalParams, &MessageResponse{})
	return res.(*MessageResponse), err
}

func (r *ReplyAble) GetUserProfilePhotos(b Bot,
	optionalParams *GetUserProfilePhotosOptionalParams) (response *UserProfileResponse, err error) {
	type data struct {
		UserId int `json:"user_id"`
	}
	d := data{UserId: r.Id}
	res, err := request("getUserProfilePhotos", b.Token, &d, optionalParams, &UserProfileResponse{})
	return res.(*UserProfileResponse), err
}

// BanChatMember bans a user in a group, a supergroup or a channel.
// In the case of supergroups and channels, the user will not be able
// to return to the chat on their own using invite links, etc.,
// unless unbanned first.
// The bot must be an administrator in the chat for this to work and must
// have the appropriate admin rights.
func (r *User) BanChatMember(b Bot, chatId int,
	optionalParams *BanChatMemberOptionalParams) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
		UserId int `json:"user_id"`
	}
	d := data{ChatId: chatId, UserId: r.Id}
	res, err := request("banChatMember", b.Token, &d, optionalParams, &BooleanResponse{})
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
	res, err := request("unbanChatMember", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *User) RestrictChatMember(b Bot, chatId int, permissions ChatPermissions,
	optionalParams *RestrictChatMemberOptionalParams) (response *BooleanResponse, err error) {
	type data struct {
		ChatId      int             `json:"chat_id"`
		UserId      int             `json:"user_id"`
		Permissions ChatPermissions `json:"permissions"`
	}
	d := data{ChatId: chatId, UserId: r.Id, Permissions: permissions}
	res, err := request("restrictChatMember", b.Token, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *User) PromoteChatMember(b Bot, chatId int,
	optionalParams *PromoteChatMemberOptionalParams) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
		UserId int `json:"user_id"`
	}
	d := data{ChatId: chatId, UserId: r.Id}
	res, err := request("promoteChatMember", b.Token, &d, optionalParams, &BooleanResponse{})
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
	res, err := request("setChatAdministratorCustomTitle", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *User) SetChatPermissions(b Bot, permissions ChatPermissions) (response *BooleanResponse, err error) {
	type data struct {
		ChatId      int             `json:"chat_id"`
		Permissions ChatPermissions `json:"permissions"`
	}
	d := data{ChatId: r.Id, Permissions: permissions}
	res, err := request("setChatPermissions", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) ExportChatInviteLink(b Bot) (response *MapResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("exportChatInviteLink", b.Token, &d, nil, &MapResponse{})
	return res.(*MapResponse), err
}

func (r *Chat) CreateChatInviteLink(b Bot,
	optionalParams *CreateChatInviteLinkOptionalParams) (response *InviteLinkResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("createChatInviteLink", b.Token, &d, optionalParams, &InviteLinkResponse{})
	return res.(*InviteLinkResponse), err
}

func (r *Chat) EditChatInviteLink(b Bot, inviteLink string,
	optionalParams *EditChatInviteLinkOptionalParams) (response *InviteLinkResponse, err error) {
	type data struct {
		ChatId     int    `json:"chat_id"`
		InviteLink string `json:"invite_link"`
	}
	d := data{ChatId: r.Id, InviteLink: inviteLink}
	res, err := request("editChatInviteLink", b.Token, &d, optionalParams, &InviteLinkResponse{})
	return res.(*InviteLinkResponse), err
}

func (r *Chat) RevokeChatInviteLink(b Bot, inviteLink string) (response *InviteLinkResponse, err error) {
	type data struct {
		ChatId     int    `json:"chat_id"`
		InviteLink string `json:"invite_link"`
	}
	d := data{ChatId: r.Id, InviteLink: inviteLink}
	res, err := request("revokeChatInviteLink", b.Token, &d, nil, &InviteLinkResponse{})
	return res.(*InviteLinkResponse), err
}

func (r *Chat) SetChatPhoto(b Bot, photo *os.File) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int      `json:"chat_id"`
		Photo  *os.File `json:"photo"`
	}
	d := data{ChatId: r.Id, Photo: photo}
	res, err := request("setChatPhoto", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) DeleteChatPhoto(b Bot) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("deleteChatPhoto", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) SetChatTitle(b Bot, title string) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int    `json:"chat_id"`
		Title  string `json:"title"`
	}
	d := data{ChatId: r.Id, Title: title}
	res, err := request("setChatTitle", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) SetChatDescription(b Bot, description string) (response *BooleanResponse, err error) {
	type data struct {
		ChatId      int    `json:"chat_id"`
		Description string `json:"description"`
	}
	d := data{ChatId: r.Id, Description: description}
	res, err := request("setChatDescription", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) PinChatMessage(b Bot, messageId int,
	optionalParams *PinChatMessageOptionalParams) (response *BooleanResponse, err error) {
	type data struct {
		ChatId    int `json:"chat_id"`
		MessageId int `json:"message_id"`
	}
	d := data{ChatId: r.Id, MessageId: messageId}
	res, err := request("pinChatMessage", b.Token, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) UnpinChatMessage(b Bot,
	optionalParams *UnpinChatMessageOptionalParams) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("unpinChatMessage", b.Token, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) UnpinAllChatMessages(b Bot) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("unpinAllChatMessages", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) LeaveChat(b Bot) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("leaveChat", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

// GetChat gets up-to-date information about the chat (current name of the user
// for one-on-one conversations, current username of a user, group or channel, etc.)
func (r *Chat) GetChat(b Bot) (response *ChatResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("getChat", b.Token, &d, nil, &ChatResponse{})
	return res.(*ChatResponse), err
}

func (r *Chat) GetChatAdministrators(b Bot) (response *ChatMemberResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("getChatAdministrators", b.Token, &d, nil, &ChatMemberResponse{})
	member := res.(*ChatMemberResponse)
	member.permissionSetter()
	return member, err
}

func (r *Chat) GetChatMemberCount(b Bot) (response *IntResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("getChatMemberCount", b.Token, &d, nil, &IntResponse{})
	return res.(*IntResponse), err
}

func (r *Chat) GetChatMember(b Bot, userId int) (response *ChatMemberResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
		UserId int `json:"user_id"`
	}
	d := data{ChatId: r.Id, UserId: userId}
	res, err := request("getChatMember", b.Token, &d, nil, &ChatMemberResponse{})
	return res.(*ChatMemberResponse), err
}

func (r *Chat) SetChatStickerSet(b Bot, stickerSetName string) (response *BooleanResponse, err error) {
	type data struct {
		ChatId         int    `json:"chat_id"`
		StickerSetName string `json:"sticker_set_name"`
	}
	d := data{ChatId: r.Id, StickerSetName: stickerSetName}
	res, err := request("setChatStickerSet", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *Chat) DeleteChatStickerSet(b Bot) (response *BooleanResponse, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	res, err := request("deleteChatStickerSet", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (r *ReplyAble) AnswerCallbackQuery(b Bot, callbackQueryId string,
	optionalParams *AnswerCallbackQueryOptionalParams) (response *BooleanResponse, err error) {
	type data struct {
		CallbackQueryId string `json:"callback_query_id"`
	}
	d := data{CallbackQueryId: callbackQueryId}
	res, err := request("answerCallbackQuery", b.Token, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (b *Bot) SetMyCommands(commands []BotCommand,
	optionalParams *MyCommandsOptionalParams) (response *BooleanResponse, err error) {
	type data struct {
		Commands []BotCommand `json:"commands"`
	}
	d := data{Commands: commands}
	res, err := request("setMyCommands", b.Token, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (b *Bot) DeleteMyCommands(
	optionalParams *MyCommandsOptionalParams) (response *BooleanResponse, err error) {
	res, err := request("deleteMyCommands", b.Token, nil, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

func (b *Bot) GetMyCommands(
	optionalParams *MyCommandsOptionalParams) (response *BotCommandResponse, err error) {
	res, err := request("getMyCommands", b.Token, nil, optionalParams, &BotCommandResponse{})
	return res.(*BotCommandResponse), err
}

func EditMessageText(b Bot, text string,
	optionalParams EditMessageTextOptionalParams) (response *MapResponse, err error) {
	type data struct {
		Text string `json:"text"`
	}
	if optionalParams.ChatId == 0 && optionalParams.MessageId == 0 && optionalParams.InlineMessageId == 0 {
		return nil, errors.New("ChatId, MessageId and InlineMessageId of optionalParams" +
			" are empty. You need to set both ChatId and MessageId, or InlineMessageId")
	}
	if (optionalParams.ChatId == 0 && optionalParams.MessageId != 0) ||
		(optionalParams.ChatId != 0 && optionalParams.MessageId == 0) {
		return nil, errors.New("ChatId or MessageId of optionalParams" +
			" are empty. you need to set both ChatId and MessageId or InlineMessageId")
	}
	d := data{Text: text}
	res, err := request("editMessageText", b.Token, &d, &optionalParams, &MapResponse{})
	return res.(*MapResponse), err
}

// EditMessageCaption edits captions of messages.
// On success, if the edited message is not an inline message, MapResponse's Result is the
// edited Message as a string, otherwise MapResponse's Result is True as a string.
func EditMessageCaption(b Bot,
	optionalParams EditMessageCaptionOptionalParams) (response *MapResponse, err error) {
	if optionalParams.ChatId == 0 && optionalParams.MessageId == 0 && optionalParams.InlineMessageId == 0 {
		return nil, errors.New("ChatId, MessageId and InlineMessageId of optionalParams" +
			" are empty. You need to set both ChatId and MessageId, or InlineMessageId")
	}
	if (optionalParams.ChatId == 0 && optionalParams.MessageId != 0) ||
		(optionalParams.ChatId != 0 && optionalParams.MessageId == 0) {
		return nil, errors.New("ChatId or MessageId of optionalParams" +
			" are empty. you need to set both ChatId and MessageId or InlineMessageId")
	}
	res, err := request("editMessageCaption", b.Token, nil, &optionalParams, &MapResponse{})
	return res.(*MapResponse), err
}

// EditMessageMedia edits animation, audio, document, photo, or video messages.
// If a message is part of a message album, then it can be edited only to an audio for audio albums,
// only to a document for document albums and to a photo or a video otherwise.
// When an inline message is edited, a new file can't be uploaded; use a previously
// uploaded file via its file_id or specify a URL.
// pass media a type of InputMediaAudio, InputMediaPhoto, InputMediaVideo or InputMediaDocument and make sure
// Media field of them is not empty. Media field can be file_id, URL or file. if you are uploding a file
//
// On success, if the edited message is not an inline message, MapResponse's Result is the
// edited Message as a string, otherwise MapResponse's Result is True as a string.
func EditMessageMedia(b Bot, media interface{},
	optionalParams EditMessageMediaOptionalParams) (response *MapResponse, err error) {
	type data struct {
		Media interface{} `json:"media"`
		// at most there could be only one file but since setMediaAndType(files *[]*os.File) accepts slices
		// File is a slice. Modifying setMediaAndType() breaks SendMediaGroup()
		File []*os.File
	}
	d := data{}
	var files []*os.File
	switch v := media.(type) {
	case InputMediaAudio:
		err := v.setMediaAndType(&files)
		if err != nil {
			return nil, err
		}
		d.Media = v
	case InputMediaPhoto:
		err := v.setMediaAndType(&files)
		if err != nil {
			return nil, err
		}
		d.Media = v
	case InputMediaVideo:
		err := v.setMediaAndType(&files)
		if err != nil {
			return nil, err
		}
		d.Media = v
	case InputMediaDocument:
		err := v.setMediaAndType(&files)
		if err != nil {
			return nil, err
		}
		d.Media = v
	default:
		return nil, errors.New("pass media a type of InputMediaAudio, InputMediaPhoto, InputMediaVideo " +
			"or InputMediaDocument not " + reflect.TypeOf(media).String())
	}
	if optionalParams.ChatId == 0 && optionalParams.MessageId == 0 && optionalParams.InlineMessageId == 0 {
		return nil, errors.New("ChatId, MessageId and InlineMessageId of optionalParams" +
			" are empty. You need to set both ChatId and MessageId, or InlineMessageId")
	}
	if (optionalParams.ChatId == 0 && optionalParams.MessageId != 0) ||
		(optionalParams.ChatId != 0 && optionalParams.MessageId == 0) {
		return nil, errors.New("ChatId or MessageId of optionalParams" +
			" are empty. you need to set both ChatId and MessageId or InlineMessageId")
	}
	d.File = files
	res, err := request("editMessageMedia", b.Token, &d, optionalParams, &MapResponse{})
	return res.(*MapResponse), err
}

func (r *Chat) StopPoll(b Bot, messageId int,
	optionalParams *StopPollOptionalParams) (response *PollResponse, err error) {
	type data struct {
		ChatId    int `json:"chat_id"`
		MessageId int `json:"message_id"`
	}
	d := data{ChatId: r.Id, MessageId: messageId}
	res, err := request("stopPoll", b.Token, &d, optionalParams, &PollResponse{})
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
func (r *Chat) DeleteMessage(b Bot, messageId int) (response *BooleanResponse, err error) {
	type data struct {
		ChatId    int `json:"chat_id"`
		MessageId int `json:"message_id"`
	}
	d := data{ChatId: r.Id, MessageId: messageId}
	res, err := request("deleteMessage", b.Token, &d, nil, &BooleanResponse{})
	return res.(*BooleanResponse), err
}

// EditMessageReplyMarkup edits only the reply markup of messages.
// On success, if the edited message is not an inline message, the edited Message is returned,
// otherwise True is returned.
func (r *Chat) EditMessageReplyMarkup(b Bot,
	optionalParams EditMessageMediaOptionalParams) (response *MapResponse, err error) {
	if optionalParams.ChatId == 0 && optionalParams.MessageId == 0 && optionalParams.InlineMessageId == 0 {
		return nil, errors.New("ChatId, MessageId and InlineMessageId of optionalParams" +
			" are empty. You need to set both ChatId and MessageId, or InlineMessageId")
	}
	if (optionalParams.ChatId == 0 && optionalParams.MessageId != 0) ||
		(optionalParams.ChatId != 0 && optionalParams.MessageId == 0) {
		return nil, errors.New("ChatId or MessageId of optionalParams" +
			" are empty. you need to set both ChatId and MessageId or InlineMessageId")
	}
	res, err := request("editMessageReplyMarkup", b.Token, nil, nil, &MapResponse{})
	return res.(*MapResponse), err
}
