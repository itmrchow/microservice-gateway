package errors

type BadRequestErr struct {
	errCode BadRequestErrCode
}
type BadRequestErrCode string

func (e BadRequestErr) Error() string {
	return string(e.errCode)
}

func NewBadRequestErr(errCode BadRequestErrCode) BadRequestErr {
	return BadRequestErr{
		errCode: errCode,
	}
}

const (
	InvalidInputDataErrCode BadRequestErrCode = "BAD_REQUEST_INVALID_INPUT_DATA"
)
