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
		return errors.New("MessageText field is mandatory")
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
	return nil
}

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Vcard       string `json:"vcard"`
}

func (i InputContactMessageContent) checkMessageContent() error {
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
	inlineKeyboardMarkup
}

func (i *InlineQueryResultArticle) checkQueryAnswer() error {
	if i.Type != "article" {
		return errors.New("type must be `article`")
	}
	if i.InputMessageContent == nil {
		return errors.New("InputMessageContent is required")
	} else {
		e := i.InputMessageContent.checkMessageContent()
		if e != nil {
			return e
		}
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.Title == "" {
		return errors.New("title is required")
	}
	return nil
}

type InlineQueryResultPhoto struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	PhotoUrl            string          `json:"photo_url"`
	PhotoId             string          `json:"photo_id"`
	ThumbUrl            string          `json:"thumb_url"`
	PhotoWidth          int             `json:"photo_width"`
	PhotoHeight         int             `json:"photo_height"`
	Title               string          `json:"title"`
	Description         string          `json:"description"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	inlineKeyboardMarkup
}

func (i *InlineQueryResultPhoto) checkQueryAnswer() error {
	if i.Type != "photo" {
		return errors.New("type must be `photo`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.PhotoUrl == "" || i.PhotoId == "" {
		return errors.New("photo_url or photo_id is required")
	}
	if i.ThumbUrl == "" {
		return errors.New("ThumbUrl is required")
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
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	GifUrl              string          `json:"gif_url"`
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
	inlineKeyboardMarkup
}

func (i *InlineQueryResultGif) checkQueryAnswer() error {
	if i.Type != "gif" {
		return errors.New("type must be `gif`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.GifUrl == "" {
		return errors.New("GifUrl is required")
	}
	if i.ThumbUrl == "" {
		return errors.New("ThumbUrl is required")
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
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Mpeg4Url            string          `json:"mpeg4_url"`
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
	inlineKeyboardMarkup
}

func (i *InlineQueryResultMpeg4Gif) checkQueryAnswer() error {
	if i.Type != "mpeg4_gif" {
		return errors.New("type must be `mpeg4_gif`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.Mpeg4Url == "" {
		return errors.New("Mpeg4Url is required")
	}
	if i.ThumbUrl == "" {
		return errors.New("ThumbUrl is required")
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
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	VideoUrl            string          `json:"video_url"`
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
	inlineKeyboardMarkup
}

func (i *InlineQueryResultVideo) checkQueryAnswer() error {
	if i.Type != "video" {
		return errors.New("type must be `video`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.VideoUrl == "" {
		return errors.New("PhotoUrl is required")
	}
	if i.MimeType == "" {
		return errors.New("MimeType is required")
	}
	if i.ThumbUrl == "" {
		return errors.New("ThumbUrl is required")
	}
	if i.Title == "" {
		return errors.New("title is required")
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
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	AudioUrl            string          `json:"audio_url"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	Performer           string          `json:"performer"`
	AudioDuration       int             `json:"audio_duration"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	inlineKeyboardMarkup
}

func (i *InlineQueryResultAudio) checkQueryAnswer() error {
	if i.Type != "audio" {
		return errors.New("type must be `audio`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.AudioUrl == "" {
		return errors.New("AudioUrl is required")
	}
	if i.Title == "" {
		return errors.New("title is required")
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
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	VoiceDuration       int             `json:"voice_duration"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	inlineKeyboardMarkup
}

func (i *InlineQueryResultVoice) checkQueryAnswer() error {
	if i.Type != "voice" {
		return errors.New("type must be `voice`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.VoiceUrl == "" {
		return errors.New("VoiceUrl is required")
	}
	if i.Title == "" {
		return errors.New("title is required")
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
	MimeType            string          `json:"mime_type"`
	Description         string          `json:"description"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	ThumbUrl            string          `json:"thumb_url"`
	ThumbWidth          int             `json:"thumb_width"`
	ThumbHeight         int             `json:"thumb_height"`
	inlineKeyboardMarkup
}

func (i *InlineQueryResultDocument) checkQueryAnswer() error {
	if i.Type != "document" {
		return errors.New("type must be `document`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.DocumentUrl == "" {
		return errors.New("AudioUrl is required")
	}
	if i.MimeType == "" {
		return errors.New("MimeType is required")
	}
	if i.Title == "" {
		return errors.New("title is required")
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
	inlineKeyboardMarkup
}

func (i *InlineQueryResultLocation) checkQueryAnswer() error {
	if i.Type != "location" {
		return errors.New("type must be `location`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.Title == "" {
		return errors.New("title is required")
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
	inlineKeyboardMarkup
}

func (i *InlineQueryResultVenue) checkQueryAnswer() error {
	if i.Type != "venue" {
		return errors.New("type must be `venue`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.Address == "" {
		return errors.New("address is required")
	}
	if i.Title == "" {
		return errors.New("title is required")
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
	inlineKeyboardMarkup
}

func (i *InlineQueryResultContact) checkQueryAnswer() error {
	if i.Type != "contact" {
		return errors.New("type must be `contact`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.FirstName == "" {
		return errors.New("first_name is required")
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
	inlineKeyboardMarkup
}

func (i *InlineQueryResultGame) checkQueryAnswer() error {
	if i.Type != "game" {
		return errors.New("type must be `game`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.GameShortName == "" {
		return errors.New("game_short_name is required")
	}
	return nil
}

type InlineQueryResultCachedPhoto struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	PhotoFileId         string          `json:"photo_file_id"`
	Title               string          `json:"title"`
	Description         string          `json:"description"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedGif struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	GifFileId           string          `json:"gif_file_id"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedMpeg4Gif struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Mpeg4FileId         string          `json:"mpeg4_file_id"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedSticker struct {
	Type                string         `json:"type"`
	Id                  string         `json:"id"`
	StickerFileId       string         `json:"sticker_file_id"`
	InputMessageContent MessageContent `json:"input_message_content"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedDocument struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Title               string          `json:"title"`
	DocumentFileId      string          `json:"document_file_id"`
	Description         string          `json:"description"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedVideo struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Title               string          `json:"title"`
	VideoFileId         string          `json:"video_file_id"`
	Description         string          `json:"description"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedVoice struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Title               string          `json:"title"`
	VoiceFileId         string          `json:"voice_file_id"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedAudio struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	AudioFileId         string          `json:"audio_file_id"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parse_mode"`
	CaptionEntities     []MessageEntity `json:"caption_entities"`
	InputMessageContent MessageContent  `json:"input_message_content"`
	inlineKeyboardMarkup
}

func (i *InlineQuery) Answer(b Bot, results []QueryAnswer,
	optionalParams *AnswerInlineQueryOP) (response *BooleanResponse, err error) {
	type data struct {
		InlineQueryId string        `json:"inline_query_id"`
		Results       []QueryAnswer `json:"results"`
	}
	if len(results) == 0 {
		return &BooleanResponse{}, errors.New("results slice is empty. pass QueryAnswer structs such as " +
			"InlineQueryResultArticle, InlineQueryResultPhoto and etc")
	}
	for _, j := range results {
		e := j.checkQueryAnswer()
		if e != nil {
			return &BooleanResponse{}, e
		}
	}
	d := data{i.Id, results}
	res, err := request("answerInlineQuery", b.Token, &d, optionalParams, &BooleanResponse{})
	return res.(*BooleanResponse), err
}
