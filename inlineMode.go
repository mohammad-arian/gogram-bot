package gogram

import "errors"

type InlineQuery struct {
	Id       string   `json:"id"`
	From     User     `json:"from"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
	ChatType string   `json:"chat_type"`
	Location Location `json:"location"`
}

type QueryAnswer interface {
	checkQueryAnswer() error
}

type MessageContent interface {
	checkMessageContent() error
}

type InputEmptyContent struct {
}

func (i InputEmptyContent) checkMessageContent() error {
	return nil
}

type InputTextMessageContent struct {
	MessageText           string          `json:"message_text"`
	ParseMode             string          `json:"parse_mode"`
	Entities              []MessageEntity `json:"entities"`
	DisableWebPagePreview bool            `json:"disable_web_page_preview"`
}

func (i InputTextMessageContent) checkMessageContent() error {
	if len(i.MessageText) == 0 {
		return errors.New("you need to set Id of InlineQueryResultPhoto to a string")
	}
	return nil
}

type InputLocationMessageContent struct {
	Location
}

func (i InputLocationMessageContent) checkMessageContent() error {
	return nil
}

type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareId    string  `json:"foursquare_id"`
	FoursquareType  string  `json:"foursquare_type"`
	GooglePlaceId   string  `json:"google_place_id"`
	GooglePlaceType string  `json:"google_place_type"`
}

func (i InputVenueMessageContent) checkMessageContent() error {
	if i.Title == "" {
		return errors.New("you need to set Title of InputVenueMessageContent to a string")
	}
	if i.Address == "" {
		return errors.New("you need to set Address of InputVenueMessageContent to a string")
	}
	return nil
}

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Vcard       string `json:"vcard"`
}

func (i InputContactMessageContent) checkMessageContent() error {
	if i.PhoneNumber == "" {
		return errors.New("you need to set PhoneNumber of InputVenueMessageContent to a string")
	}
	if i.FirstName == "" {
		return errors.New("you need to set FirstName of InputVenueMessageContent to a string")
	}
	return nil
}

type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int            `json:"max_tip_amount"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts"`
	ProviderData              string         `json:"provider_data"`
	PhotoUrl                  string         `json:"photo_url"`
	PhotoSize                 int            `json:"photo_size"`
	PhotoWidth                int            `json:"photo_width"`
	PhotoHeight               int            `json:"photo_height"`
	NeedName                  bool           `json:"need_name"`
	NeedPhoneNumber           bool           `json:"need_phone_number"`
	NeedEmail                 bool           `json:"need_email"`
	NeedShippingAddress       bool           `json:"need_shipping_address"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider"`
	SendEmailToProvider       bool           `json:"send_email_to_provider"`
	IsFlexible                bool           `json:"is_flexible"`
}

func (i InputInvoiceMessageContent) checkMessageContent() error {
	if i.Title == "" {
		return errors.New("you need to set Title of InputInvoiceMessageContent to a string")
	}
	if i.Description == "" {
		return errors.New("you need to set Description of InputInvoiceMessageContent to a string")
	}
	if i.Payload == "" {
		return errors.New("you need to set Payload of InputInvoiceMessageContent to a string")
	}
	if i.ProviderData == "" {
		return errors.New("you need to set ProviderData of InputInvoiceMessageContent to a string")
	}
	if i.Currency == "" {
		return errors.New("you need to set Currency of InputInvoiceMessageContent to a string")
	}
	if len(i.Prices) == 0 {
		return errors.New("you need to set Prices of InputInvoiceMessageContent to a slice of " +
			"LabeledPrice struct")
	}
	return nil
}

type InlineQueryResultArticle struct {
	Type                string         `json:"type"`
	Id                  string         `json:"id"`
	Title               string         `json:"title"`
	InputMessageContent MessageContent `json:"input_message_content"`
	Url                 string         `json:"url"`
	HideUrl             bool           `json:"hide_url"`
	Description         string         `json:"description"`
	ThumbUrl            string         `json:"thumb_url"`
	ThumbWidth          int            `json:"thumb_width"`
	ThumbHeight         int            `json:"thumb_height"`
	InlineKeyboard
}

func (i *InlineQueryResultArticle) checkQueryAnswer() error {
	i.Type = "article"
	if i.InputMessageContent == nil {
		return errors.New("you need to set InputMessageContent of InlineQueryResultArticle to a MessageContent" +
			" such as InputTextMessageContent, InputLocationMessageContent etc")
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultArticle to a unique string")
	}
	if i.Title == "" {
		return errors.New("you need to set Title of InlineQueryResultArticle to a string")
	}
	return nil
}

