package gogram

import "errors"

type InlineQuery struct {
	Id       string   `json:"id"`
	From     User     `json:"from"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
	ChatType string   `json:"chatType"`
	Location Location `json:"location"`
}

type QueryAnswer interface {
	checkQueryAnswer() error
}

type MessageContent interface {
	checkMessageContent() error
}

type InputTextMessageContent struct {
	MessageText           string          `json:"messageText"`
	ParseMode             string          `json:"parseMode"`
	Entities              []MessageEntity `json:"Entities"`
	DisableWebPagePreview bool            `json:"disableWebPagePreview"`
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
	FoursquareId    string  `json:"foursquareId"`
	FoursquareType  string  `json:"foursquareType"`
	GooglePlaceId   string  `json:"googlePlaceId"`
	GooglePlaceType string  `json:"googlePlaceType"`
}

func (i InputVenueMessageContent) checkMessageContent() error {
	return nil
}

type InputContactMessageContent struct {
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Vcard       string `json:"vcard"`
}

func (i InputContactMessageContent) checkMessageContent() error {
	return nil
}

type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"providerToken"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int            `json:"maxTipAmount"`
	SuggestedTipAmounts       []int          `json:"suggestedTipAmounts"`
	ProviderData              string         `json:"providerData"`
	PhotoUrl                  string         `json:"photoUrl"`
	PhotoSize                 int            `json:"photoSize"`
	PhotoWidth                int            `json:"photoWidth"`
	PhotoHeight               int            `json:"photoHeight"`
	NeedName                  bool           `json:"needName"`
	NeedPhoneNumber           bool           `json:"needPhoneNumber"`
	NeedEmail                 bool           `json:"needEmail"`
	NeedShippingAddress       bool           `json:"needShippingAddress"`
	SendPhoneNumberToProvider bool           `json:"sendPhoneNumberToProvider"`
	SendEmailToProvider       bool           `json:"sendEmailToProvider"`
	IsFlexible                bool           `json:"isFlexible"`
}

func (i InputInvoiceMessageContent) checkMessageContent() error {
	return nil
}

type InlineQueryResultArticle struct {
	Type                string         `json:"type"`
	Id                  string         `json:"id"`
	Title               string         `json:"title"`
	InputMessageContent MessageContent `json:"inputMessageContent"`
	Url                 string         `json:"url"`
	HideUrl             bool           `json:"hideUrl"`
	Description         string         `json:"description"`
	ThumbUrl            string         `json:"thumbUrl"`
	ThumbWidth          int            `json:"thumbWidth"`
	ThumbHeight         int            `json:"thumbHeight"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultArticle) checkQueryAnswer() error {
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
	return nil
}

type InlineQueryResultPhoto struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	PhotoUrl            string          `json:"photoUrl"`
	ThumbUrl            string          `json:"thumbUrl"`
	PhotoWidth          int             `json:"photoWidth"`
	PhotoHeight         int             `json:"photoHeight"`
	Title               string          `json:"title"`
	Description         string          `json:"description"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultPhoto) checkQueryAnswer() error {
	if i.Type != "photo" {
		return errors.New("type must be `photo`")
	}
	if i.InputMessageContent == nil {
		return errors.New("InputMessageContent is required")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.PhotoUrl == "" {
		return errors.New("PhotoUrl is required")
	}
	if i.ThumbUrl == "" {
		return errors.New("ThumbUrl is required")
	}
	return nil
}

type InlineQueryResultGif struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	GifUrl              string          `json:"gifUrl"`
	GifWidth            int             `json:"gifWidth"`
	GifHeight           int             `json:"gifHeight"`
	GifDuration         int             `json:"gifDuration"`
	ThumbUrl            string          `json:"thumbUrl"`
	ThumbMimeType       string          `json:"thumbMimeType"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultGif) checkQueryAnswer() error {
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
	return nil
}

