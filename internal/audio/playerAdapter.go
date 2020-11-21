package audio

import "github.com/GomeBox/gome/adapters/shared"

type playerAdapter interface {
	shared.Unloader
	Play() error
}
