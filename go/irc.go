package main

import (
	"fmt"
	"os"

	"github.com/gempir/go-twitch-irc/v4"
)

func startIRCClient() {
	fmt.Println("Initialising IRC Client")
	ircClient := initIRCClient()
	fmt.Println("Setting handlers")
	setHandlers(ircClient)
	fmt.Println("Connecting")
	err := ircClient.Connect()
	if err != nil {
		panic(err)
	}
}

func setHandlers(ircClient *twitch.Client) {
	setConnectHandler(ircClient)
	setOnPrivateMessageHandler(ircClient)
}

func initIRCClient() *twitch.Client {
	ircClient := twitch.NewClient(chatterName, fmt.Sprintf("oauth:"+os.Getenv("USER_ACCESS_TOKEN")))
	// ircClient.IrcAddress = "irc-ws.chat.twitch.tv:443"
	ircClient.Join(broadcasterName)
	return ircClient
}

func setConnectHandler(ircClient *twitch.Client) {
	ircClient.OnConnect(func() {
		fmt.Println("Connected")
		ircClient.Say(broadcasterName, "Hi, I'm "+chatterName)
	})
}

func setOnPrivateMessageHandler(ircClient *twitch.Client) {
	ircClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.User.DisplayName, ":", message.Message)
	})
}
