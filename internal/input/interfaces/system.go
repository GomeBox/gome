package interfaces

type System interface {
	Keyboard() Keyboard
	Update() error
}
