package shared

type Unloader interface {
	//Unload unloads the loaded resources
	Unload() error
}
