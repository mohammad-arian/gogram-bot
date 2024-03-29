package gogram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Update from webhook
type Update struct {
	UpdateId      int           `json:"update_id"`
	Message       Message       `json:"message"`
	InlineQuery   InlineQuery   `json:"inline_query"`
	CallbackQuery CallbackQuery `json:"callback_query"`
	Poll          Poll          `json:"poll"`
}

func (u Update) String() string {
	return fmt.Sprintf("Update: %#v\n", u)
}

type Message struct {
	MessageId             int               `json:"message_id"`
	User                  User              `json:"from"`
	Chat                  Chat              `json:"chat"`
	SenderChat            Chat              `json:"sender_chat"`
	ForwardFrom           User              `json:"forward_from"`
	ForwardFromChat       Chat              `json:"forward_from_chat"`
	ForwardSignature      string            `json:"forward_signature"`
	ForwardSenderName     string            `json:"forward_sender_name"`
	ForwardDate           int               `json:"forward_date"`
	IsAutomaticForward    bool              `json:"is_automatic_forward"`
	ReplyToMessage        *Message          `json:"reply_to_message"`
	ViaBot                User              `json:"via_bot"`
	EditDate              int               `json:"edit_date"`
	HasProtectedContent   bool              `json:"has_protected_content"`
	MediaGroupId          string            `json:"media_group_id"`
	AuthorSignature       string            `json:"author_signature"`
	Text                  string            `json:"text"`
	Entities              []MessageEntity   `json:"entities"`
	Animation             Animation         `json:"animation"`
	Photo                 []PhotoSize       `json:"photo"`
	Audio                 Audio             `json:"audio"`
	Document              Document          `json:"document"`
	Sticker               Sticker           `json:"sticker"`
	Video                 Video             `json:"video"`
	VideoNote             VideoNote         `json:"video_note"`
	Voice                 Voice             `json:"voice"`
	Caption               string            `json:"caption"`
	CaptionEntities       []MessageEntity   `json:"caption_entities"`
	Contact               Contact           `json:"contact"`
	Dice                  Dice              `json:"dice"`
	Game                  Game              `json:"game"`
	Date                  int               `json:"date"`
	ReplyMarkup           InlineKeyboard    `json:"reply_markup"`
	Poll                  Poll              `json:"poll"`
	Venue                 Venue             `json:"venue"`
	Location              Location          `json:"location"`
	LeftChatMember        User              `json:"left_chat_member"`
	NewChatPhoto          []PhotoSize       `json:"new_chat_photo"`
	NewChatTitle          string            `json:"new_chat_title"`
	NewChatMembers        []User            `json:"new_chat_members"`
	DeleteChatPhoto       bool              `json:"delete_chat_photo"`
	GroupChatCreated      bool              `json:"group_chat_created"`
	SupergroupChatCreated bool              `json:"supergroup_chat_created"`
	ChannelChatCreated    bool              `json:"channel_chat_created"`
	MigrateToChatId       int               `json:"migrate_to_chat_id"`
	MigrateFromChatId     int               `json:"migrate_from_chat_id"`
	PinnedMessage         *Message          `json:"pinned_message"`
	PassportData          PasswordData      `json:"passport_data"`
	Invoice               Invoice           `json:"invoice"`
	SuccessfulPayment     SuccessfulPayment `json:"successful_payment"`
	ConnectedWebsite      string            `json:"connected_website"`
}

const (
	TypeText              = "Text"
	TypePhoto             = "Photo"
	TypeAnimation         = "Animation"
	TypeForwardFrom       = "ForwardFrom"
	TypeReply             = "Reply"
	TypeAudio             = "Audio"
	TypeDocument          = "Document"
	TypeSticker           = "Sticker"
	TypeVideo             = "Video"
	TypeVideoNote         = "VideoNote"
	TypeVoice             = "Voice"
	TypeContact           = "Contact"
	TypeDice              = "Dice"
	TypeGame              = "Game"
	TypePoll              = "Poll"
	TypeVenue             = "Venue"
	TypeLocation          = "Location"
	TypeMemberLeftChat    = "MemberLeftChat"
	TypeNewChatTitle      = "NewChatTitle"
	TypeNewChatPhoto      = "NewChatPhoto"
	TypeDeleteChatPhoto   = "DeleteChatPhoto"
	TypeGroupCreated      = "GroupCreated"
	TypeSuperGroupCreated = "SuperGroupCreated"
	TypeChannelCreated    = "ChannelCreated"
	TypeMigrateToChatId   = "MigrateToChatId"
	TypeMigrateFromChatId = "MigrateFromChatId"
	TypePinnedMessage     = "PinnedMessage"
	TypeInvoice           = "Invoice"
	TypeSuccessfulPayment = "SuccessfulPayment"
	TypePassport          = "Passport"
	TypeUnknown           = "Unknown"
)