type InlineQueryResultPhoto struct {
	Id                  string          `json:"id"`
	PhotoUrl            string          `json:"photo_url"`
	PhotoFileId         string          `json:"photo_file_id"`
	ThumbUrl            string          `json:"thumb_url"`
	Type                string          `json:"type"`
	PhotoWidth          int             `json:"photo_width"`
	PhotoHeight         int             `json:"photo_height"`
	Title               string          `json:"title"`
	Description         string          `json:"description"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	InlineKeyboard
}

func (i *InlineQueryResultPhoto) checkQueryAnswer() error {
	i.Type = "photo"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultPhoto to a unique string")
	}
	if i.PhotoUrl == "" && i.PhotoFileId == "" {
		return errors.New("you need to set PhotoUrl or PhotoFileId of InlineQueryResultPhoto to a photo url or " +
			"file id on telegram server")
	}
	if i.PhotoUrl != "" && i.PhotoFileId != "" {
		return errors.New("set PhotoUrl or PhotoFileId of InlineQueryResultPhoto, not both")
	}
	if i.ThumbUrl == "" && i.PhotoUrl != "" {
		return errors.New("ThumbUrl is required if you are setting PhotoUrl of " +
			"InlineQueryResultPhoto to a url")
	}
	if i.InputMessageContent == nil {
		i.InputMessageContent = InputEmptyContent{}
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	return nil
}

type InlineQueryResultGif struct {
	Id                  string          `json:"id"`
	GifUrl              string          `json:"gif_url"`
	GifFileId           string          `json:"gif_File_Id"`
	Type                string          `json:"type"`
	GifWidth            int             `json:"gif_width"`
	GifHeight           int             `json:"gif_height"`
	GifDuration         int             `json:"gif_duration"`
	ThumbUrl            string          `json:"thumb_url"`
	ThumbMimeType       string          `json:"thumb_mime_type"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	InlineKeyboard
}

func (i *InlineQueryResultGif) checkQueryAnswer() error {
	i.Type = "gif"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultGif to a unique string")
	}
	if i.GifUrl == "" && i.GifFileId == "" {
		return errors.New("you need to set GifUrl or GifFileId of InlineQueryResultGif to a gif url or " +
			"file id on telegram server")
	}
	if i.GifUrl != "" && i.GifFileId != "" {
		return errors.New("set GifUrl or GifFileId of InlineQueryResultGif, not both")
	}
	if i.ThumbUrl == "" && i.GifUrl != "" {
		return errors.New("ThumbUrl is required if you are setting GifUrl of InlineQueryResultGif to a url")
	}
	if i.InputMessageContent == nil {
		i.InputMessageContent = InputEmptyContent{}
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	return nil
}

