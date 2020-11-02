package audio

import "github.com/GomeBox/gome/internal/audio/interfaces"

//Sound is used to Play a sound effect
type Sound interface {
	//Play plays the sound effect once
	Play() error
}

type sound struct {
	internalSound interfaces.Sound
}

func (s *sound) Play() error {
	return s.internalSound.Play()
}
