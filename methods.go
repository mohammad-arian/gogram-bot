package gogram

// SendText sends message to a User.
func (r *ReplyAble) SendText(b Bot, data TextData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

func (r *ReplyAble) SendPhoto(b Bot, data PhotoData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

func (r *ReplyAble) SendVideo(b Bot, data VideoData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

// SendAudio sends audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format.
// On success, the sent Message is returned.
// Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
func (r *ReplyAble) SendAudio(b Bot, data AudioData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

func (r *ReplyAble) SendDocument(b Bot, data DocumentData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

// SendVoice sends audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with
// OPUS (other formats may be sent as Audio or Document).
// On success, the sent Message is returned.
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (r *ReplyAble) SendVoice(b Bot, data VoiceData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

func (r *ReplyAble) SendAnimation(b Bot, data AnimationData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

func (r *ReplyAble) SendDice(b Bot, data DiceData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

func (r *ReplyAble) SendVideoNote(b Bot, data VideoNoteData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

func (r *ReplyAble) SendPoll(b Bot, data PollData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

// SendMediaGroup sends a group of photos, videos, documents or audios as an album.
// Documents and audio files can be only grouped in an album with messages of the same type.
// On success, an array of Messages that were sent is returned.
// You can add file_ids as string to send a media that exists on the Telegram servers (recommended),
// HTTP URLs as string for Telegram to get a media from the Internet, or a file of type *os.File to
// photos, videos, documents and audios slices.
func (r *ReplyAble) SendMediaGroup(b Bot, data MediaGroupData) (response *SliceMessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

func (r *ReplyAble) SendLocation(b Bot, data LocationData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

func (r *ReplyAble) SendContact(b Bot, data ContactData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

// SendChatAction tells the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot,
// Telegram clients clear its typing status). Returns True on success.
// action parameter is the type of the action. Choose one, depending on what the user is about to receive:
// "typing" for text messages, "upload_photo" for photos, "record_video" or "upload_video" for videos,
// "record_voice" or "upload_voice" for voice notes, "upload_document" for "general" files,
// "find_location" for location data, "record_video_note" or "upload_video_note" for video notes.
func (r *ReplyAble) SendChatAction(b Bot, data SendChatActionData) (response *BooleanResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

func (m *Message) ForwardMessage(b Bot, data ForwardMessageData) (response *MessageResponse, err error) {
	data.MessageId = m.MessageId
	return data.Send(b)
}

// CopyMessage copies messages of any kind.
// Service messages and invoice messages can't be copied.
// The method is analogous to the method forwardMessage,
// but the copied message doesn't have a link to the original
// message. Returns the MessageId of the sent message on success.
func (m *Message) CopyMessage(b Bot, data CopyMessageData) (response *MessageResponse, err error) {
	data.MessageId = m.MessageId
	return data.Send(b)
}

func (u *User) GetUserProfilePhotos(b Bot, data UserProfilePhotosData) (response *UserProfileResponse, err error) {
	data.UserId = u.Id
	return data.Send(b)
}

// GetFile gets basic info about a file and prepare it for downloading.
// For the moment, bots can download files of up to 20 MB in size.
// On success, a File object is returned.
//The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>,
//where <file_path> is taken from the response.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile again.
func (f *File) GetFile(b Bot) (response *FileResponse, err error) {
	d := GetFileData{FileId: f.fileId}
	return d.Send(b)
}

// BanChatMember bans a user in a group, a supergroup or a channel.
// In the case of supergroups and channels, the user will not be able
// to return to the chat on their own using invite links, etc.,
// unless unbanned first.
// The bot must be an administrator in the chat for this to work and must
// have the appropriate admin rights.
func (u *User) BanChatMember(b Bot, data BanChatMemberData) (response *BooleanResponse, err error) {
	data.UserId = u.Id
	return data.Send(b)
}

// UnbanChatMember unbans a previously banned user in a supergroup or channel.
// The user will not return to the group or channel automatically,
// but will be able to join via link, etc.
// The bot must be an administrator for this to work.
// This method guarantees that after the call the user is not a member of the chat,
// but will be able to join it. So if the user is a member of the chat they will also be removed
// from the chat. If you don't want this, set onlyIfBanned to true, otherwise set to false.
func (u *User) UnbanChatMember(b Bot, data UnbanChatMemberData) (response *BooleanResponse, err error) {
	data.UserId = u.Id
	return data.Send(b)
}

func (u *User) RestrictChatMember(b Bot, data RestrictChatMemberData) (response *BooleanResponse, err error) {
	data.UserId = u.Id
	return data.Send(b)
}

func (u *User) PromoteChatMember(b Bot, data PromoteChatMemberData) (response *BooleanResponse, err error) {
	data.UserId = u.Id
	return data.Send(b)
}

// SetChatAdministratorCustomTitle sets a custom title for an administrator in a supergroup promoted by the bot.
func (u *User) SetChatAdministratorCustomTitle(b Bot,
	data SetChatAdministratorCustomTitleData) (response *BooleanResponse, err error) {
	data.UserId = u.Id
	return data.Send(b)
}

func (c *Chat) SetChatPermissions(b Bot, data SetChatPermissionsData) (response *BooleanResponse, err error) {
	data.ChatId = c.Id
	return data.Send(b)
}

func (c *Chat) ExportChatInviteLink(b Bot) (response *MapResponse, err error) {
	d := ExportChatInviteLinkData{ChatId: c.Id}
	return d.Send(b)
}

func (c *Chat) CreateChatInviteLink(b Bot, data CreateChatInviteLinkData) (response *InviteLinkResponse, err error) {
	data.ChatId = c.Id
	return data.Send(b)
}

func (c *Chat) EditChatInviteLink(b Bot, data EditChatInviteLinkData) (response *InviteLinkResponse, err error) {
	data.ChatId = c.Id
	return data.Send(b)
}

func (c *Chat) RevokeChatInviteLink(b Bot, data RevokeChatInviteLinkData) (response *InviteLinkResponse, err error) {
	data.ChatId = c.Id
	return data.Send(b)
}

func (c *Chat) SetChatPhoto(b Bot, data SetChatPhotoData) (response *BooleanResponse, err error) {
	data.ChatId = c.Id
	return data.Send(b)
}

func (c *Chat) DeleteChatPhoto(b Bot) (response *BooleanResponse, err error) {
	d := DeleteChatPhotoData{ChatId: c.Id}
	return d.Send(b)
}

func (c *Chat) SetChatTitle(b Bot, data SetChatTitleData) (response *BooleanResponse, err error) {
	data.ChatId = c.Id
	return data.Send(b)
}

func (c *Chat) SetChatDescription(b Bot, data SetChatDescriptionData) (response *BooleanResponse, err error) {
	data.ChatId = c.Id
	return data.Send(b)
}

func (c *Chat) PinChatMessage(b Bot, data PinChatMessageData) (response *BooleanResponse, err error) {
	data.ChatId = c.Id
	return data.Send(b)
}

func (c *Chat) UnpinChatMessage(b Bot, data UnpinChatMessageData) (response *BooleanResponse, err error) {
	data.ChatId = c.Id
	return data.Send(b)
}

func (c *Chat) UnpinAllChatMessages(b Bot) (response *BooleanResponse, err error) {
	d := UnpinAllChatMessagesData{ChatId: c.Id}
	return d.Send(b)
}

func (c *Chat) Leave(b Bot) (response *BooleanResponse, err error) {
	d := LeaveData{ChatId: c.Id}
	return d.Send(b)
}

// GetChat gets up-to-date information about the chat (current name of the user
// for one-on-one conversations, current username of a user, group or channel, etc.)
func (c *Chat) GetChat(b Bot) (response *ChatResponse, err error) {
	d := GetChatData{ChatId: c.Id}
	return d.Send(b)
}

func (c *Chat) GetChatAdministrators(b Bot) (response *ChatMemberResponse, err error) {
	d := GetChatAdministratorsData{ChatId: c.Id}
	return d.Send(b)
}

func (c *Chat) GetChatMemberCount(b Bot) (response *IntResponse, err error) {
	d := GetChatMemberCountData{ChatId: c.Id}
	return d.Send(b)
}

func (u *User) GetChatMember(b Bot, data GetChatMemberData) (response *ChatMemberResponse, err error) {
	data.UserId = u.Id
	return data.Send(b)
}

func (c *Chat) SetChatStickerSet(b Bot, data SetChatStickerSetData) (response *BooleanResponse, err error) {
	data.ChatId = c.Id
	return data.Send(b)
}

func (c *Chat) DeleteChatStickerSet(b Bot) (response *BooleanResponse, err error) {
	d := DeleteChatStickerSetData{ChatId: c.Id}
	return d.Send(b)
}

func (m *Message) EditMessageText(b Bot, data EditMessageTextData) (response *MapResponse, err error) {
	data.MessageId = m.MessageId
	return data.Send(b)
}

func (c *CallbackQuery) EditMessageText(b Bot, data EditMessageTextData) (response *MapResponse, err error) {
	data.InlineMessageId = c.Id
	return data.Send(b)
}

// EditMessageCaption edits captions of messages.
// On success, if the edited message is not an inline message, MapResponse's Result is the
// edited Message as a string, otherwise MapResponse's Result is True as a string.
func (m *Message) EditMessageCaption(b Bot, data EditMessageCaptionData) (response *MapResponse, err error) {
	data.MessageId = m.MessageId
	return data.Send(b)
}

// EditMessageCaption edits captions of messages.
// On success, if the edited message is not an inline message, MapResponse's Result is the
// edited Message as a string, otherwise MapResponse's Result is True as a string.
func (c *CallbackQuery) EditMessageCaption(b Bot, data EditMessageCaptionData) (response *MapResponse, err error) {
	data.InlineMessageId = c.Id
	return data.Send(b)
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
func (m *Message) EditMessageMedia(b Bot, data EditMessageMediaData) (response *MapResponse, err error) {
	data.MessageId = m.MessageId
	return data.Send(b)
}

// EditMessageMedia is the same as Message.EditMessageMedia(), except it sets the InlineMessageId field of
// EditMessageMediaData to CallbackQuery's id
func (c *CallbackQuery) EditMessageMedia(b Bot, data EditMessageMediaData) (response *MapResponse, err error) {
	data.InlineMessageId = c.Id
	return data.Send(b)
}

// EditMessageReplyMarkup edits only the reply markup of messages.
// On success, if the edited message is not an inline message, the edited Message is returned,
// otherwise True is returned.
func (m *Message) EditMessageReplyMarkup(b Bot, data EditMessageReplyMarkupData) (response *MapResponse, err error) {
	data.MessageId = m.MessageId
	return data.Send(b)
}

// EditMessageReplyMarkup edits only the reply markup of messages.
// On success, if the edited message is not an inline message, the edited Message is returned,
// otherwise True is returned.
func (c *CallbackQuery) EditMessageReplyMarkup(b Bot, data EditMessageReplyMarkupData) (response *MapResponse, err error) {
	data.InlineMessageId = c.Id
	return data.Send(b)
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
func (m *Message) DeleteMessage(b Bot, data DeleteMessageData) (response *BooleanResponse, err error) {
	data.MessageId = m.MessageId
	return data.Send(b)
}

func (c *Chat) StopPoll(b Bot, data StopPollData) (response *PollResponse, err error) {
	data.ChatId = c.Id
	return data.Send(b)
}

func (c *CallbackQuery) AnswerCallbackQuery(b Bot, data AnswerCallbackQueryData) (response *BooleanResponse, err error) {
	data.CallbackQueryId = c.Id
	return data.Send(b)
}

func (r *ReplyAble) SendSticker(b Bot, data SendStickerData) (response *MessageResponse, err error) {
	data.ChatId = r.Id
	return data.Send(b)
}

func (s StickerSet) GetStickerSet(b Bot, data GetStickerSetData) (response *StickerSetResponse, err error) {
	data.Name = s.Name
	return data.Send(b)
}

func (u *User) UploadStickerFile(b Bot, data UploadStickerFileData) (response *FileResponse, err error) {
	data.UserId = u.Id
	return data.Send(b)
}

func (u *User) CreateNewStickerSet(b Bot, data CreateNewStickerSetData) (response *BooleanResponse, err error) {
	data.UserId = u.Id
	return data.Send(b)
}

// AddStickerToSet adds a new sticker to a set created by the bot.
// You must use exactly one of the fields png_sticker or tgs_sticker of AddStickerToSetOptionalParams.
// Animated stickers can be added to animated sticker sets and only to them.
// Animated sticker sets can have up to 50 stickers.
// Static sticker sets can have up to 120 stickers.
// Returns True on success.
func (u *User) AddStickerToSet(b Bot, data AddStickerToSetData) (response *BooleanResponse, err error) {
	data.UserId = u.Id
	return data.Send(b)
}

func (s Sticker) SetStickerPositionInSet(b Bot, data SetStickerPositionInSetData) (response *BooleanResponse, err error) {
	return data.Send(b)
}

func (s Sticker) DeleteStickerFromSet(b Bot, data DeleteStickerFromSetData) (response *BooleanResponse, err error) {
	return data.Send(b)
}

func (u *User) SetStickerSetThumb(b Bot, data SetStickerSetThumbData) (response *BooleanResponse, err error) {
	data.UserId = u.Id
	return data.Send(b)
}
