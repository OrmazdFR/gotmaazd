package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/nicklaw5/helix/v2"
)

var client *helix.Client = loadClient()
var channels = []string{"channel1", "channel2"}
var scopes = []string{
	"analytics:read:games",
	"bits:read",
	"channel:manage:broadcast",
	"channel:manage:moderators",
	"channel:manage:polls",
	"channel:manage:predictions",
	"channel:manage:raids",
	"channel:manage:redemptions",
	"channel:read:hype_train",
	"channel:read:polls",
	"channel:read:predictions",
	"channel:read:redemptions",
	"channel:read:subscriptions",
	"channel:read:vips",
	"channel:manage:vips",
	"moderation:read",
	"moderator:manage:announcements",
	"moderator:manage:banned_users",
	"moderator:manage:blocked_terms",
	"moderator:manage:chat_messages",
	"moderator:read:chat_settings",
	"moderator:manage:chat_settings",
	"moderator:read:chatters",
	"moderator:read:shoutouts",
	"moderator:manage:shoutouts",
	"channel:moderate",
	"chat:edit",
	"chat:read",
}

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

	tokenIsValid, _, err := client.ValidateToken(os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	if !tokenIsValid {
		fmt.Println("Please insert this token in the .env file (ACCESS_TOKEN field)")
		resp, err := client.RequestAppAccessToken(scopes)
		if err != nil {
			// handle error
		}

		fmt.Printf("%+v\n", resp.Data.AccessToken)
		return
	}

	getUsers()
}

func getUsers() {
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
