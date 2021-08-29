package gogram

func TypeIndicator(message Message) string {
	switch {
	case message.Text != "":
		return "Text"
	case message.Animation != Animation{}:
		return "Animation"
	default:
		return ""
	}
}
