package error

import (
	"fmt"
	"net/http"
)

type Op string

type Kind int

const (
	KindNotFound      = http.StatusNotFound
	KindBadRequest    = http.StatusBadRequest
	KindInternalError = http.StatusInternalServerError
	KindUnexpected    = http.StatusInternalServerError
)

type Error struct {
	Err  error
	Op   Op
	Kind Kind
}

func (e *Error) Error() string {
	return fmt.Sprintf("operation %s: err %v", e.Op, e.Err)
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch v := arg.(type) {
		case error:
			e.Err = v
		case Op:
			e.Op = v
		case Kind:
			e.Kind = v
		default:
			panic("bad call to E")
		}
	}
	return e
}

func GetKind(err error) Kind {
	e, ok := err.(*Error)
	if !ok {
		return KindUnexpected
	}
	if e.Kind != 0 {
		return e.Kind
	}
	return GetKind(e.Err)
}
