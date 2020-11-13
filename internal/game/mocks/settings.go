package mocks

import "github.com/GomeBox/gome/adapters/graphics"

type Settings struct {
	OnWindowSettings func() *graphics.WindowSettings
}

func (s *Settings) WindowSettings() *graphics.WindowSettings {
	if s.OnWindowSettings != nil {
		return s.OnWindowSettings()
	}
	return nil
}
