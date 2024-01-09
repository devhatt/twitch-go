package client

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v4"
)

type Twitch struct {
	Client *twitch.Client
}

func NewTwitchClient() *Twitch {
	return &Twitch{
		Client: twitch.NewAnonymousClient(),
	}
}

func (t *Twitch) JoinChannel(c string) {
	t.Client.Join(c)
	fmt.Println("Joining the channel: " + c)
}

func (t *Twitch) OnMessage(cb func(message twitch.PrivateMessage)) {
	t.Client.OnPrivateMessage(cb)
}

func (t *Twitch) Connect() error {
	fmt.Println("-------Conectado-------")
	fmt.Println("Para sair pressione CTRL + C")
	err := t.Client.Connect()

	if err != nil {
		return err
	}

	return nil
}
