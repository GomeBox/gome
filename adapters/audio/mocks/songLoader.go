package mocks

import "github.com/GomeBox/gome/adapters/audio"

type SongLoader struct {
	OnLoad func(fileName string) (audio.Song, error)
}

func (s SongLoader) Load(fileName string) (audio.Song, error) {
	if s.OnLoad != nil {
		return s.OnLoad(fileName)
	}
	return nil, nil
}
