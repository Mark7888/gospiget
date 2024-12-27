package gospiget

import "fmt"

// NotFoundError represents a 404 Not Found error
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not Found: %s", e.Message)
}

// UnexpectedStatusCodeError represents an unexpected status code error
type UnexpectedStatusCodeError struct {
	StatusCode int
}

func (e *UnexpectedStatusCodeError) Error() string {
	return fmt.Sprintf("unexpected status code: %d", e.StatusCode)
}

// UnmarshalError represents an error during unmarshalling
type UnmarshalError struct {
	Message string
}

func (e *UnmarshalError) Error() string {
	return fmt.Sprintf("failed to unmarshal response: %s", e.Message)
}

// RequestError represents an error during the request
type RequestError struct {
	Message string
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("request error: %s", e.Message)
}
