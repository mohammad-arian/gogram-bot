package gogram

import (
	"errors"
	"os"
	"reflect"
)

// Update from webhook
type Update struct {
	UpdateId      int           `json:"update_id"`
	Message       Message       `json:"message"`
	InlineQuery   InlineQuery   `json:"inline_query"`
	CallbackQuery CallbackQuery `json:"callback_query"`
	Poll          Poll          `json:"poll"`
}

type CallbackQuery struct {
	Id           string  `json:"id"`
	Message      Message `json:"message"`
	Chat         Chat    `json:"chat"`
	ChatInstance string  `json:"chat_instance"`
	Data         string  `json:"data"`
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
	// fileId is Identifier for this file, which can be used to download or reuse the file
	fileId string
	// fileUniqueId isUnique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	fileUniqueId string
	// fileUniqueId is File size in bytes, if known. Optional
	fileSize int
	// filePath is File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file. Optional
	filePath string
}

type Message struct {
	MessageId       int                  `json:"message_id"`
	User            User                 `json:"from"`
	Chat            Chat                 `json:"chat"`
	Text            string               `json:"text"`
	Animation       Animation            `json:"animation"`
	Photo           []PhotoSize          `json:"photo"`
	Date            int                  `json:"date"`
	ReplyMarkup     inlineKeyboardMarkup `json:"reply_markup"`
	Poll            Poll                 `json:"poll"`
	NewChatPhoto    []PhotoSize          `json:"new_chat_photo"`
	NewChatTitle    string               `json:"new_chat_title"`
	NewChatMembers  []User               `json:"new_chat_members"`
	DeleteChatPhoto bool                 `json:"delete_chat_photo"`
}

// TypeIndicator function returns the type of message
// This make it easier to know which fields are empty and which aren't
// TypeIndicator may return "Text", "Animation", "Photo" and etc
func (m Message) TypeIndicator() string {
	switch {
	case m.Text != "":
		return "Text"
	case m.Animation != Animation{}:
		return "Animation"
	case m.Photo != nil:
		return "Photo"
	case m.DeleteChatPhoto == true:
		return "DeleteChatPhoto"
	case m.NewChatPhoto != nil:
		return "NewChatPhoto"
	default:
		return "Unknown"
	}
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

type Location struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontalAccuracy"`
	LivePeriod           int     `json:"livePeriod"`
	Heading              int     `json:"heading"`
	ProximityAlertRadius int     `json:"proximityAlertRadius"`
}

type Venue struct {
	Location
	Title           string `json:"title"`
	Address         string `json:"address"`
	FoursquareId    string `json:"foursquareId"`
	FoursquareType  string `json:"foursquareType"`
	GooglePlaceId   string `json:"googlePlaceId"`
	GooglePlaceType string `json:"googlePlaceType"`
}

type ChatLocation struct {
	Address  string   `json:"address"`
	Location Location `json:"location"`
}

type Animation struct {
	FileId string `json:"file_id"`
}

// InputMediaPhoto Represents a photo to be sent.
type InputMediaPhoto struct {
	// Type of the result, must be "photo"
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass a file of type *os.File
	// to upload a new one.
	Media interface{} `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
}

func (i *InputMediaPhoto) setMediaAndType(files *[]*os.File) error {
	switch v := i.Media.(type) {
	case *os.File:
		i.Media = "attach://" + v.Name()
		*files = append(*files, v)
	case string:
	default:
		return errors.New("your Media is of type " + reflect.TypeOf(i.Media).String() + ". " +
			"Media must be of type *os.File for files or string for file_ids and urls")
	}
	i.Type = "photo"
	return nil
}

type InputMediaVideo struct {
	// Type of the result, must be "video"
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass a file of type *os.File
	// to upload a new one.
	Media interface{} `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption           string          `json:"caption"`
	ParseMode         string          `json:"parse_mode"`
	Width             int             `json:"width"`
	Height            int             `json:"height"`
	Duration          int             `json:"duration"`
	SupportsStreaming bool            `json:"supports_streaming"`
	CaptionEntities   []MessageEntity `json:"caption_entities"`
}

