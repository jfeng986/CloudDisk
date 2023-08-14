package utils

import (
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/jordan-wright/email"
)

func MailSendCode(toEmail string) error {
	err := godotenv.Load("../.env")
	if err != nil {
		return err
	}
	emailAddress := os.Getenv("EMAIL_ADDRESS")
	emailPassword := os.Getenv("APP_PASSWORD")
	code := GenCode()
	e := email.NewEmail()
	e.From = "Get <jfeng986@gmail.com>"
	e.To = []string{toEmail}
	e.Subject = "Code has been sent, please check it out!"
	e.HTML = []byte("Your auth code: <b>" + code + "</b>")
	auth := smtp.PlainAuth("", emailAddress, emailPassword, "smtp.gmail.com")
	return e.Send("smtp.gmail.com:587", auth)
}

func GenCode() string {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	res := ""
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rng.Intn(10))
	}
	return res
}
