package plume

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v3"
	"golang.org/x/time/rate"
)

// User represents a user with the intent to join
// a Plume chat session
type User struct {
	Name  string `json:"name"`
	Email string `json:"-"`
}

const domain = "mail.plume.chat"

var apiKey = os.Getenv("MAILGUN_APIKEY")

// 1 every 4 seconds, max 1 call at once
var mailLimiter = rate.NewLimiter(0.25, 1)

func (u User) sendAuthEmail(token string) {
	if !mailLimiter.Allow() {
		log.Printf("Mail send rate limit exceeded by %s\n", u.Email)
		return
	}

	log.Printf("Sending token for %s: %s", u.Name, token)

	from := "Plume Chat <hi@plume.chat>"
	subject := "Your Plume login code"
	body := fmt.Sprintf(
		"Hi @%s! Your login code to Plume.chat is \"%s\".",
		u.Name,
		token,
	)

	if environment != "production" {
		log.Printf("Not sending login email as ENV != production\n")
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
