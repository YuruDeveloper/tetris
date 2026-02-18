package packagederror

import "fmt"

var _ error = (*PackagedError)(nil)

type ErrorCode uint

const (
	FailGLFWInitError = 001
	FailGLInitError   = 002
	FailCreateWindow  = 003

	FailReadFile = 100
	FailOpenFile = 101

	FailDecodeImage  = 200
	FailConvertImage = 201

	ShaderCompileFail = 300
	ProgramLinkFail   = 301

	AlreadyCompiled         = 400
	DataArrayIsEmpty        = 401
	FailCreateRenderingData = 402
	UnSupportedDataType     = 403
	AlreadyDeleted = 404

	UnknownAsset = 500
	FailAssetTypeConvert = 501
	FailRegisterDuplicate = 502
)

func NewError(code ErrorCode, str string) *PackagedError {
	return &PackagedError{
		errorCode:   code,
		errorString: str,
	}
}

type PackagedError struct {
	errorCode   ErrorCode
	errorString string
}

func (instance *PackagedError) Error() string {
	return fmt.Sprintf("error[%-3d]  %s", instance.errorCode, instance.errorString)
}
