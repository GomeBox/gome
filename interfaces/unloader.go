package interfaces

type Unloader interface {
	//Unload unloads the loaded resources
	Unload() error
	//Unloaded returns true, if Unload has been called successfully
	Unloaded() bool
}
