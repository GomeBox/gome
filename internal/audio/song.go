package audio

import adapters "github.com/GomeBox/gome/adapters/audio"

type song struct {
	songPlayer adapters.Song
}

func (song *song) Play() error {
	return song.songPlayer.Play()
}
