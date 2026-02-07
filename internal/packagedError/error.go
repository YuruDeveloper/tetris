package packagederror

import "fmt"

var _ error = (*PackagedError)(nil)

type ErrorCode uint

const (
	FailGLFWInitError = 001
	FailGLInitError = 002
	FailCreateWindow = 003
	
	FailReadFile = 100	

	ShaderCompileFail = 200
	ProgramLinkFail = 201

	AlreadyCompiled = 300
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
