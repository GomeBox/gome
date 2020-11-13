package interfaces

type System interface {
	LoadSound(fileName string) (Player, error)
	LoadSong(fileName string) (Player, error)
}
