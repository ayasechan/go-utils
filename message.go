package goutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	gomail "gopkg.in/mail.v2"
)

func SMTPMessageSender(host, password, from, to string) func(subject, content string) error {
	d := gomail.NewDialer(host, 465, from, password)
	return func(subject, content string) error {
		m := gomail.NewMessage()
		m.SetHeaders(map[string][]string{
			"From":    {from},
			"To":      {to},
			"Subject": {subject},
		})
		m.SetBody("text/plain", content)
		return d.DialAndSend(m)
	}

}

func TelegramMessageSender(token, chatID string) func(text string) (*http.Response, error) {
	return func(text string) (*http.Response, error) {
		data, err := json.Marshal(map[string]string{
			"chat_id": chatID,
			"text":    text,
		})
		if err != nil {
			return nil, err
		}
		return http.Post(
			fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token),
			"Content-Type: application/json",
			bytes.NewBuffer(data),
		)
	}
}

func DefaultMessageTemplate(title, from, level, detail string) string {
	return fmt.Sprintf(
		"title: %s\nfrom: %s\nlevel: %s\ntime: %s\ndetail:\n%s",
		title, from, level, time.Now().Format(time.RFC3339), detail,
	)
}
