package errors_test

import (
	goError "errors"
	"testing"

	errors "github.com/bytheway/VetFailure"
)

func TestDemo(t *testing.T) {
	errors.Demo()
	_ = goError.New("error")
}
