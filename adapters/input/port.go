package input

//Port is to be implemented by the input adapter
type Port interface {
	//Keyboard returns the Keyboard adapter
	Keyboard() Keyboard
	//ControllerCount returns the number of connected controllers
	ControllerCount() int
	//Controller returns the specified controller
	Controller(number int) (*Controller, error)
}
