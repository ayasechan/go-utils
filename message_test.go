package goutils

import (
	"os"
	"testing"
)

func TestSendSMTPMessage(t *testing.T) {

	subject := `hello`
	content := `hello world`
	err := SendSMTPMessage(
		os.Getenv("EMAIL_HOST"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("EMAIL_FROM"),
		os.Getenv("EMAIL_TO"),
		subject,
		content,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendTelegramMessage(t *testing.T) {
	text := "hello world"
	if _, err := SendTelegramMessage(
		os.Getenv("TELEGRAM_TOKEN"),
		os.Getenv("TELEGRAM_CHATID"),
		text,
	); err != nil {
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
