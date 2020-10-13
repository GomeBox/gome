package input

//Controller gives access to a connected gamecontroller
type Controller interface {
	ButtonPressed(buttonType ButtonType) (bool, error)
}

//ButtonType defines a controller button
type ButtonType int

//goland:noinspection GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst,GoUnusedConst
const (
	ButtonA ButtonType = iota
	ButtonB
	ButtonX
	ButtonY
	ButtonStart
	ButtonSelect
)
