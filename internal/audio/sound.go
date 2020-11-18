package audio

import adapters "github.com/GomeBox/gome/adapters/audio"

type sound struct {
	soundPlayer adapters.Sound
}

func (sound *sound) Play() error {
	return sound.soundPlayer.Play()
}

func (sound *sound) Unload() error {
	return nil
}

func (sound *sound) Unloaded() bool {
	return false
}
