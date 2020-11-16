package interfaces

//WindowSettings contains all the settings needed to show the game window
type WindowSettings interface {
	//Fullscreen defines if the game window should use the whole screen.
	//The size of the game window is the Resolution, if Fullscreen is false
	SetFullscreen(value bool) WindowSettings
	//Resolution defines the resolution of the game window
	SetResolution(width, height int) WindowSettings
	//Title is the game's title that is used as the window title
	SetTitle(value string) WindowSettings
}
