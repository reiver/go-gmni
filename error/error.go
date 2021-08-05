package geminierror

type Error interface {
	error
	GeminiReponseStatusCode() int
	GeminiReponseStatusName() string
}
