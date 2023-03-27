package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/nicklaw5/helix/v2"
)

var client *helix.Client = loadHelixClient()
var scopes = []string{
	"chat:edit",
	"chat:read",
	"moderator:manage:announcements",
	"channel:read:subscriptions",
}
var broadcasterName = "ormaazd"

func loadHelixClient() *helix.Client {
	client, err := helix.NewClient(&helix.Options{
		ClientID:        os.Getenv("CLIENT_ID"),
		ClientSecret:    os.Getenv("CLIENT_SECRET"),
		AppAccessToken:  os.Getenv("ACCESS_TOKEN"),
		UserAccessToken: os.Getenv("ADMIN_ACCESS_TOKEN"),
		RedirectURI:     os.Getenv("REDIRECT_URL"),
	})
	if err != nil {
		panic(err)
	}

	return client
}

func main() {
	checkAuthRoutine()
	startIRCClient()
}

func getUserId(username string) string {
	resp, err := client.GetUsers(&helix.UsersParams{
		Logins: []string{username},
	})
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Erreur lors de la récupération des Users")
		panic(resp.ErrorStatus)
	}

	return resp.Data.Users[0].ID
}

func getUsers(channels []string) {
	resp, err := client.GetUsers(&helix.UsersParams{
		Logins: channels,
	})
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("Erreur lors de la récupération des Users")
		return
	}

	fmt.Println("Users :")
	for _, user := range resp.Data.Users {
		fmt.Printf("ID: %s Name: %s\n", user.ID, user.DisplayName)
	}
}

func sendChatAnnouncement(message string, broadcaster string, moderator string) {
	resp, err := client.SendChatAnnouncement(&helix.SendChatAnnouncementParams{
		Message:       message,
		BroadcasterID: getUserId(broadcaster),
		ModeratorID:   getUserId(moderator),
	})
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 204 {
		fmt.Println("Erreur lors de l'envoi de l'annonce : ", message)
		return
	}

	fmt.Println("Annonce envoyée : ", message)
}

func getSubscribersInfos(broadcaster string) {
	resp, err := client.GetSubscriptions(&helix.SubscriptionsParams{
		BroadcasterID: getUserId(broadcaster),
	})
	if err != nil {
		fmt.Println("Error getting subsrivers infos : ", err)
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("Not authorized to retrieve", broadcaster, "subscribers")
		return
	}

	for _, subscriber := range resp.Data.Subscriptions {
		fmt.Println(subscriber.UserName, "Tier:", subscriber.Tier[0:1])
	}
}
