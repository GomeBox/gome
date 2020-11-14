package interfaces

//Settings gives access to the game's settings
type Settings interface {
	//WindowSettings contains settings about the game window
	WindowSettings() WindowSettings
}
