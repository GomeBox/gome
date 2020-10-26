package audio

type SongLoader interface {
	Load(fileName string) (Song, error)
}
