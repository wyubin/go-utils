package e

import (
	"fmt"
)

var (
	ErrMapKeyNotExist     = fmt.Errorf("%w: key not exist in map", ErrNoDataExist)
	ErrSliceIndexNotExist = fmt.Errorf("%w: index not exist in slice", ErrNoDataExist)
	ErrRepoDBNotExist     = fmt.Errorf("%w: dbName not exist in repo", ErrNoDataExist)
	ErrRepoCollNotExist   = fmt.Errorf("%w: collName not exist in repo", ErrNoDataExist)
)

func Unwrap(err error) error {
	switch x := err.(type) {
	case interface{ Unwrap() error }:
		err = x.Unwrap()
		if err == nil {
			return nil
		}
		return Unwrap(err)
	case interface{ Error() string }:
		return err
	default:
		return nil
	}
}
