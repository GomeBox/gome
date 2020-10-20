package graphics

//Adapters gives access to all graphics adapters
type Adapters interface {
	//TextureLoader returns the adapter that is used to load textures
	TextureLoader() TextureLoader
	//TextureCreator returns the adapter that is used to create simple unicolored textures
	TextureCreator() TextureCreator
	//FontLoader returns the adapter that is used to load fonts
	FontLoader() FontLoader
	//WindowAdapter returns the adapter that is used to open and close the game window
	WindowAdapter() WindowAdapter
	//ScreenPresenter returns the adapter that is used to update the screen
	ScreenPresenter() ScreenPresenter
}
