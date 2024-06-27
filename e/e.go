package e

import "fmt"

type ErrorCode int

const (
	ErrNoError ErrorCode = iota
	ErrNoDataExist
)

var code2Msg = map[ErrorCode]string{
	ErrNoError:     "NoError",
	ErrNoDataExist: "NoDataExist",
}

func (e ErrorCode) Error() string {
	msg, found := code2Msg[e]
	if !found {
		return fmt.Sprintf("ErrorCode[%d]", e)
	}
	return msg
}

func (e ErrorCode) Code() int {
	return int(e)
}
