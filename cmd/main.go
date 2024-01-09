package main

import (
	"hatbot-twitch/client"
	"hatbot-twitch/configs"
	"hatbot-twitch/helpers"
	"hatbot-twitch/webhook"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	t := client.NewTwitchClient()
	helpers.GetBearerToken()

	t.OnMessage(func(message twitch.PrivateMessage) {
		var urlWebhook string
		now := time.Now().Hour()
		if now >= 14 && now <= 15 {
			urlWebhook = cfg.DSCWebHookOctopost
		}

		if now >= 16 && now <= 17 {
			urlWebhook = cfg.DSCWebHookPetDex
		}

		err = webhook.SendMessage(message, urlWebhook)
		if err != nil {
			panic(err)
		}
	})

	t.JoinChannel("devhatt")

	err = t.Client.Connect()
	if err != nil {
		panic(err)
	}
}
