package gogram

// TextOptionalParams represents optional parameters
// that SendText function can use.
// All fields are optional.
type TextOptionalParams struct {
	parseMode                string
	entities                 interface{}
	disableWebPagePreview    bool
	disableNotification      bool
	replyToMessageId         int
	allowSendingWithoutReply bool
	replyMarkup              interface{}
}

//// TextOptionalParams represents optional parameters
//// that SendText function can use.
//// All fields are optional.
//type PhotoOptionalParams struct {
//	parseMode                string
//	entities                 interface{}
//	disableWebPagePreview    bool
//	disableNotification      bool
//	replyToMessageId         int
//	allowSendingWithoutReply bool
//	replyMarkup              interface{}
//}
