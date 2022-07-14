package chessgo

type Error interface{}

type ErrorFriendlyFire struct {
	msg string
}

func (e ErrorFriendlyFire) Error() string {
	return e.msg
}
