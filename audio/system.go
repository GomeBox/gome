package audio

//System contains functions to access the audio system
type System interface {
	//LoadSound loads a sound effect from a file
	LoadSound(fileName string) (Sound, error)
	//LoadSong loads music from a file
	LoadSong(fileName string) (Song, error)
}
