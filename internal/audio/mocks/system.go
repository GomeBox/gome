package mocks

import "github.com/GomeBox/gome/audio"

type System struct {
}

func (s System) LoadSound(fileName string) (audio.Sound, error) {
	panic("implement me")
}

func (s System) LoadSong(fileName string) (audio.Song, error) {
	panic("implement me")
}
