package mocks

import "github.com/GomeBox/gome/interfaces"

type Game struct {
	OnInitialize func(context interfaces.Context) error
	OnSetup      func(settings interfaces.Settings)
	OnUpdate     func(timeDelta float32, context interfaces.Context) (keepRunning bool, error error)
	OnDraw       func(timeDelta float32, context interfaces.Context) error
}

func (g *Game) Setup(settings interfaces.Settings) {
	if g.OnSetup != nil {
		g.OnSetup(settings)
	}
}

func (g *Game) Initialize(context interfaces.Context) error {
	if g.OnInitialize != nil {
		return g.OnInitialize(context)
	}
	return nil
}

func (g *Game) Update(timeDelta float32, context interfaces.Context) (keepRunning bool, error error) {
	if g.OnUpdate != nil {
		return g.OnUpdate(timeDelta, context)
	}
	return true, nil
}

func (g *Game) Draw(timeDelta float32, context interfaces.Context) error {
	if g.OnDraw != nil {
		return g.OnDraw(timeDelta, context)
	}
	return nil
}