func (i *InputMediaVideo) setMediaAndType(files *[]*os.File) error {
	switch v := i.Media.(type) {
	case *os.File:
		i.Media = "attach://" + v.Name()
		*files = append(*files, v)
	case string:
	default:
		return errors.New("your Media is of type " + reflect.TypeOf(i.Media).String() + ". " +
			"Media must be of type *os.File for files or string for file_ids and urls")
	}
	i.Type = "video"
	return nil
}

type InputMediaDocument struct {
	// Type of the result, must be "document"
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass a file of type *os.File
	// to upload a new one.
	Media interface{} `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Optional. Disables automatic server-side content type detection for files uploaded using
	// multipart/form-data. Always true, if the document is sent as part of an album.
	DisableContentTypeDetection bool `json:"disable_content_type_detection"`
}

func (i *InputMediaDocument) setMediaAndType(files *[]*os.File) error {
	switch v := i.Media.(type) {
	case *os.File:
		i.Media = "attach://" + v.Name()
		*files = append(*files, v)
	case string:
	default:
		return errors.New("your Media is of type " + reflect.TypeOf(i.Media).String() + ". " +
			"Media must be of type *os.File for files or string for file_ids and urls")
	}
	i.Type = "document"
	return nil
}

type InputMediaAudio struct {
	// Type of the result, must be "audio"
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass a file of type *os.File
	// to upload a new one.
	Media interface{} `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	Duration        int             `json:"duration"`
	Performer       string          `json:"performer"`
	Tile            string          `json:"tile"`
}

func (i *InputMediaAudio) setMediaAndType(files *[]*os.File) error {
	switch v := i.Media.(type) {
	case *os.File:
		i.Media = "attach://" + v.Name()
		*files = append(*files, v)
	case string:
	default:
		return errors.New("your Media is of type " + reflect.TypeOf(i.Media).String() + ". " +
			"Media must be of type *os.File for files or string for file_ids and urls")
	}
	i.Type = "audio"
	return nil
}

type InputMediaAnimation struct {
	// Type of the result, must be "animation"
	Type string `json:"type"`
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass
	// "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name>
	Media interface{} `json:"media"`
	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode"`
	CaptionEntities []MessageEntity `json:"caption_entities"`
	Width           int             `json:"width"`
	Height          int             `json:"height"`
}

func (i *InputMediaAnimation) setMediaAndType(files *[]*os.File) error {
	switch v := i.Media.(type) {
	case *os.File:
		i.Media = "attach://" + v.Name()
		*files = append(*files, v)
	case string:
	default:
		return errors.New("your Media is of type " + reflect.TypeOf(i.Media).String() + ". " +
			"Media must be of type *os.File for files or string for file_ids and urls")
	}
	i.Type = "animation"
	return nil
}

type MaskPosition struct {
	Point  string  `json:"point"`
	YShift float64 `json:"YShift"`
	XShift float64 `json:"XShift"`
	Scale  float64 `json:"Scale"`
}

type Sticker struct {
	FileId       string       `json:"fileId"`
	FileUniqueId string       `json:"fileUniqueId"`
	Width        int          `json:"width"`
	Height       int          `json:"height"`
	IsAnimated   bool         `json:"isAnimated"`
	Thumb        PhotoSize    `json:"thumb"`
	Emoji        string       `json:"emoji"`
	SetName      string       `json:"setName"`
	MaskPosition MaskPosition `json:"maskPosition"`
	FileSize     int          `json:"fileSize"`
}