type InlineQueryResultMpeg4Gif struct {
	Id                  string          `json:"id"`
	Mpeg4Url            string          `json:"mpeg4_url"`
	Mpeg4FileId         string          `json:"mpeg4_file_id"`
	Type                string          `json:"type"`
	Mpeg4Width          int             `json:"mpeg4_width"`
	Mpeg4Height         int             `json:"mpeg4_height"`
	Mpeg4Duration       int             `json:"mpeg4_duration"`
	ThumbUrl            string          `json:"thumb_url"`
	ThumbMimeType       string          `json:"thumb_mime_type"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	InlineKeyboard
}

func (i *InlineQueryResultMpeg4Gif) checkQueryAnswer() error {
	i.Type = "mpeg4_gif"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultMpeg4Gif to a unique string")
	}
	if i.Mpeg4Url == "" && i.Mpeg4FileId == "" {
		return errors.New("you need to set Mpeg4Url or Mpeg4FileId of InlineQueryResultMpeg4Gif to a " +
			"video animation (H.264/MPEG-4 AVC video without sound) url or file id on telegram server")
	}
	if i.Mpeg4Url != "" && i.Mpeg4FileId != "" {
		return errors.New("set Mpeg4Url or Mpeg4FileId of InlineQueryResultMpeg4Gif, not both")
	}
	if i.ThumbUrl == "" && i.Mpeg4Url != "" {
		return errors.New("ThumbUrl is required if you are setting Mpeg4Url of " +
			"InlineQueryResultMpeg4Gif to a url")
	}
	if i.InputMessageContent == nil {
		i.InputMessageContent = InputEmptyContent{}
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	return nil
}

type InlineQueryResultVideo struct {
	Id                  string          `json:"id"`
	VideoUrl            string          `json:"video_url"`
	VideoFileId         string          `json:"video_file_id"`
	Type                string          `json:"type"`
	MimeType            string          `json:"mime_type"`
	ThumbUrl            string          `json:"thumb_url"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	VideoWidth          int             `json:"video_width"`
	VideoHeight         int             `json:"video_height"`
	VideoDuration       int             `json:"video_duration"`
	Description         string          `json:"description"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	InlineKeyboard
}

func (i *InlineQueryResultVideo) checkQueryAnswer() error {
	i.Type = "video"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultVideo to a unique string")
	}
	if i.VideoUrl == "" && i.VideoFileId == "" {
		return errors.New("you need to set VideoUrl or VideoFileId of InlineQueryResultVideo to a " +
			"video url or file id on telegram server")
	}
	if i.VideoUrl != "" && i.VideoFileId != "" {
		return errors.New("set VideoUrl or VideoFileId of InlineQueryResultVideo, not both")
	}
	if i.ThumbUrl == "" && i.VideoFileId != "" {
		return errors.New("ThumbUrl is required if you are setting VideoUrl of " +
			"InlineQueryResultVideo to a url")
	}
	if i.Title == "" {
		return errors.New("you need to set Title of InlineQueryResultVideo to a string")
	}
	if i.MimeType == "" && i.VideoUrl != "" {
		return errors.New("MimeType is required if you are setting VideoUrl of InlineQueryResultVideo to a url. " +
			"you need to set MimeType to Mime type of the content of video url, “text/html” or “video/mp4”")
	}
	if i.InputMessageContent == nil {
		i.InputMessageContent = InputEmptyContent{}
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	return nil
}

type InlineQueryResultAudio struct {
	Id                  string          `json:"id"`
	AudioUrl            string          `json:"audio_url"`
	AudioFileId         string          `json:"audio_file_id"`
	Type                string          `json:"type"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	Performer           string          `json:"performer"`
	AudioDuration       int             `json:"audio_duration"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	InlineKeyboard
}

func (i *InlineQueryResultAudio) checkQueryAnswer() error {
	i.Type = "audio"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultAudio to a unique string")
	}
	if i.AudioUrl == "" && i.AudioFileId == "" {
		return errors.New("you need to set AudioUrl or AudioFileId of InlineQueryResultAudio to a " +
			"audio url or file id on telegram server")
	}
	if i.AudioUrl != "" && i.AudioFileId != "" {
		return errors.New("set AudioUrl or AudioFileId of InlineQueryResultAudio, not both")
	}
	if i.Title == "" {
		return errors.New("you need to set Title of InlineQueryResultAudio to a string")
	}
	if i.InputMessageContent == nil {
		i.InputMessageContent = InputEmptyContent{}
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	return nil
}

type InlineQueryResultVoice struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	VoiceUrl            string          `json:"voice_url"`
	VoiceFileId         string          `json:"voice_file_id"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	VoiceDuration       int             `json:"voice_duration"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	InlineKeyboard
}

func (i *InlineQueryResultVoice) checkQueryAnswer() error {
	i.Type = "voice"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultVoice to a unique string")
	}
	if i.VoiceUrl == "" && i.VoiceFileId == "" {
		return errors.New("you need to set VoiceUrl or VoiceFileId of InlineQueryResultVoice to a " +
			"audio url or file id on telegram server")
	}
	if i.VoiceUrl != "" && i.VoiceFileId != "" {
		return errors.New("set VoiceUrl or VoiceFileId of InlineQueryResultVoice, not both")
	}
	if i.Title == "" {
		return errors.New("you need to set Title of InlineQueryResultVoice to a string")
	}
	if i.InputMessageContent == nil {
		i.InputMessageContent = InputEmptyContent{}
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	return nil
}

type InlineQueryResultDocument struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	DocumentUrl         string          `json:"document_url"`
	DocumentFileId      string          `json:"document_file_id"`
	MimeType            string          `json:"mime_type"`
	Description         string          `json:"description"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	ThumbUrl            string          `json:"thumb_url"`
	ThumbWidth          int             `json:"thumb_width"`
	ThumbHeight         int             `json:"thumb_height"`
	InlineKeyboard
}

func (i *InlineQueryResultDocument) checkQueryAnswer() error {
	i.Type = "document"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultDocument to a unique string")
	}
	if i.DocumentUrl == "" && i.DocumentFileId == "" {
		return errors.New("you need to set DocumentUrl or DocumentFileId of InlineQueryResultDocument to a " +
			"audio url or file id on telegram server")
	}
	if i.DocumentUrl != "" && i.DocumentFileId != "" {
		return errors.New("set DocumentUrl or DocumentFileId of InlineQueryResultDocument, not both")
	}
	if i.MimeType == "" && i.DocumentUrl != "" {
		return errors.New("MimeType is required if you are setting DocumentUrl of InlineQueryResultDocument to " +
			"a url. you need to set MimeType to Mime type of the content of " +
			"the file, either “application/pdf” or “application/zip”")
	}
	if i.Title == "" {
		return errors.New("you need to set Title of InlineQueryResultDocument to a string")
	}
	if i.InputMessageContent == nil {
		i.InputMessageContent = InputEmptyContent{}
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	return nil
}

