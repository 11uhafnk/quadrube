package main

type qerr string

// Errors ...
const (
	ErrWrongDimention qerr = "wrong dimention deep"
	ErrMove           qerr = "box cann't move to this posistion"
	ErrWrongDirection qerr = "wrong direction to rotate"
)

// Error implementation interface error
func (e qerr) Error() string {
	return string(e)
}
