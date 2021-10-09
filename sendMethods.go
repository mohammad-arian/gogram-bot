package gogram

import (
	"encoding/json"
	"errors"
	"os"
)

// SendText sends message to a User.
// b Bot parameter indicated which bot to send
// the message with. This way you can send messages with different bots
// text is the message that will be sent
// pass nil or *TextOptionalParams struct to optionalParams to add optional
// parameters to request
func (r *ReplyAble) SendText(b Bot, text string, optionalParams *TextOptionalParams) (response string, err error) {
	type data struct {
		ChatId int    `json:"chat_id"`
		Text   string `json:"text"`
	}
	d := data{ChatId: r.Id, Text: text}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendMessage", b.Token, d, op)
}

func (r *ReplyAble) SendPhoto(b Bot, photo interface{}, optionalParams *PhotoOptionalParams) (response string, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Photo  interface{} `json:"photo"`
	}
	d := data{ChatId: r.Id, Photo: photo}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendPhoto", b.Token, d, op)
}

func (r *ReplyAble) SendVideo(b Bot, video interface{}, optionalParams *VideoOptionalParams) (response string, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Video  interface{} `json:"video"`
	}
	d := data{ChatId: r.Id, Video: video}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendVideo", b.Token, d, op)
}

// SendAudio sends audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format.
// On success, the sent Message is returned.
// Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
func (r *ReplyAble) SendAudio(b Bot, audio interface{}, optionalParams *AudioOptionalParams) (response string, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Audio  interface{} `json:"audio"`
	}
	d := data{ChatId: r.Id, Audio: audio}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendAudio", b.Token, d, op)
}

func (r *ReplyAble) SendDocument(b Bot, document interface{}, optionalParams *DocumentOptionalParams) (response string, err error) {
	type data struct {
		ChatId   int         `json:"chat_id"`
		Document interface{} `json:"document"`
	}
	d := data{ChatId: r.Id, Document: document}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendDocument", b.Token, d, op)
}

// SendVoice sends audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with
// OPUS (other formats may be sent as Audio or Document).
// On success, the sent Message is returned.
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (r *ReplyAble) SendVoice(b Bot, voice interface{}, optionalParams *VoiceOptionalParams) (response string, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Voice  interface{} `json:"voice"`
	}
	d := data{ChatId: r.Id, Voice: voice}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendVoice", b.Token, d, op)
}

func (r *ReplyAble) SendAnimation(b Bot, animation interface{}, optionalParams *AnimationOptionalParams) (response string, err error) {
	type data struct {
		ChatId    int         `json:"chat_id"`
		Animation interface{} `json:"animation"`
	}
	d := data{ChatId: r.Id, Animation: animation}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendAnimation", b.Token, d, op)
}

func (r *ReplyAble) SendDice(b Bot, optionalParams *DiceOptionalParams) (response string, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendDice", b.Token, d, op)
}

func (r *ReplyAble) SendVideoNote(b Bot, videoNote interface{}, optionalParams *VideoNoteOptionalParams) (response string, err error) {
	type data struct {
		ChatId    int         `json:"chat_id"`
		VideoNote interface{} `json:"videoNote"`
	}
	d := data{ChatId: r.Id, VideoNote: videoNote}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendVideoNote", b.Token, d, op)
}

func (r *ReplyAble) SendPoll(b Bot, question string, options []string, optionalParams *PollOptionalParams) (response string, err error) {
	if options == nil {
		return "", errors.New("options slice is empty")
	}
	type data struct {
		ChatId   int         `json:"chat_id"`
		Question interface{} `json:"question"`
		Options  []string    `json:"options"`
	}
	d := data{ChatId: r.Id, Question: question, Options: options}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendPoll", b.Token, d, op)
}

// SendMediaGroup sends a group of photos, videos, documents or audios as an album.
// Documents and audio files can be only grouped in an album with messages of the same type.
// On success, an array of Messages that were sent is returned.
// You can add file_ids as string to send a media that exists on the Telegram servers (recommended),
// HTTP URLs as string for Telegram to get a media from the Internet, or a file of type *os.File to
// photos, videos, documents and audios slices.
func (r *ReplyAble) SendMediaGroup(b Bot, optionalParams *MediaGroupOptionalParams, photos []interface{}, videos []interface{},
	documents []interface{}, audios []interface{}) (response string, err error) {
	type data struct {
		ChatId int    `json:"chat_id"`
		Media  string `json:"media"`
		Files  []*os.File
	}
	d := data{ChatId: r.Id}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	var media []interface{}
	for _, i := range photos {
		switch v := i.(type) {
		case *os.File:
			d.Files = append(d.Files, v)
			media = append(media, inputMediaPhoto{Media: "attach://" + v.Name(), Type: "photo"})
		case string:
			media = append(media, inputMediaPhoto{Media: v, Type: "photo"})
		}
	}
	for _, i := range videos {
		switch v := i.(type) {
		case *os.File:
			d.Files = append(d.Files, v)
			media = append(media, inputMediaVideo{Media: "attach://" + v.Name(), Type: "video"})
		case string:
			media = append(media, inputMediaVideo{Media: v, Type: "video"})
		}
	}
	for _, i := range documents {
		switch v := i.(type) {
		case *os.File:
			d.Files = append(d.Files, v)
			media = append(media, inputMediaDocument{Media: "attach://" + v.Name(), Type: "document"})
		case string:
			media = append(media, inputMediaDocument{Media: v, Type: "documents"})
		}
	}
	for _, i := range audios {
		switch v := i.(type) {
		case *os.File:
			d.Files = append(d.Files, v)
			media = append(media, inputMediaAudio{Media: "attach://" + v.Name(), Type: "audio"})
		case string:
			media = append(media, inputMediaAudio{Media: v, Type: "audio"})
		}
	}
	if media == nil {
		return "", errors.New("you did not pass any file, file_id or URL")
	}
	mediaToJson, _ := json.Marshal(media)
	d.Media = string(mediaToJson)
	return request(r.Id, "sendMediaGroup", b.Token, d, op)
}

