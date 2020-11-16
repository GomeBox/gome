package input

import "github.com/GomeBox/gome/interfaces"

type System interface {
	interfaces.Input
	Update() error
}
