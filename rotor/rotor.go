package rotor

type Rotor struct {
	currentPos  int
	startingPos int
}

func New(currentPos int) Rotor {
	var startingPos = currentPos
	r := Rotor{currentPos, startingPos}
	return r
}

func (r Rotor) Spin() {
	r.currentPos = r.currentPos + 1

	if r.currentPos > 25 {
		r.currentPos = 0
	}
}

func (r Rotor) GetCurrentPos() int {
	return r.currentPos
}
