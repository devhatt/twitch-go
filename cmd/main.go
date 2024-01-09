package main

import (
	"fmt"
	"hatbot-twitch/client"
	"hatbot-twitch/configs"
	"hatbot-twitch/helpers"
	"hatbot-twitch/webhook"

	"github.com/gempir/go-twitch-irc/v4"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	t := client.NewTwitchClient()
	urlWebhook := cfg.DSCWebHookOctopost

	t.OnMessage(func(message twitch.PrivateMessage) {
		if message.Message == "-pet" && helpers.ValidUsers(message.User.Name) {
			urlWebhook = cfg.DSCWebHookPetDex
			fmt.Println("WebHook alterado para PetDex")
		}

		if message.Message == "-octo" && helpers.ValidUsers(message.User.Name) {
			urlWebhook = cfg.DSCWebHookOctopost
			fmt.Println("WebHook alterado para Octopost")
		}

		err = webhook.SendMessage(message, urlWebhook)
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("Usando Webhook da Octopost")
	t.JoinChannel("devhatt")

	err = t.Client.Connect()
	if err != nil {
		panic(err)
	}
}
