package mocks

import "github.com/GomeBox/gome/adapters/graphics"

type Settings struct {
}

func (s *Settings) WindowSettings() *graphics.WindowSettings {
	panic("implement me")
}
