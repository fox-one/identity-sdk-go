package identity

// Pagination Pagination
type Pagination struct {
	Offset uint64 `json:"offset,omitempty"`
	Page   uint64 `json:"page,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`
	Total  uint64 `json:"total,omitempty"`
	Pages  uint64 `json:"pages,omitempty"`
}

// AppError AppError
type AppError struct {
	ErrorCode int    `json:"code"`
	ErrorMsg  string `json:"message"`
}

// Error Error
func (err *AppError) Error() string {
	return err.ErrorMsg
}

// NewAppError NewAppError
func NewAppError(msg string) *AppError {
	return &AppError{
		ErrorCode: 5000,
		ErrorMsg:  msg,
	}
}
