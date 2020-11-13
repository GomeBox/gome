package interfaces

//Audio contains functions to access the audio system
type Audio interface {
	//LoadSound loads a sound effect from a file
	LoadSound(fileName string) (Player, error)
	//LoadSong loads music from a file
	LoadSong(fileName string) (Player, error)
}
