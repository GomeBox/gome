package input

type Interface interface {
	Keyboard() *Keyboard
	ControllerCount() int
	Controller(number int) (*Controller, error)
}
