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
// pass nil or *TextOptionalParams struct to optionalParams. It adds some optional
// parameters to request, like reply_markup, disable_notification and ...
func (r *ReplyAble) SendText(b Bot, text string, optionalParams *TextOptionalParams) (response string, err error) {
	type data struct {
		ChatId int    `json:"chat_id"`
		Text   string `json:"text"`
	}
	d := data{ChatId: r.Id, Text: text}
	return request(r.Id, "Message", b.Token, false, d, optionalParams)
}

func (r *ReplyAble) SendPhoto(b Bot, photo interface{}, optionalParams *PhotoOptionalParams) (response string, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Photo  interface{} `json:"photo"`
	}
	d := data{ChatId: r.Id, Photo: photo}
	switch photo.(type) {
	case *os.File:
		return request(r.Id, "Photo", b.Token, true, d, optionalParams)
	case string:
		return request(r.Id, "Photo", b.Token, false, d, optionalParams)
	default:
		return "", errors.New("SendPhoto function accepts only string and *os.File types")
	}
}

func (r *ReplyAble) SendVideo(b Bot, video interface{}, optionalParams *VideoOptionalParams) (response string, err error) {
	type data struct {
		ChatId int         `json:"chat_id"`
		Video  interface{} `json:"video"`
	}
	d := data{ChatId: r.Id, Video: video}
	switch video.(type) {
	case *os.File:
		return request(r.Id, "Video", b.Token, true, d, optionalParams)
	case string:
		return request(r.Id, "Video", b.Token, false, d, optionalParams)
	default:
		return "", errors.New("SendVideo function accepts only string and *os.File types")
	}
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
	switch audio.(type) {
	case *os.File:
		return request(r.Id, "Audio", b.Token, true, d, optionalParams)
	case string:
		return request(r.Id, "Audio", b.Token, false, d, optionalParams)
	default:
		return "", errors.New("SendAudio function accepts only string and *os.File types")
	}
}

func (r *ReplyAble) SendDocument(b Bot, document interface{}, optionalParams *DocumentOptionalParams) (response string, err error) {
	type data struct {
		ChatId   int         `json:"chat_id"`
		Document interface{} `json:"document"`
	}
	d := data{ChatId: r.Id, Document: document}
	switch document.(type) {
	case *os.File:
		return request(r.Id, "Document", b.Token, true, d, optionalParams)
	case string:
		return request(r.Id, "Document", b.Token, false, d, optionalParams)
	default:
		return "", errors.New("SendDocument function accepts only string and *os.File types")
	}
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
	switch voice.(type) {
	case *os.File:
		return request(r.Id, "Voice", b.Token, true, d, optionalParams)
	case string:
		return request(r.Id, "Voice", b.Token, false, d, optionalParams)
	default:
		return "", errors.New("SendVoice function accepts only string and *os.File types")
	}
}

func (r *ReplyAble) SendAnimation(b Bot, animation interface{}, optionalParams *AnimationOptionalParams) (response string, err error) {
	type data struct {
		ChatId    int         `json:"chat_id"`
		Animation interface{} `json:"animation"`
	}
	d := data{ChatId: r.Id, Animation: animation}
	switch animation.(type) {
	case *os.File:
		return request(r.Id, "Animation", b.Token, true, d, optionalParams)
	case string:
		return request(r.Id, "Animation", b.Token, false, d, optionalParams)
	default:
		return "", errors.New("SendAnimation function accepts only string and *os.File types")
	}
}

func (r *ReplyAble) SendDice(b Bot, optionalParams *DiceOptionalParams) (response string, err error) {
	type data struct {
		ChatId int `json:"chat_id"`
	}
	d := data{ChatId: r.Id}
	return request(r.Id, "Dice", b.Token, false, d, optionalParams)
}

func (r *ReplyAble) SendVideoNote(b Bot, videoNote interface{}, optionalParams *VideoNoteOptionalParams) (response string, err error) {
	type data struct {
		ChatId    int         `json:"chat_id"`
		VideoNote interface{} `json:"videoNote"`
	}
	d := data{ChatId: r.Id, VideoNote: videoNote}
	switch videoNote.(type) {
	case *os.File:
		return request(r.Id, "VideoNote", b.Token, true, d, optionalParams)
	case string:
		return request(r.Id, "VideoNote", b.Token, false, d, optionalParams)
	default:
		return "", errors.New("SendVideoNote function accepts only string and *os.File types")
	}
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
	return request(r.Id, "Poll", b.Token, false, d, optionalParams)
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
	return request(r.Id, "MediaGroup", b.Token, true, d, optionalParams)
}

func (r *ReplyAble) SendLocation(b Bot, latitude float64, longitude float64, optionalParams *LocationOptionalParams) (response string, err error) {
	type data struct {
		ChatId    int     `json:"chat_id"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	d := data{ChatId: r.Id, Latitude: latitude, Longitude: longitude}
	return request(r.Id, "Location", b.Token, false, d, optionalParams)
}

func (r *ReplyAble) SendContact(b Bot, phoneNumber string, firstName string, optionalParams *ContactOptionalParams) (response string, err error) {
	type data struct {
		ChatId      int    `json:"chat_id"`
		PhoneNumber string `json:"phone_number"`
		FirstName   string `json:"first_name"`
	}
	d := data{ChatId: r.Id, PhoneNumber: phoneNumber, FirstName: firstName}
	return request(r.Id, "Contact", b.Token, false, d, optionalParams)
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
	d := data{ChatId: r.Id, Action: action}
	return request(r.Id, "ChatAction", b.Token, false, d, nil)
}
