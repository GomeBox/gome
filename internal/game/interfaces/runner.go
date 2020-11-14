package interfaces

import (
	"github.com/GomeBox/gome/interfaces"
)

type Runner interface {
	Run(game interfaces.Game) error
}
