package input

//Adapters gives access to all input adapters
type Adapters interface {
	//Keyboard returns the KeyboardAdapter
	Keyboard() KeyboardAdapter
}
