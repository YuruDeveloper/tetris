package packagederror

import "fmt"

var _ error = (*PackagedError)(nil)

type ErrorCode uint

const (
	FailReadFile = 001	

	ShaderCompileFail = 100
	ProgramLinkFail = 101

	FailGLFWInitError = 200
	FailGLInitError = 201
	FailCreateWindow = 202

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
	return fmt.Sprintf("error[%-3d]  %s",instance.errorCode,instance.errorString)
}
