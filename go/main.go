package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/nicklaw5/helix/v2"
)

var client *helix.Client = loadClient()
var channels = []string{"channel1", "channel2"}

func loadClient() *helix.Client {
	client, err := helix.NewClient(&helix.Options{
		ClientID:       os.Getenv("CLIENT_ID"),
		ClientSecret:   os.Getenv("CLIENT_SECRET"),
		AppAccessToken: os.Getenv("ACCESS_TOKEN"),
	})
	if err != nil {
		panic(err)
	}

	return client
}

func main() {

	resp, err := client.GetUsers(&helix.UsersParams{
		Logins: channels,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status code: %d\n", resp.StatusCode)
	fmt.Printf("Rate limit: %d\n", resp.GetRateLimit())
	fmt.Printf("Rate limit remaining: %d\n", resp.GetRateLimitRemaining())
	fmt.Printf("Rate limit reset: %d\n\n", resp.GetRateLimitReset())

	for _, user := range resp.Data.Users {
		fmt.Printf("ID: %s Name: %s\n", user.ID, user.DisplayName)
	}
}
