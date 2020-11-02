package audio

import "github.com/GomeBox/gome/internal/audio/interfaces"

//Song is used to play music
type Song interface {
	//Play plays the song once
	Play() error
}

type song struct {
	internalSong interfaces.Song
}

func (s *song) Play() error {
	return s.internalSong.Play()
}
