package webhook

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gempir/go-twitch-irc/v4"
)

func SendMessage(message twitch.PrivateMessage, url string) error {

	contentMsg := map[string]string{
		"content":    message.Message,
		"username":   message.User.Name + " da Twitch",
		"avatar_url": "https://media.discordapp.net/attachments/1187477693726478477/1194333111480225813/chapeu_NFT-2.png?ex=65aff863&is=659d8363&hm=808df8c5413e1db83d99da0bba9ca5d0e47ba7766e832112b46cbda399810125&=&format=webp&quality=lossless&width=670&height=670",
	}

	msg, err := json.Marshal(contentMsg)
	if err != nil {
		return err
	}

	payload := bytes.NewBuffer(msg)
	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	return nil
}
