package goutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	gomail "gopkg.in/mail.v2"
)

func SendSMTPMessage(host, password, from, to, subject, content string) error {
	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {from},
		"To":      {to},
		"Subject": {subject},
	})
	m.SetBody("text/plain", content)
	return gomail.NewDialer(host, 465, from, password).DialAndSend(m)
}

func SendTelegramMessage(token, chatID, text string) (*http.Response, error) {
	data, err := json.Marshal(map[string]string{
		"chat_id": chatID,
		"text":    text,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token),
		"Content-Type: application/json",
		bytes.NewBuffer(data),
	)
	return resp, err
}

func DefaultMessageTemplate(title, from, level, detail string) string {
	return fmt.Sprintf(
		"title: %s\nfrom: %s\nlevel: %s\ntime: %s\ndetail:\n%s",
		title, from, level, time.Now().Format(time.RFC3339), detail,
	)
}
