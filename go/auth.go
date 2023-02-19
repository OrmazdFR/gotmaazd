package main

import (
	"fmt"
	"os"

	"github.com/nicklaw5/helix/v2"
)

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
