package errors

type InternalErr error

type InternalErrCode string

const (
	SystemUnavailableErrCode InternalErrCode = "SYSTEM_UNAVAILABLE"
)
