package audio

import adapters "github.com/GomeBox/gome/adapters/audio"

type sound struct {
	soundPlayer adapters.Sound
}

func (sound *sound) Play() error {
	return sound.soundPlayer.Play()
}
