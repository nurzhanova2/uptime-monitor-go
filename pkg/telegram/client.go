package telegram

import (
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	Token  string
	ChatID int64
}

func NewClient(token string, chatID int64) *Client {
	return &Client{Token: token, ChatID: chatID}
}

func (c *Client) Send(msg string) error {
	base := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.Token)
	data := url.Values{}
	data.Set("chat_id", fmt.Sprintf("%d", c.ChatID))
	data.Set("text", msg)

	_, err := http.PostForm(base, data)
	return err
}