func (r *ReplyAble) SendLocation(b Bot, latitude float64, longitude float64, optionalParams *LocationOptionalParams) (response string, err error) {
	type data struct {
		ChatId    int     `json:"chat_id"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	d := data{ChatId: r.Id, Latitude: latitude, Longitude: longitude}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendLocation", b.Token, d, op)
}

func (r *ReplyAble) SendContact(b Bot, phoneNumber string, firstName string, optionalParams *ContactOptionalParams) (response string, err error) {
	type data struct {
		ChatId      int    `json:"chat_id"`
		PhoneNumber string `json:"phone_number"`
		FirstName   string `json:"first_name"`
	}
	d := data{ChatId: r.Id, PhoneNumber: phoneNumber, FirstName: firstName}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "sendContact", b.Token, d, op)
}

// SendChatAction tells the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot,
// Telegram clients clear its typing status). Returns True on success.
// action parameter is the type of the action. Choose one, depending on what the user is about to receive:
// "typing" for text messages, "upload_photo" for photos, "record_video" or "upload_video" for videos,
// "record_voice" or "upload_voice" for voice notes, "upload_document" for "general" files,
// "find_location" for location data, "record_video_note" or "upload_video_note" for video notes.
func (r *ReplyAble) SendChatAction(b Bot, action string) (response string, err error) {
	type data struct {
		ChatId int    `json:"chat_id"`
		Action string `json:"action"`
	}
	var actions = []string{"typing", "upload_photo", "record_video", "upload_video", "general",
		"upload_document", "upload_voice", "record_voice", "find_location", "record_video_note", "upload_video_note"}
	for _, v := range actions {
		if v == action {
			d := data{ChatId: r.Id, Action: action}
			return request(r.Id, "sendChatAction", b.Token, d, nil)
		}
	}
	return "", errors.New(action + " is an unknown action, read the document.")
}

func (r *ReplyAble) ForwardMessage(b Bot, targetChatId int, messageId int,
	optionalParams *ForwardMessageOptionalParams) (response string, err error) {
	type data struct {
		ChatId     int `json:"chat_id"`
		FromChatId int `json:"from_chat_id"`
		MessageId  int `json:"message_id"`
	}
	d := data{ChatId: targetChatId, FromChatId: r.Id, MessageId: messageId}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "forwardMessage", b.Token, d, op)
}

// CopyMessage copies messages of any kind.
// Service messages and invoice messages can't be copied.
// The method is analogous to the method forwardMessage,
// but the copied message doesn't have a link to the original
// message. Returns the MessageId of the sent message on success.
func (r *ReplyAble) CopyMessage(b Bot, targetChatId int, messageId int,
	optionalParams *CopyMessageOptionalParams) (response string, err error) {
	type data struct {
		ChatId     int `json:"chat_id"`
		FromChatId int `json:"from_chat_id"`
		MessageId  int `json:"message_id"`
	}
	d := data{ChatId: targetChatId, FromChatId: r.Id, MessageId: messageId}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "forwardMessage", b.Token, d, op)
}

func (r *ReplyAble) GetUserProfilePhotos(b Bot,
	optionalParams *GetUserProfilePhotosOptionalParams) (response UserProfileResponse, err error) {
	type data struct {
		UserId int `json:"user_id"`
	}
	d := data{UserId: r.Id}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	res, err := request(r.Id, "getUserProfilePhotos", b.Token, d, op)
	if err == nil {
		u := UserProfileResponse{}
		err := json.Unmarshal([]byte(res), &u)
		if err != nil {
			return u, err
		}
		return u, nil
	}
	return UserProfileResponse{}, err
}

// BanChatMember bans a user in a group, a supergroup or a channel.
// In the case of supergroups and channels, the user will not be able
// to return to the chat on their own using invite links, etc.,
// unless unbanned first.
// The bot must be an administrator in the chat for this to work and must
// have the appropriate admin rights.
func (r *User) BanChatMember(b Bot, chatId int,
	optionalParams *BanChatMemberOptionalParams) (response string, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
		UserId int `json:"user_id"`
	}
	d := data{ChatId: chatId, UserId: r.Id}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "banChatMember", b.Token, d, op)
}

// UnbanChatMember unbans a previously banned user in a supergroup or channel.
// The user will not return to the group or channel automatically,
// but will be able to join via link, etc.
// The bot must be an administrator for this to work.
// This method guarantees that after the call the user is not a member of the chat,
// but will be able to join it. So if the user is a member of the chat they will also be removed
// from the chat. If you don't want this, set onlyIfBanned to true, otherwise set to false.
func (r *User) UnbanChatMember(b Bot, chatId int, onlyIfBanned bool,
	optionalParams *BanChatMemberOptionalParams) (response string, err error) {
	type data struct {
		ChatId       int  `json:"chat_id"`
		UserId       int  `json:"user_id"`
		OnlyIfBanned bool `json:"only_if_banned"`
	}
	d := data{ChatId: chatId, UserId: r.Id, OnlyIfBanned: onlyIfBanned}
	var op interface{}
	if optionalParams != nil {
		op = *optionalParams
	}
	return request(r.Id, "unbanChatMember", b.Token, d, op)
}
