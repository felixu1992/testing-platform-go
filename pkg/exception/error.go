package exception

import (
	"fmt"
)

type TestingPlatformError struct {
	error
	Args      []any
	ErrorInfo ErrorInfo
}

func PlatformError(errorInfo ErrorInfo, args ...any) error {
	err := &errorInfo
	if args != nil && len(args) > 0 {
		err.Message = fmt.Sprintf(err.Message, args...)
	}
	return TestingPlatformError{Args: args, ErrorInfo: *err}
}
