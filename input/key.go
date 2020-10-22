package input

import "github.com/GomeBox/gome/primitives"

type Key interface {
	KeyType() primitives.KeyType
	IsPressed() bool
	WasPressed() bool
}
