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
	"moderator:manage:announcements",
}
var broadcasterId = "555891"
var moderatorId = "555891"

func loadClient() *helix.Client {
	client, err := helix.NewClient(&helix.Options{
		ClientID:        os.Getenv("CLIENT_ID"),
		ClientSecret:    os.Getenv("CLIENT_SECRET"),
		AppAccessToken:  os.Getenv("ACCESS_TOKEN"),
		UserAccessToken: os.Getenv("USER_ACCESS_TOKEN"),
		RedirectURI:     os.Getenv("REDIRECT_URL"),
	})
	if err != nil {
		panic(err)
	}

	return client
}

func main() {
	checkAuthRoutine()
	getUsers()
	sendChatAnnouncement("It's working !")
}

func getUsers() {
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

func sendChatAnnouncement(message string) {
	resp, err := client.SendChatAnnouncement(&helix.SendChatAnnouncementParams{
		BroadcasterID: broadcasterId,
		ModeratorID:   moderatorId,
		Message:       message,
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
