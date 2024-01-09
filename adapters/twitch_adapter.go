package adapters

type TwitchAdapter interface {
	JoinChannel(c string)
	Connect() error
}
