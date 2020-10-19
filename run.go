package gome

import (
	"errors"
	"github.com/GomeBox/gome/game"
	"github.com/GomeBox/gome/graphics"
	"github.com/GomeBox/gome/input"
	"github.com/GomeBox/gome/internal/core"
)

//Gome is used to run a Interface
type Gome interface {
	//Run starts the Interface. Returns when the game ends
	Run(game game.Interface, settings game.Settings) error
	//ChangeGameRunner is used for testing purposes only
	ChangeGameRunner(runner core.GameRunner)
}

//New returns a new instance of Gome
func New() Gome {
	g := new(gome)
	g.gameRunner = core.NewGameRunner()
	return g
}

type gome struct {
	gameRunner core.GameRunner
}

func (gome *gome) Run(game game.Interface, settings game.Settings) error {
	if gome.gameRunner.Running() {
		return errors.New("game is already running")
	}
	err := gome.gameRunner.Init(game.CreateAdapters, settings)
	if err != nil {
		return err
	}
	c := newContextWrapper(gome.gameRunner)
	update := func(timeDelta float32) error {
		return game.Update(timeDelta, c)
	}
	draw := func(timeDelta float32) error {
		return game.Draw(timeDelta, c)
	}
	initialize := func() error {
		return game.Initialize(c)
	}
	err = gome.gameRunner.Loop(initialize, update, draw)
	if err != nil {
		return err
	}
	return nil
}

func (gome *gome) ChangeGameRunner(runner core.GameRunner) {
	gome.gameRunner = runner
}

type contextWrapper struct {
	runner   core.GameRunner
	graphics graphics.System
	input    input.System
}

func newContextWrapper(runner core.GameRunner) *contextWrapper {
	context := new(contextWrapper)
	context.runner = runner
	textureLoader := runner.AdapterSystem().Graphics().TextureLoader()
	windowAdapter := runner.AdapterSystem().Graphics().WindowAdapter()
	fontLoader := runner.AdapterSystem().Graphics().FontLoader()
	context.graphics = graphics.NewSystem(
		graphics.Adapters{
			TextureLoader: textureLoader,
			WindowAdapter: windowAdapter,
			FontLoader:    fontLoader})

	return context
}

func (context *contextWrapper) QuitGame() {
	context.runner.Quit()
}

func (context *contextWrapper) Graphics() graphics.System {
	return context.graphics
}

func (context *contextWrapper) Input() input.System {
	return context.input
}
