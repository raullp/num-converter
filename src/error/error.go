package error

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
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
	Err      error
	Op       Op
	Kind     Kind
	Severity log.Level
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
		case log.Level:
			e.Severity = v
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

func SystemError(err error) {
	e, ok := err.(*Error)
	if !ok {
		log.Error(err)
	}
	entry := log.WithFields(log.Fields{
		"operations": e.Op,
		"kind":       e.Kind,
	})

	switch e.Severity {
	case log.PanicLevel:
		entry.Panicf("%s: %v", e.Op, err)
	case log.ErrorLevel:
		entry.Errorf("%s: %v", e.Op, err)
	case log.FatalLevel:
		entry.Fatalf("%s: %v", e.Op, err)
	case log.WarnLevel:
		entry.Warnf("%s: %v", e.Op, err)
	case log.InfoLevel:
		entry.Infof("%s: %v", e.Op, err)
	case log.DebugLevel:
		entry.Debugf("%s: %v", e.Op, err)
	case log.TraceLevel:
		entry.Tracef("%s: %v", e.Op, err)
	}
}
