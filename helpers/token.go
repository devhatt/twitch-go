package helpers

import (
	"bytes"
	"encoding/json"
	"hatbot-twitch/configs"
	"io"
	"log"
	"net/http"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

var Tkn TokenResponse

func GetBearerToken() error {
	msg := bytes.NewBuffer([]byte{})
	config := configs.GetConfig()

	resp, err := http.Post(config.TwTTokenURL, "application/json", msg)
	if err != nil {
		log.Fatalf("error on http post %s", err)
		return err
	}

	body, err := io.ReadAll(resp.Body)

	json.Unmarshal(body, &Tkn)

	if err != nil {
		log.Fatalf("error on read %v", err)
		return err
	}

	defer resp.Body.Close()
	return nil
}
