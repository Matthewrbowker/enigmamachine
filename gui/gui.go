package gui

type GUI struct {
}

func New() GUI {
	return GUI{}
}

func (g GUI) Exec() int {
	return 0
}
