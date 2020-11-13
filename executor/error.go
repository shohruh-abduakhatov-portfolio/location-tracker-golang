package executor

import "errors"

type ExecutorError error

var (
	err                = ExecutorError(errors.New("Error"))
	errUnexpectedField = ExecutorError(errors.New("Field type unexcepcted"))
)

func IsExecutorError(err error) bool {
	_, ok := err.(ExecutorError)
	return ok
}
