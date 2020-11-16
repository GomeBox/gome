package graphics

import (
	"github.com/GomeBox/gome/interfaces"
	"github.com/GomeBox/gome/primitives"
)

type SystemMock struct {
	CallCntClear   int
	CallCntPresent int
}

func (s *SystemMock) LoadTexture(fileName string) (interfaces.Texture, error) {
	panic("implement me")
}

func (s *SystemMock) LoadFont(fileName string, size int) (interfaces.Font, error) {
	panic("implement me")
}

func (s *SystemMock) CreateTexture(dimensions primitives.Dimensions, color primitives.Color) (interfaces.Texture, error) {
	panic("implement me")
}

func (s *SystemMock) Window() interfaces.Window {
	panic("implement me")
}

func (s *SystemMock) OpenWindow(settings WindowSettings) error {
	panic("implement me")
}

func (s *SystemMock) Clear() error {
	s.CallCntClear++
	return nil
}

func (s *SystemMock) Present() error {
	s.CallCntPresent++
	return nil
}
