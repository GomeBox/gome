package input

type Controller interface {
	ButtonPressed(buttonType ButtonType) (bool, error)
}

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
