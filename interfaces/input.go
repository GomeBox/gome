package interfaces

//Input gives access to the input devices
type Input interface {
	//Keyboard returns the Keyboard
	Keyboard() Keyboard
}
