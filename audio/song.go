package audio

//Song is used to play music
type Song interface {
	//Play plays the song once
	Play() error
}
