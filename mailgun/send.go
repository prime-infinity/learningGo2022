package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func main() {

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "API_KEY"
	}

	id, err := SendSimpleMessage("sandbox8ceb7e7e2ac140c7975102cacdd94eb6.mailgun.org", apiKey)
	fmt.Println(id)
	fmt.Println(err)

}

func SendSimpleMessage(domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	//When you have an EU-domain, you must specify the endpoint:
	// mg.SetAPIBase("https://api.eu.mailgun.net")
	m := mg.NewMessage(
		"Mailgun Sandbox <postmaster@sandbox8ceb7e7e2ac140c7975102cacdd94eb6.mailgun.org>",
		"Hello Akonofua Osamede",
		"Congratulations Akonofua Osamede, you just sent an email with Mailgun! You are truly awesome!",
		"Akonofua Osamede <admin@sellperview.com>",
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}
