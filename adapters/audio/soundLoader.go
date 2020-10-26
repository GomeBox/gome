package audio

type SoundLoader interface {
	Load(fileName string) (Sound, error)
}
