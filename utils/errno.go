package errno

import "fmt"

var (
	// common errors
	OK                  = &Errno{Code: 0, Msg: "OK"}
	HomoErr             = &Errno{Code: 114514, Msg: "Error"}
)

type Errno struct {
	Code int
	Msg  string
	Err  error
}

func NewErrno(code int, msg string, err error) *Errno {
	return &Errno{Code: code, Msg: msg, Err: err}
}

func (e *Errno) Error() string {
	return fmt.Sprintf("errno: %d, msg: %s,err: %v", e.Code, e.Msg, e.Err)
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Msg
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.Code, typed.Msg
	default:
		return HomoErr.Code, HomoErr.Msg
	}
}