type StickerSet struct {
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	IsAnimated    bool      `json:"isAnimated"`
	ContainsMasks bool      `json:"containsMasks"`
	Stickers      []Sticker `json:"stickers"`
	Thumb         PhotoSize `json:"thumb"`
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

type Response struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type UserProfileResponse struct {
	Result UserProfilePhotos `json:"result"`
	Response
}

type BooleanResponse struct {
	Result bool `json:"result"`
	Response
}

type IntResponse struct {
	Result int `json:"result"`
	Response
}

type MapResponse struct {
	Result interface{} `json:"result"`
	Response
}

type PollResponse struct {
	Result Poll `json:"result"`
	Response
}

type UserResponse struct {
	Result User `json:"result"`
	Response
}

type MessageResponse struct {
	Result Message `json:"result"`
	Response
}

type InviteLinkResponse struct {
	Result ChatInviteLink `json:"result"`
	Response
}

type ChatResponse struct {
	Result Chat `json:"result"`
	Response
}

type FileResponse struct {
	Result File `json:"result"`
	Response
}

type ChatMemberResponse struct {
	Result []ChatMember `json:"result"`
	Response
}

type StickerSetResponse struct {
	Result StickerSet `json:"result"`
	Response
}

type BotCommandResponse struct {
	Result []BotCommand `json:"result"`
	Response
}

func (c *ChatMemberResponse) permissionSetter() {
	for j := range c.Result {
		if c.Result[j].Status != "restricted" {
			c.Result[j].UntilDate = -1
		}
		if c.Result[j].Status == "creator" {
			c.Result[j].IsMember = true
			c.Result[j].CanPostMessages = true
			c.Result[j].CanInviteUsers = true
			c.Result[j].CanSendPolls = true
			c.Result[j].CanAddWebPagePreviews = true
			c.Result[j].CanChangeInfo = true
			c.Result[j].CanSendOtherMessages = true
			c.Result[j].CanSendMessages = true
			c.Result[j].CanDeleteMessages = true
			c.Result[j].CanManageChat = true
			c.Result[j].CanPromoteMembers = true
			c.Result[j].CanSendMediaMessages = true
			c.Result[j].CanRestrictMembers = true
			c.Result[j].CanPinMessages = true
			c.Result[j].CanManageVoiceChats = true
			c.Result[j].CanEditMessages = true
			c.Result[j].CanSendPolls = true
		}
		if c.Result[j].Status == "administrator" {
			c.Result[j].IsMember = true
			c.Result[j].CanSendPolls = true
			c.Result[j].CanSendMediaMessages = true
			c.Result[j].CanSendOtherMessages = true
			c.Result[j].CanAddWebPagePreviews = true
			c.Result[j].CanSendMessages = true
		}
		if c.Result[j].Status == "member" {
			c.Result[j].IsMember = true
			c.Result[j].CanPostMessages = true
			c.Result[j].CanInviteUsers = true
			c.Result[j].CanSendPolls = true
			c.Result[j].CanAddWebPagePreviews = true
			c.Result[j].CanChangeInfo = true
			c.Result[j].CanSendMessages = true
			c.Result[j].CanPinMessages = true
			c.Result[j].CanBeEdited = true
		}
	}
}

type ChatMember struct {
	Status              string `json:"status"`
	User                User   `json:"user"`
	IsAnonymous         bool   `json:"is_anonymous"`
	CustomTitle         string `json:"custom_title"`
	IsMember            bool   `json:"is_member"`
	CanBeEdited         bool   `json:"can_be_edited"`
	CanManageChat       bool   `json:"can_manage_chat"`
	CanDeleteMessages   bool   `json:"can_delete_messages"`
	CanManageVoiceChats bool   `json:"can_manage_voice_chats"`
	CanRestrictMembers  bool   `json:"can_restrict_members"`
	CanPromoteMembers   bool   `json:"can_promote_members"`
	CanPostMessages     bool   `json:"can_post_messages"`
	CanEditMessages     bool   `json:"can_edit_messages"`
	// if member is restricted, UntilDate is the date when restrictions will be lifted for this user;
	// unix time. If 0, then the user is restricted forever. If -1 user is not
	// restricted.
	UntilDate int `json:"until_date"`
	ChatPermissions
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

// BotCommandScope Represents the scope of bot commands.
type BotCommandScope struct {
	// Type is the scope type. It can be:
	// "default"                 -> Default commands are used if no commands with a narrower
	//                              scope are specified for the user.
	// "chat_member"             -> covers a specific member of a group or supergroup chat.
	// "all_private_chats"       -> covers all private chats.
	// "all_group_chats"         -> covers all group and supergroup chats.
	// "all_chat_administrators" -> covers all group and supergroup chat administrators.
	// "chat"                    -> covers a specific chat.
	// "chat_administrators"     -> covers all administrators of a specific group or supergroup chat.
	Type string `json:"type"`
	// ChatId is unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername).
	// Required only if Type is "chat_administrators", "chat" or "chat_member".
	ChatId int `json:"chat_id"`
	// UserId is unique identifier of the target user.
	// Required only if Type is "chat_member"
	UserId int `json:"user_id"`
}
