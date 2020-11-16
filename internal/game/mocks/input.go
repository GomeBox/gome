package mocks

import "github.com/GomeBox/gome/interfaces"

type Input struct {
	OnUpdate      func() error
	CallCntUpdate int
}

func (i *Input) Update() error {
	i.CallCntUpdate++
	if i.OnUpdate != nil {
		return i.OnUpdate()
	}
	return nil
}

func (i *Input) Keyboard() interfaces.Keyboard {
	panic("implement me")
}
