package packagederror

import "fmt"

var _ error = (*PackagedError)(nil)

type ErrorCode uint

const (
	ShaderCompileFail = 100
	ProgramLinkFail = 101

	AlreadyCompiled = 1001
)

func NewError(code ErrorCode,str string) *PackagedError {
	return &PackagedError{
		errorCode: code,
		errorString: str,
	}
}

type PackagedError struct {
	errorCode ErrorCode
	errorString string
}

func (instance *PackagedError) Error() string {
	return fmt.Sprintf("error[%d]  %s",instance.errorCode,instance.errorString)
}
