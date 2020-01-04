package plume

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v3"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"-"`
}

const domain = "mail.plume.chat"

var apiKey = os.Getenv("MAILGUN_APIKEY")

func (u User) SendAuthEmail(token string) {
	log.Printf("Sending token for %s: %s", u.Name, token)

	from := "Plume Chat <hi@plume.chat>"
	subject := "Your Plume login code"
	body := fmt.Sprintf(
		"Hi %s! Your login code to plume.chat is \"%s\".",
		u.Name,
		u.Email,
	)

	if Environment != "production" {
		log.Printf("Not sending login email as ENV != production")
		return
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	mail := mg.NewMessage(
		from,
		subject,
		body,
		u.Email, // recipient
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	resp, id, err := mg.Send(ctx, mail)
	if err != nil {
		log.Printf("Error sending login token email: %v", err)
	} else {
		log.Printf("Sent token to %s: resp(%s) id(%s)", u.Email, resp, id)
	}
}