// TypeIndicator function returns the type of message.
func (m Message) TypeIndicator() string {
	switch {
	case m.Text != "":
		return TypeText
	case m.Animation != Animation{}:
		return TypeAnimation
	case m.Photo != nil:
		return TypePhoto
	case m.DeleteChatPhoto == true:
		return TypeDeleteChatPhoto
	case m.NewChatPhoto != nil:
		return TypeNewChatPhoto
	case m.ForwardFrom != User{}:
		return TypeForwardFrom
	case m.ReplyToMessage != nil:
		return TypeReply
	case m.Audio != Audio{}:
		return TypeAudio
	case m.Sticker != Sticker{}:
		return TypeSticker
	case m.Document != Document{}:
		return TypeDocument
	case m.Location != Location{}:
		return TypeLocation
	case m.Video != Video{}:
		return TypeVideo
	case m.VideoNote != VideoNote{}:
		return TypeVideoNote
	case m.Voice != Voice{}:
		return TypeVoice
	case m.Contact != Contact{}:
		return TypeContact
	case m.Dice != Dice{}:
		return TypeDice
	case m.Game.Title != "":
		return TypeGame
	case m.Poll.Id != "":
		return TypePoll
	case m.Venue != Venue{}:
		return TypeVenue
	case m.LeftChatMember != User{}:
		return TypeMemberLeftChat
	case m.NewChatTitle != "":
		return TypeNewChatTitle
	case m.GroupChatCreated == true:
		return TypeGroupCreated
	case m.SupergroupChatCreated == true:
		return TypeSuperGroupCreated
	case m.ChannelChatCreated == true:
		return TypeChannelCreated
	case m.MigrateToChatId != 0:
		return TypeMigrateToChatId
	case m.MigrateFromChatId != 0:
		return TypeMigrateFromChatId
	case m.PinnedMessage != nil:
		return TypePinnedMessage
	case m.Invoice != Invoice{}:
		return TypeInvoice
	case m.SuccessfulPayment != SuccessfulPayment{}:
		return TypeSuccessfulPayment
	case len(m.PassportData.data) != 0:
		return TypePassport
	default:
		return TypeUnknown
	}
}

type CallbackQuery struct {
	Id              string  `json:"id"`
	Message         Message `json:"message"`
	From            User    `json:"chat"`
	InlineMessageId string  `json:"inline_message_id"`
	ChatInstance    string  `json:"chat_instance"`
	Data            string  `json:"data"`
	GameShortName   string  `json:"game_short_name"`
}

