package interfaces

//Context is passed to Game.Update and Game.Draw to access the game's systems, etc.
type Context interface {
	//Graphics returns the graphics adapter
	Graphics() Graphics
	//Graphics returns the input adapter
	Input() Input
	//Audio returns the audio adapter
	Audio() Audio
}
