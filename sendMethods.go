package gogram

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// SendMessageToUser sends message to a User.
// b Bot parameter indicated which bot to send
// the message with. This way you can send messages
// with different bots
func (u User) SendMessageToUser(b Bot, text string) {
	if u.Id == 0 {
		log.Fatalln("User's Id field is empty")
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token), nil)
	if err != nil {
		log.Fatalln(err)
	}
	q := req.URL.Query()
	q.Add("chat_id", strconv.Itoa(u.Id))
	q.Add("text", text)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
}
