package audio

import (
	"errors"
	adapters "github.com/GomeBox/gome/adapters/audio"
)

type song struct {
	songPlayer adapters.Song
	unloaded   bool
}

func (song *song) Play() error {
	return song.songPlayer.Play()
}

func (song *song) Unload() error {
	if song.unloaded {
		return errors.New("resource already unloaded")
	}
	err := song.songPlayer.Unload()
	if err != nil {
		return err
	}
	song.unloaded = true
	return nil
}

func (song *song) Unloaded() bool {
	return song.unloaded
}
