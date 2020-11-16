package mocks

import "github.com/GomeBox/gome/interfaces"

type Game struct {
}

func (g Game) Setup(settings interfaces.Settings) {

}

func (g Game) Initialize(context interfaces.Context) error {
	return nil
}

func (g Game) Update(timeDelta float32, context interfaces.Context) (keepRunning bool, error error) {
	return true, nil
}

func (g Game) Draw(timeDelta float32, context interfaces.Context) error {
	return nil
}
