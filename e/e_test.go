package e

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// test same error type
func TestOneErrorType(t *testing.T) {
	// same error type with same Msg
	errCustom := ErrNoError
	srcErr := ErrNoDataExist
	errCustom1 := fmt.Errorf("errCustom1: %w", srcErr)
	assert.True(t, errors.Is(errCustom1, srcErr))
	// same error type without same Msg
	assert.False(t, errors.Is(errCustom1, errCustom))
	// same error source
	errCustom2 := fmt.Errorf("%w: errCustom2", errCustom1)
	errCustom3 := fmt.Errorf("%w: errCustom3", srcErr)
	assert.True(t, errors.Is(errCustom2, srcErr))
	// fmt.Printf("errCustom2: %+v\n", errCustom2)
	// fmt.Printf("errCustom3: %+v\n", errCustom3)
	assert.True(t, errors.Is(errCustom3, srcErr))
}

// test Unwrap
func TestUnwrap(t *testing.T) {
	srcErr := ErrNoDataExist
	errCustom1 := fmt.Errorf("%w: errCustom1", srcErr)
	errCustom2 := fmt.Errorf("%w: errCustom2", errCustom1)
	getSrc := Unwrap(errCustom2)
	assert.Equal(t, 1, getSrc.(ErrorCode).Code())
}

// compare error code
func TestErrorCode(t *testing.T) {
	srcErr := ErrNoDataExist
	assert.Equal(t, ErrorCode(1), srcErr)
}
