package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hatbot-twitch/configs"

	"io"
	"net/http"
)

type UserResponse struct {
	Data []Data `json:"data"`
}

type Data struct {
	Id              string `json:"id"`
	ProfileImageUrl string `json:"profile_image_url"`
	Login           string `json:"login"`
}

func GetUserImage(username string) (string, error) {
	buf := bytes.NewBuffer([]byte{})
	url := "https://api.twitch.tv/helix/users?login=" + username
	req, err := http.NewRequest("GET", url, buf)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+Tkn.AccessToken)
	req.Header.Add("Client-Id", configs.GetConfig().TwTClientID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var u UserResponse

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	json.Unmarshal(body, &u)

	return u.Data[0].ProfileImageUrl, nil
}
