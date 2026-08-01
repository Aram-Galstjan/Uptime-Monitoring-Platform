package eventbus

type Bus struct{}

func New() *Bus {
	return &Bus{}
}
