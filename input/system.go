package input

//System gives access to the input devices
type System interface {
	//Keyboard returns the Keyboard
	Keyboard() Keyboard
}