type ShippingOptions struct {
	Id     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

type Audio struct {
	FileId       string    `json:"file_Id"`
	FileUniqueId string    `json:"file_unique_id"`
	Duration     int       `json:"duration"`
	Performer    string    `json:"performer"`
	Title        string    `json:"title"`
	FileName     string    `json:"file_name"`
	MimeType     string    `json:"mime_type"`
	FileSize     int       `json:"file_size"`
	Thumb        PhotoSize `json:"thumb"`
}

type Document struct {
	FileId       string    `json:"file_Id"`
	FileUniqueId string    `json:"file_unique_id"`
	Thumb        PhotoSize `json:"thumb"`
	FileName     string    `json:"file_name"`
	MimeType     string    `json:"mime_type"`
	FileSize     int       `json:"file_size"`
}

type Video struct {
	FileId       string    `json:"file_Id"`
	FileUniqueId string    `json:"file_unique_id"`
	Duration     int       `json:"duration"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Thumb        PhotoSize `json:"thumb"`
	FileName     string    `json:"file_name"`
	MimeType     string    `json:"mime_type"`
	FileSize     int       `json:"file_size"`
}

type VideoNote struct {
	FileId       string    `json:"file_Id"`
	FileUniqueId string    `json:"file_unique_id"`
	Length       int       `json:"length"`
	Duration     int       `json:"duration"`
	Thumb        PhotoSize `json:"thumb"`
	FileSize     int       `json:"file_size"`
}

type Voice struct {
	FileId       string `json:"file_Id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type"`
	FileSize     int    `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      int    `json:"user_id"`
	// Additional data about the contact in the form of a vCard
	Vcard string `json:"vcard"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    Animation       `json:"animation"`
}

type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type OrderInfo struct {
	Name            string          `json:"name"`
	PhoneNumber     string          `json:"phone_number"`
	Email           string          `json:"email"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type SuccessfulPayment struct {
	Currency                string    `json:"currency"`
	TotalAmount             int       `json:"total_amount"`
	InvoicePayload          string    `json:"invoice_payload"`
	ShippingOptionId        string    `json:"shipping_option_id"`
	OrderInfo               OrderInfo `json:"order_info"`
	TelegramPaymentChargeId string    `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string    `json:"provider_payment_charge_id"`
}

type MessageEntity struct {
	Type     string
	offset   int
	length   int
	url      string
	user     User
	language string
}

type Poll struct {
	Id              string       `json:"id"`
	Question        string       `json:"question"`
	Options         []PollOption `json:"options"`
	TotalVoterCount int          `json:"total_voter_count"`
	IsAnonymous     bool         `json:"is_anonymous,"`
	CloseDate       int          `json:"close_date"`
	// Pass True, if the poll needs to be immediately closed.
	// This can be useful for poll preview.
	IsClosed                 bool            `json:"is_closed"`
	Type                     string          `json:"type"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	CorrectOptionId          int             `json:"correct_option_id"`
	Explanation              string          `json:"explanation"`
	ExplanationEntities      []MessageEntity `json:"explanation_entities"`
	// Amount of time in seconds the poll will be active after
	// creation, 5-600. Can't be used together with close_date.
	OpenPeriod int `json:"open_period"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

type File struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// file size in bytes, if known. Optional
	FileSize int `json:"file_size"`
	// file path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file. Optional
	FilePath string `json:"file_path"`
}

type ReplyAble struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// User id is a unique identification number of a particular Telegram user.
// However, the Telegram Chat id is a unique identification
// number of a Telegram chat (personal or group chat).
// Use Chat id for groups, and User id for a specific user
type User struct {
	// SupportsInlineQueries shows if Bot supports inline queries
	// This field is only for bots
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
	LanguageCode            string `json:"language_code"`
	IsBot                   bool   `json:"is_bot"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	ReplyAble
}

// Chat id is a unique identification number of a Telegram chat (personal or group chat).
// However, the Telegram User id is a unique identification number of a particular Telegram user.
// Use Chat id for groups, and User id for a specific user
type Chat struct {
	Type                  string          `json:"type"`
	Title                 string          `json:"title"`
	Bio                   string          `json:"bio"`
	Photo                 ChatPhoto       `json:"photo"`
	Location              ChatLocation    `json:"location"`
	Description           string          `json:"description"`
	InviteLink            string          `json:"invite_link"`
	PinnedMessage         *Message        `json:"pinned_message"`
	Permissions           ChatPermissions `json:"permissions"`
	SlowModeDelay         int             `json:"slow_mode_delay"`
	MessageAutoDeleteTime int             `json:"message_auto_delete_time"`
	StickerSetName        string          `json:"sticker_set_name"`
	CanSetStickerSet      bool            `json:"can_set_sticker_set"`
	LinkedChatId          int             `json:"linked_chat_id"`
	ReplyAble
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

type ChatMemberOwner struct {
	Status      string `json:"status"`
	User        User   `json:"user"`
	IsAnonymous bool   `json:"is_anonymous"`
	CustomTitle string `json:"custom_title"`
}

type ChatMemberAdministrator struct {
	Status              string `json:"status"`
	User                User   `json:"user"`
	IsAnonymous         bool   `json:"is_anonymous"`
	CustomTitle         string `json:"custom_title"`
	CanBeEdited         bool   `json:"can_be_edited"`
	CanManageChat       bool   `json:"can_manage_chat"`
	CanDeleteMessages   bool   `json:"can_delete_messages"`
	CanManageVoiceChats bool   `json:"can_manage_voice_chats"`
	CanRestrictMembers  bool   `json:"can_restrict_members"`
	CanPromoteMembers   bool   `json:"can_promote_members"`
	CanChangeInfo       bool   `json:"can_change_info"`
	CanInviteUsers      bool   `json:"can_invite_users"`
	CanPostMessages     bool   `json:"can_post_messages"`
	CanEditMessages     bool   `json:"can_edit_messages"`
	CanPinMessages      bool   `json:"can_pin_messages"`
}

type ChatMemberMember struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

type ChatMemberRestricted struct {
	Status                string `json:"status"`
	User                  User   `json:"user"`
	IsMember              bool   `json:"is_member"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	UntilDate             bool   `json:"until_date"`
}

type ChatMemberLeft struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

type ChatMemberBanned struct {
	Status    string `json:"status"`
	User      User   `json:"user"`
	UntilDate bool   `json:"until_date"`
}

type Location struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`
	LivePeriod           int     `json:"live_period"`
	Heading              int     `json:"heading"`
	ProximityAlertRadius int     `json:"proximity_alert_radius"`
}

type Venue struct {
	Location
	Title           string `json:"title"`
	Address         string `json:"address"`
	FoursquareId    string `json:"foursquare_id"`
	FoursquareType  string `json:"foursquare_type"`
	GooglePlaceId   string `json:"google_place_id"`
	GooglePlaceType string `json:"google_place_type"`
}

type ChatLocation struct {
	Address  string   `json:"address"`
	Location Location `json:"location"`
}

type Animation struct {
	FileId string `json:"file_id"`
}

type InputMedia interface {
	// returnFile return *os.File if Media field of InputMedia is a file and automatically set the Media to
	// the correct value: attach://<file name>
	// Methods like MediaGroupData.Send() use this method to add those files to request.
	// Be aware that after running returnFile, the Media field is no longer a file, it is a string.
	// So if you intend to use a InputMedia twice, set the Media field again.
	returnFile() *os.File
}

type InputMediaPhoto struct {
	// Type of the result, must be "photo"
	Type string `json:"type"`
	// Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet or pass a *os.File
	Media any `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
}

func (i *InputMediaPhoto) returnFile() *os.File {
	i.Type = "photo"
	switch m := i.Media.(type) {
	case *os.File:
		i.Media = "attach://" + m.Name()
		return m
	}
	return nil
}

type InputMediaVideo struct {
	// Type of the result, must be "video"
	Type string `json:"type"`
	// Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet or pass a *os.File
	Media any `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption           string          `json:"caption"`
	ParseMode         string          `json:"parse_mode"`
	Width             int             `json:"width"`
	Height            int             `json:"height"`
	Duration          int             `json:"duration"`
	SupportsStreaming bool            `json:"supports_streaming"`
	CaptionEntities   []MessageEntity `json:"caption_entities"`
}

func (i *InputMediaVideo) returnFile() *os.File {
	i.Type = "video"
	switch m := i.Media.(type) {
	case *os.File:
		i.Media = "attach://" + m.Name()
		return m
	}
	return nil
}

type InputMediaDocument struct {
	// Type of the result, must be "document"
	Type string `json:"type"`
	// Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet or pass a *os.File
	Media any `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Optional. Disables automatic server-side content type detection for files uploaded using
	// multipart/form-data. Always true, if the document is sent as part of an album.
	DisableContentTypeDetection bool `json:"disable_content_type_detection"`
}

func (i *InputMediaDocument) returnFile() *os.File {
	i.Type = "document"
	switch m := i.Media.(type) {
	case *os.File:
		i.Media = "attach://" + m.Name()
		return m
	}
	return nil
}

type InputMediaAudio struct {
	// Type of the result, must be "audio"
	Type string `json:"type"`
	// Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet or pass a *os.File
	Media any `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	Duration        int             `json:"duration"`
	Performer       string          `json:"performer"`
	Tile            string          `json:"tile"`
}

func (i *InputMediaAudio) returnFile() *os.File {
	i.Type = "audio"
	switch m := i.Media.(type) {
	case *os.File:
		i.Media = "attach://" + m.Name()
		return m
	}
	return nil
}

type InputMediaAnimation struct {
	// Type of the result, must be "animation"
	Type string `json:"type"`
	// Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet or pass a *os.File
	Media any `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	Width           int             `json:"width"`
	Height          int             `json:"height"`
}

func (i *InputMediaAnimation) returnFile() *os.File {
	i.Type = "animation"
	switch m := i.Media.(type) {
	case *os.File:
		i.Media = "attach://" + m.Name()
		return m
	}
	return nil
}

type LoginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text"`
	BotUsername        string `json:"bot_username"`
	RequestWriteAccess bool   `json:"request_write_access"`
}

func (i LoginUrl) check() error {
	if i.Url == "" {
		return errors.New("url of LoginUrl is empty")
	}
	return nil
}

// CallbackGame is a placeholder, currently holds no information. Use BotFather to set up your game
// and set Active to true.
type CallbackGame struct {
	Active bool
}

type KeyboardButtonPollType struct {
	// Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode.
	// If regular is passed, only regular polls will be allowed.
	// Otherwise, the user will be allowed to create a poll of any type.
	Type string `json:"type"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	YShift float64 `json:"y_shift"`
	XShift float64 `json:"x_shift"`
	Scale  float64 `json:"scale"`
}

type Sticker struct {
	FileId       string       `json:"file_id"`
	FileUniqueId string       `json:"file_unique_id"`
	Width        int          `json:"width"`
	Height       int          `json:"height"`
	IsAnimated   bool         `json:"is_animated"`
	Thumb        PhotoSize    `json:"thumb"`
	Emoji        string       `json:"emoji"`
	SetName      string       `json:"set_name"`
	MaskPosition MaskPosition `json:"mask_position"`
	FileSize     int          `json:"file_size"`
}

type StickerSet struct {
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	IsAnimated    bool      `json:"is_animated"`
	ContainsMasks bool      `json:"contains_masks"`
	Stickers      []Sticker `json:"stickers"`
	Thumb         PhotoSize `json:"thumb"`
}

type ReplyKeyboardOP struct {
	Horizontal            bool   `json:"horizontal"`
	OneTimeKeyboard       bool   `json:"one_time_keyboard"`
	ResizeKeyboard        bool   `json:"resize_keyboard"`
	InputFieldPlaceholder string `json:"input_field_placeholder"`
	Selective             bool   `json:"selective"`
}

type ChatInviteLink struct {
	InviteLink  string `json:"invite_link"`
	Creator     User   `json:"creator"`
	IsPrimary   bool   `json:"is_primary"`
	IsRevoked   bool   `json:"is_revoked"`
	ExpireDate  int    `json:"expire_date"`
	MemberLimit int    `json:"member_limit"`
}

type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

type UserProfilePhotos struct {
	// Total number of profile pictures the target user has
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

type Response interface {
	print()
	getResult() any
	set(*http.Response) (ResponseImpl, error)
	isOk() bool
	getDescription() string
	getErrorCode() int
}

type ResponseImpl struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Result      any    `json:"result"`
}

func (r ResponseImpl) print() {

}

func (r ResponseImpl) getResult() any {
	return r.Result
}

func (r ResponseImpl) set(res *http.Response) (ResponseImpl, error) {
	readRes, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(readRes, &r)
	if err != nil {
		return r, err
	}
	if r.Ok != true {
		return r, errors.New("telegram returned an error. check response for more details")
	}
	return r, nil
}

func (r ResponseImpl) isOk() bool {
	return r.Ok
}

func (r ResponseImpl) getDescription() string {
	return r.Description
}

func (r ResponseImpl) getErrorCode() int {
	return r.ErrorCode
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type BotCommandScope interface {
	checkScope() error
}

type BotCommandScopeDefault struct {
	Type string `json:"type"`
}

func (b BotCommandScopeDefault) checkScope() error {
	if b.Type != "default" {
		return errors.New("botCommandScope type must be default")
	}
	return nil
}

type BotCommandScopeAllGroupChats struct {
	Type string `json:"type"`
}

func (b BotCommandScopeAllGroupChats) checkScope() error {
	if b.Type != "all_group_chats" {
		return errors.New("botCommandScope type must be all_group_chats")
	}
	return nil
}

type BotCommandScopeChat struct {
	Type   string `json:"type"`
	ChatId int    `json:"chat_id"`
}

func (b BotCommandScopeChat) checkScope() error {
	if b.Type != "chat" {
		return errors.New("botCommandScope type must be chat")
	}
	return nil
}

type BotCommandScopeChatAdministrators struct {
	Type   string `json:"type"`
	ChatId int    `json:"chat_id"`
}

func (b BotCommandScopeChatAdministrators) checkScope() error {
	if b.Type != "chat_administrators" {
		return errors.New("botCommandScope type must be chat_administrators")
	}
	return nil
}

type BotCommandScopeChatMember struct {
	Type   string `json:"type"`
	ChatId int    `json:"chat_id"`
	UserId int    `json:"user_id"`
}

func (b BotCommandScopeChatMember) checkScope() error {
	if b.Type != "chat_member" {
		return errors.New("botCommandScope type must be chat_member")
	}
	return nil
}

type InlineKeyboard struct {
	Buttons [][]InlineButton `json:"inline_keyboard"`
}

func (i *InlineKeyboard) AddInlineButtons(horizontal bool, a ...InlineButton) error {
	var buttons [][]InlineButton
	if horizontal {
		buttons = append(buttons, []InlineButton{})
		for _, button := range a {
			if err := button.check(); err != nil {
				return err
			}
			buttons[0] = append(buttons[0], button)
		}
	} else {
		for _, button := range a {
			if err := button.check(); err != nil {
				return err
			}
			buttons = append(buttons, []InlineButton{button})
		}
	}
	i.Buttons = append(i.Buttons, buttons...)
	return nil
}

type ReplyKeyboard struct {
	Keyboard              [][]ReplyButton `json:"keyboard"`
	OneTimeKeyboard       bool            `json:"one_time_keyboard"`
	ResizeKeyboard        bool            `json:"resize_keyboard"`
	InputFieldPlaceholder string          `json:"input_field_placeholder"`
	Selective             bool            `json:"selective"`
}

func (r *ReplyKeyboard) AddReplyButtons(optionalParams ReplyKeyboardOP, a ...ReplyButton) error {
	r.OneTimeKeyboard = optionalParams.OneTimeKeyboard
	r.Selective = optionalParams.Selective
	r.InputFieldPlaceholder = optionalParams.InputFieldPlaceholder
	r.ResizeKeyboard = optionalParams.ResizeKeyboard
	var buttons [][]ReplyButton
	if optionalParams.Horizontal {
		buttons = append(buttons, []ReplyButton{})
		for _, button := range a {
			if err := button.check(); err != nil {
				return err
			}
			buttons[0] = append(buttons[0], button)
		}
	} else {
		for _, button := range a {
			if err := button.check(); err != nil {
				return err
			}
			buttons = append(buttons, []ReplyButton{button})
		}
	}
	r.Keyboard = append(r.Keyboard, buttons...)
	return nil
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

// Remove sets ReplyKeyboardRemove fields.
// Set selective to true if you want to remove the keyboard for specific users only.
// Targets: 1) users that are @mentioned in the text of the Message object;
//          2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
// Example: A user votes in a poll, bot returns confirmation message in reply
// to the vote and removes the keyboard for that user,
// while still showing the keyboard with poll options to users who haven't voted yet.
func (r *ReplyKeyboardRemove) Remove(selective bool) {
	r.RemoveKeyboard = true
	r.Selective = selective
}

type ForceReply struct {
	IsForceReply          bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder"`
	Selective             bool   `json:"selective"`
}

func (t *ForceReply) SetForceReply(selective bool, inputFieldPlaceholder string) {
	t.IsForceReply = true
	t.Selective = selective
	t.InputFieldPlaceholder = inputFieldPlaceholder
}

type Keyboard struct {
	ReplyMarkup any `json:"reply_markup"`
}

func (k *Keyboard) SetInlineKeyboard(horizontal bool, a ...InlineButton) error {
	inlineKeyboard, ok := (k.ReplyMarkup).(InlineKeyboard)
	if ok {
		if err := inlineKeyboard.AddInlineButtons(horizontal, a...); err != nil {
			return err
		}
		k.ReplyMarkup = inlineKeyboard
		return nil
	}
	i := InlineKeyboard{}
	if err := i.AddInlineButtons(horizontal, a...); err != nil {
		return err
	}
	k.ReplyMarkup = i
	return nil
}

// SetReplyKeyboard adds reply keyboard to message. optionalParams is optional and you can pass
// an empty ReplyKeyboardOP
func (k *Keyboard) SetReplyKeyboard(optionalParams ReplyKeyboardOP, a ...ReplyButton) error {
	replyKeyboard, ok := (k.ReplyMarkup).(ReplyKeyboard)
	if ok {
		if err := replyKeyboard.AddReplyButtons(optionalParams, a...); err != nil {
			return err
		}
		k.ReplyMarkup = replyKeyboard
		return nil
	}
	r := ReplyKeyboard{}
	if err := r.AddReplyButtons(optionalParams, a...); err != nil {
		return err
	}
	k.ReplyMarkup = r
	return nil
}

// RemoveReplyKeyboard removes reply keyboard.
// Set selective to true if you want to remove the keyboard for specific users only.
// Targets: 1) users that are @mentioned in the text of the Message object;
//          2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
// Example: A user votes in a poll, bot returns confirmation message in reply
// to the vote and removes the keyboard for that user,
// while still showing the keyboard with poll options to users who haven't voted yet.
func (k *Keyboard) RemoveReplyKeyboard(selective bool) {
	i := ReplyKeyboardRemove{}
	i.Remove(selective)
	k.ReplyMarkup = i
}

func (k *Keyboard) ForceReply(selective bool, inputFieldPlaceholder string) {
	i := ForceReply{}
	i.SetForceReply(selective, inputFieldPlaceholder)
	k.ReplyMarkup = i
}

// InlineButton represents one button of an inline keyboard.
// You must use exactly one of the optional fields.
type InlineButton struct {
	// Label text on the button
	Text string `json:"text"`
	// Optional. HTTP or tg:// url to be opened
	// when button is pressed
	Url      string   `json:"url"`
	LoginUrl LoginUrl `json:"login_url"`
	// Optional. Data to be sent in a callback query
	// to the bot when button is pressed
	CallbackData string `json:"callback_data"`
	// Optional. If set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	// Can be empty, in which case just the bot's username will be inserted.
	// Note: This offers an easy way for users to start using your bot in inline mode
	// when they are currently in a private chat with it.
	// Especially useful when combined with switch_pm… actions – in this case the user will be
	// automatically returned to the chat they switched from, skipping the chat selection screen.
	SwitchInlineQuery string `json:"switch_inline_query"`
	// Optional. If set, pressing the button will insert the bot's username and the specified
	// inline query in the current chat's input field.
	// Can be empty, in which case only the bot's username will be inserted.
	// This offers a quick way for the user to open your bot in inline mode
	// in the same chat – good for selecting something from multiple options.
	SwitchInlineQueryCurrentChat string       `json:"switch_inline_query_current_chat"`
	CallbackGame                 CallbackGame `json:"callback_game"`
	// Optional. Specify True, to send a Pay button.
	// NOTE: This type of button must always be the first button in the first row.
	Pay bool `json:"pay"`
}

func (i InlineButton) check() error {
	if i.Text == "" {
		return errors.New("text of InlineButton is empty")
	}
	notEmpty := 0
	if i.Url != "" {
		notEmpty += 1
	}
	if i.Pay == true {
		notEmpty += 1
	}
	if i.CallbackData != "" {
		notEmpty += 1
	}
	if i.SwitchInlineQuery != "" {
		notEmpty += 1
	}
	if i.SwitchInlineQueryCurrentChat != "" {
		notEmpty += 1
	}
	if i.LoginUrl.check() == nil {
		notEmpty += 1
	}
	if i.CallbackGame.Active == true {
		notEmpty += 1
	}
	if notEmpty != 1 {
		return errors.New("you must set exactly one of the optional fields of InlineButton")
	}
	return nil
}

// ReplyButton represents one button of a reply keyboard.
type ReplyButton struct {
	// Text of the button. If none of the optional fields are used,
	// it will be sent as a message when the button is pressed
	Text string `json:"text"`
	// Optional. If True, the user's phone number will be sent as a contact when the button is pressed.
	// Available in private chats only
	RequestContact bool `json:"request_contact"`
	// Optional. If True, the user's current location will be sent when the button is pressed.
	// Available in private chats only
	RequestLocation bool `json:"request_location"`
	// Optional. If specified, the user will be asked to create a poll and send it to the bot
	// when the button is pressed. Available in private chats only
	RequestPoll KeyboardButtonPollType `json:"request_poll"`
}

func (i ReplyButton) check() error {
	if i.Text == "" {
		return errors.New("text of ReplyButton is empty")
	}
	return nil
}
