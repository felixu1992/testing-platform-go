package exception

import (
	"fmt"
	"github.com/pkg/errors"
)

type TestingPlatformError struct {
	error
	Args      []any
	ErrorInfo ErrorInfo
}

func PlatformError(errorInfo ErrorInfo, args ...any) error {
	return errors.Wrap(TestingPlatformError{Args: args, ErrorInfo: errorInfo}, fmt.Sprintf(errorInfo.Message, args...))
}
