package errors

import "errors"

var (
	ErrInvalidJSON   = errors.New("invalid json response")
	ErrLLMFailure    = errors.New("llm failure")
	ErrRetryExceeded = errors.New("retry attempts exceeded")
)
