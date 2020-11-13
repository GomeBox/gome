package audio

import "github.com/GomeBox/gome/internal/audio/interfaces"

type Player interface {
	Play() error
}

type player struct {
	internalPlayer interfaces.Player
}

func (s *player) Play() error {
	return s.internalPlayer.Play()
}
