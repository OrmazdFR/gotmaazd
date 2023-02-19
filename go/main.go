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

func checkAuthRoutine() {
	checkUserToken()
	checkAppToken()
}

func checkUserToken() {
	userAccessToken := os.Getenv("USER_ACCESS_TOKEN")

	isValid, _, err := client.ValidateToken(userAccessToken)
	if err != nil {
		fmt.Println("Error during User Token check")
		os.Exit(0)
	}

	if !isValid {
		fmt.Println("USER TOKEN INCORRECT")
		getUserAccessToken()
		os.Exit(0)
	}
}

func getUserAccessToken() {
	code := os.Getenv("CODE")

	resp, err := client.RequestUserAccessToken(code)
	if err != nil {
		// handle error
		os.Exit(0)
	}

	if resp.StatusCode != 200 {
		checkCode()
	}

	fmt.Printf("VEUILLEZ INSÉRER CE TOKEN DANS LE .env (USER_ACCESS_TOKEN) : %+v\n", resp.Data.AccessToken)
}

func checkAppToken() {
	appTokenIsValid, _, err := client.ValidateToken(os.Getenv("ACCESS_TOKEN"))
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	if !appTokenIsValid {
		fmt.Println("Please insert this token in the .env file (ACCESS_TOKEN field)")
		resp, err := client.RequestAppAccessToken(scopes)
		if err != nil {
			// handle error
		}

		fmt.Printf("%+v\n", resp.Data.AccessToken)
		return
	}
}

func checkCode() {
	if len(os.Getenv("CODE")) <= 0 {
		fmt.Println("LA VARIABLE ENVIRONEMENT 'CODE' N'EST PAS DÉFINIE, VEUILLEZ RÉCUPÉRER LE CODE DEPUIS CE LIEN ET L'INSÉRER DANS LE .ENV :")
		getCodeURL()
		os.Exit(0)
	}

	code := os.Getenv("CODE")
	resp, err := client.RequestUserAccessToken(code)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("VARIABLE ENVIRONEMENT 'CODE' INCORRECTE, VEUILLEZ RÉCUPÉRER LE CODE DEPUIS CE LIEN ET L'INSÉRER DANS LE .ENV :")
		getCodeURL()
		os.Exit(0)
	}
	fmt.Println("Code bon")
}

func getCodeURL() {
	url := client.GetAuthorizationURL(&helix.AuthorizationURLParams{
		ResponseType: "code", // or "token"
		Scopes:       scopes,
		State:        "some-state",
		ForceVerify:  false,
	})

	fmt.Println(url)
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