type InlineQueryResultLocation struct {
	Type                string         `json:"type"`
	Id                  string         `json:"id"`
	Title               string         `json:"title"`
	InputMessageContent MessageContent `json:"input_message_content"`
	ThumbUrl            string         `json:"thumb_url"`
	ThumbWidth          int            `json:"thumb_width"`
	ThumbHeight         int            `json:"thumb_height"`
	Location
	InlineKeyboard
}

func (i *InlineQueryResultLocation) checkQueryAnswer() error {
	i.Type = "location"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultLocation to a unique string")
	}
	if i.Title == "" {
		return errors.New("you need to set Title of InlineQueryResultLocation to a string")
	}
	if i.InputMessageContent == nil {
		i.InputMessageContent = InputEmptyContent{}
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	return nil
}

type InlineQueryResultVenue struct {
	Type                string         `json:"type"`
	Id                  string         `json:"id"`
	Latitude            float64        `json:"latitude"`
	Longitude           float64        `json:"longitude"`
	Title               string         `json:"title"`
	Address             string         `json:"address"`
	FoursquareId        string         `json:"foursquare_id"`
	FoursquareType      string         `json:"foursquare_type"`
	GooglePlaceId       string         `json:"google_place_id"`
	GooglePlaceType     string         `json:"google_place_type"`
	InputMessageContent MessageContent `json:"input_message_content"`
	ThumbUrl            string         `json:"thumb_url"`
	ThumbWidth          int            `json:"thumb_width"`
	ThumbHeight         int            `json:"thumb_height"`
	InlineKeyboard
}

func (i *InlineQueryResultVenue) checkQueryAnswer() error {
	i.Type = "venue"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultVenue to a unique string")
	}
	if i.Address == "" {
		return errors.New("you need to set Address of InlineQueryResultVenue to a string")
	}
	if i.Title == "" {
		return errors.New("you need to set Title of InlineQueryResultVenue to a string")
	}
	if i.InputMessageContent == nil {
		i.InputMessageContent = InputEmptyContent{}
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	return nil
}

type InlineQueryResultContact struct {
	Type                string         `json:"type"`
	Id                  string         `json:"id"`
	PhoneNumber         string         `json:"phone_number"`
	FirstName           string         `json:"first_name"`
	LastName            string         `json:"last_name"`
	Vcard               string         `json:"vcard"`
	InputMessageContent MessageContent `json:"input_message_content"`
	ThumbUrl            string         `json:"thumb_url"`
	ThumbWidth          int            `json:"thumb_width"`
	ThumbHeight         int            `json:"thumb_height"`
	InlineKeyboard
}

func (i *InlineQueryResultContact) checkQueryAnswer() error {
	i.Type = "contact"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultContact to a unique string")
	}
	if i.FirstName == "" {
		return errors.New("you need to set FirstName of InlineQueryResultContact to a string")
	}
	if i.InputMessageContent == nil {
		i.InputMessageContent = InputEmptyContent{}
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	return nil
}

type InlineQueryResultGame struct {
	Type          string `json:"type"`
	Id            string `json:"id"`
	GameShortName string `json:"game_short_name"`
	InlineKeyboard
}

func (i *InlineQueryResultGame) checkQueryAnswer() error {
	i.Type = "game"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultGame to a unique string")
	}
	if i.GameShortName == "" {
		return errors.New("you need to set GameShortName of InlineQueryResultGame to a string")
	}
	return nil
}

type InlineQueryResultSticker struct {
	Type                string         `json:"type"`
	Id                  string         `json:"id"`
	StickerFileId       string         `json:"sticker_file_id"`
	InputMessageContent MessageContent `json:"input_message_content"`
	InlineKeyboard
}

func (i *InlineQueryResultSticker) checkQueryAnswer() error {
	i.Type = "sticker"
	if i.Id == "" {
		return errors.New("you need to set Id of InlineQueryResultSticker to a unique string")
	}
	if i.StickerFileId == "" {
		return errors.New("you need to set GameShortName of InlineQueryResultGame to a file id")
	}
	return nil
}

func (i *InlineQuery) Answer(b Bot, data AnswerInlineQueryData) (response *BooleanResponse, err error) {
	data.InlineQueryId = i.Id
	return data.Send(b)
}