type InlineQueryResultMpeg4Gif struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Mpeg4Url            string          `json:"mpeg4Url"`
	Mpeg4Width          int             `json:"mpeg4Width"`
	Mpeg4Height         int             `json:"mpeg4Height"`
	Mpeg4Duration       int             `json:"mpeg4Duration"`
	ThumbUrl            string          `json:"thumbUrl"`
	ThumbMimeType       string          `json:"thumbMimeType"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultMpeg4Gif) checkQueryAnswer() error {
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
	return nil
}

type InlineQueryResultVideo struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	VideoUrl            string          `json:"videoUrl"`
	MimeType            string          `json:"mimeType"`
	ThumbUrl            string          `json:"thumbUrl"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	VideoWidth          int             `json:"videoWidth"`
	VideoHeight         int             `json:"videoHeight"`
	VideoDuration       int             `json:"videoDuration"`
	Description         string          `json:"description"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultVideo) checkQueryAnswer() error {
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
	return nil
}

type InlineQueryResultAudio struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	AudioUrl            string          `json:"audioUrl"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	Performer           string          `json:"performer"`
	AudioDuration       int             `json:"audioDuration"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultAudio) checkQueryAnswer() error {
	if i.Type != "audio" {
		return errors.New("type must be `audio`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.AudioUrl == "" {
		return errors.New("AudioUrl is required")
	}
	return nil
}

type InlineQueryResultVoice struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	VoiceUrl            string          `json:"voiceUrl"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	VoiceDuration       int             `json:"voiceDuration"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultVoice) checkQueryAnswer() error {
	if i.Type != "voice" {
		return errors.New("type must be `voice`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	if i.VoiceUrl == "" {
		return errors.New("VoiceUrl is required")
	}
	return nil
}

type InlineQueryResultDocument struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	DocumentUrl         string          `json:"documentUrl"`
	MimeType            string          `json:"mimeType"`
	Description         string          `json:"description"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	ThumbUrl            string          `json:"thumbUrl"`
	ThumbWidth          int             `json:"thumbWidth"`
	ThumbHeight         int             `json:"thumbHeight"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultDocument) checkQueryAnswer() error {
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
	return nil
}

type InlineQueryResultLocation struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Location
	Title               string         `json:"title"`
	InputMessageContent MessageContent `json:"inputMessageContent"`
	ThumbUrl            string         `json:"thumbUrl"`
	ThumbWidth          int            `json:"thumbWidth"`
	ThumbHeight         int            `json:"thumbHeight"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultLocation) checkQueryAnswer() error {
	if i.Type != "location" {
		return errors.New("type must be `location`")
	}
	if i.Id == "" {
		return errors.New("id is required")
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
	FoursquareId        string         `json:"foursquareId"`
	FoursquareType      string         `json:"foursquareType"`
	GooglePlaceId       string         `json:"googlePlaceId"`
	GooglePlaceType     string         `json:"googlePlaceType"`
	InputMessageContent MessageContent `json:"inputMessageContent"`
	ThumbUrl            string         `json:"thumbUrl"`
	ThumbWidth          int            `json:"thumbWidth"`
	ThumbHeight         int            `json:"thumbHeight"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultVenue) checkQueryAnswer() error {
	if i.Type != "venue" {
		return errors.New("type must be `venue`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	return nil
}

type InlineQueryResultContact struct {
	Type                string         `json:"type"`
	Id                  string         `json:"id"`
	PhoneNumber         string         `json:"phoneNumber"`
	FirstName           string         `json:"firstName"`
	LastName            string         `json:"lastName"`
	Vcard               string         `json:"vcard"`
	InputMessageContent MessageContent `json:"inputMessageContent"`
	ThumbUrl            string         `json:"thumbUrl"`
	ThumbWidth          int            `json:"thumbWidth"`
	ThumbHeight         int            `json:"thumbHeight"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultContact) checkQueryAnswer() error {
	if i.Type != "contact" {
		return errors.New("type must be `contact`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	return nil
}

type InlineQueryResultGame struct {
	Type          string `json:"type"`
	Id            string `json:"id"`
	GameShortName string `json:"gameShortName"`
	inlineKeyboardMarkup
}

func (i InlineQueryResultGame) checkQueryAnswer() error {
	if i.Type != "game" {
		return errors.New("type must be `game`")
	}
	if i.Id == "" {
		return errors.New("id is required")
	}
	return nil
}

type InlineQueryResultCachedPhoto struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	PhotoFileId         string          `json:"photoFileId"`
	Title               string          `json:"title"`
	Description         string          `json:"description"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedGif struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	GifFileId           string          `json:"gifFileId"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedMpeg4Gif struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Mpeg4FileId         string          `json:"mpeg4FileId"`
	Title               string          `json:"title"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedSticker struct {
	Type                string         `json:"type"`
	Id                  string         `json:"id"`
	StickerFileId       string         `json:"stickerFileId"`
	InputMessageContent MessageContent `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedDocument struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Title               string          `json:"title"`
	DocumentFileId      string          `json:"documentFileId"`
	Description         string          `json:"description"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedVideo struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Title               string          `json:"title"`
	VideoFileId         string          `json:"videoFileId"`
	Description         string          `json:"description"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedVoice struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	Title               string          `json:"title"`
	VoiceFileId         string          `json:"voiceFileId"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

type InlineQueryResultCachedAudio struct {
	Type                string          `json:"type"`
	Id                  string          `json:"id"`
	AudioFileId         string          `json:"audioFileId"`
	Caption             string          `json:"caption"`
	ParseMode           string          `json:"parseMode"`
	CaptionEntities     []MessageEntity `json:"captionEntities"`
	InputMessageContent MessageContent  `json:"inputMessageContent"`
	inlineKeyboardMarkup
}

func (i *InlineQuery) Answer(b Bot, results []QueryAnswer,
	optionalParams *AnswerInlineQueryOP) (response *BooleanResponse, err error) {
	type data struct {
		InlineQueryId string        `json:"inlineQueryId"`
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
