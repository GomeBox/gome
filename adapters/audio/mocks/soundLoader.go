package mocks

import "github.com/GomeBox/gome/adapters/audio"

type SoundLoader struct {
	OnLoad func(fileName string) (audio.Sound, error)
}

func (s SoundLoader) Load(fileName string) (audio.Sound, error) {
	if s.OnLoad != nil {
		return s.OnLoad(fileName)
	}
	return nil, nil
}
