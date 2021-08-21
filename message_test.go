package goutils

import (
	"os"
	"testing"
)

func TestSMTPMessageSender(t *testing.T) {
	subject := `hello`
	content := `hello world`
	sendMail := SMTPMessageSender(
		os.Getenv("EMAIL_HOST"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("EMAIL_FROM"),
		os.Getenv("EMAIL_TO"),
	)
	err := sendMail(subject, content)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTelegramMessageSender(t *testing.T) {
	text := "hello world"
	sendTG := TelegramMessageSender(
		os.Getenv("TELEGRAM_TOKEN"),
		os.Getenv("TELEGRAM_CHATID"),
	)
	if _, err := sendTG(text); err != nil {
		t.Fatal(err)
	}

}

func TestDefaultMessageTemplate(t *testing.T) {
	title := "this is title"
	from := "go-utils"
	level := "debug"
	detail := "more info...."
	r := DefaultMessageTemplate(title, from, level, detail)
	t.Log(r)
}
