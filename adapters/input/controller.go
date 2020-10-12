package input

type Controller interface {
	ButtonPressed(buttonType ButtonType) (bool, error)
}

type ButtonType int

const (
	Button_A ButtonType = iota
	Button_B
	Button_X
	Button_Y
	Button_Start
	Button_Select
)
