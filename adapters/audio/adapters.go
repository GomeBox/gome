package audio

type Adapters interface {
	SoundLoader() SoundLoader
	SongLoader() SongLoader
}
