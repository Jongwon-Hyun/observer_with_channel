package pub_sub

type Observer interface {
	Receiver(message string)
	GetName() string
	ListenOnChannel()
}
