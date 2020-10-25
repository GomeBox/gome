package audio

type System interface {
	LoadSound(fileName string) (Sound, error)
	LoadSong(fileName string) (Song, error)
}
