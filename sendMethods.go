package gogram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
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
	type data struct {
		ChatId   int         `json:"chat_id"`
		Question interface{} `json:"question"`
		Options  []string
	}
	d := data{ChatId: r.Id, Question: question, Options: options}
	return request(r.Id, "VideoPoll", b.Token, false, d, optionalParams)
}

// SendMediaGroup sends a group of photos, videos, documents or audios as an album.
// Documents and audio files can be only grouped in an album with messages of the same type.
// On success, an array of Messages that were sent is returned.
// You can add file_ids as string to send a media that exists on the Telegram servers (recommended),
// HTTP URLs as string for Telegram to get a media from the Internet, or a file of type *os.File to
// photos, videos, documents and audios slices.
func (r *ReplyAble) SendMediaGroup(b Bot, optionalParams *MediaGroupOptionalParams, photos []interface{}, videos []interface{},
	documents []interface{}, audios []interface{}) (response string, err error) {
	var id = r.Id
	if id == 0 {
		return "", errors.New("id field is empty")
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.telegram.org/bot%s/sendMediaGroup", b.Token),
		nil)
	if err != nil {
		return "", err
	}
	var body = &bytes.Buffer{}
	w := multipart.NewWriter(body)
	w.WriteField("chat_id", strconv.Itoa(id))
	var media []interface{}
	for _, i := range photos {
		switch v := i.(type) {
		case *os.File:
			file, _ := w.CreateFormFile(v.Name(), v.Name())
			_, err = io.Copy(file, v)
			_, err = v.Seek(0, io.SeekStart)
			media = append(media, inputMediaPhoto{Media: "attach://" + v.Name(), Type: "photo"})
		case string:
			media = append(media, inputMediaPhoto{Media: v, Type: "photo"})
		}
	}
	for _, i := range videos {
		switch v := i.(type) {
		case *os.File:
			file, _ := w.CreateFormFile(v.Name(), v.Name())
			_, err = io.Copy(file, v)
			_, err = v.Seek(0, io.SeekStart)
			media = append(media, inputMediaVideo{Media: "attach://" + v.Name(), Type: "video"})
		case string:
			media = append(media, inputMediaVideo{Media: v, Type: "video"})
		}
	}
	for _, i := range documents {
		switch v := i.(type) {
		case *os.File:
			file, _ := w.CreateFormFile(v.Name(), v.Name())
			_, err = io.Copy(file, v)
			_, err = v.Seek(0, io.SeekStart)
			media = append(media, inputMediaDocument{Media: "attach://" + v.Name(), Type: "document"})
		case string:
			media = append(media, inputMediaDocument{Media: v, Type: "documents"})
		}
	}
	for _, i := range audios {
		switch v := i.(type) {
		case *os.File:
			file, _ := w.CreateFormFile(v.Name(), v.Name())
			_, err = io.Copy(file, v)
			_, err = v.Seek(0, io.SeekStart)
			media = append(media, inputMediaAudio{Media: "attach://" + v.Name(), Type: "audio"})
		case string:
			media = append(media, inputMediaAudio{Media: v, Type: "audio"})
		}
	}
	if media == nil {
		return "", errors.New("you did not pass any file, file_id or URL")
	}
	mediaToJson, _ := json.Marshal(media)
	w.WriteField("media", string(mediaToJson))
	err = w.Close()
	req.Header.Add("Content-Type", w.FormDataContentType())
	req.Body = ioutil.NopCloser(bytes.NewReader(body.Bytes()))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	resToString, _ := ioutil.ReadAll(res.Body)
	return string(resToString), nil
}
