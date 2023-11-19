package entity

type RequestError struct {
	RequestID uint32
	Err       error
}
