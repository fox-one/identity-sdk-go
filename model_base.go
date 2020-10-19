package identity

import "fmt"

const ErrorOk = 0

// AppError AppError
type AppError struct {
	ErrorCode int    `json:"code"`
	ErrorMsg  string `json:"message"`
}

// Error Error
func (err *AppError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", err.ErrorCode, err.ErrorMsg)
}

// NewAppError NewAppError
func NewAppError(msg string) *AppError {
	return &AppError{
		ErrorCode: 5000,
		ErrorMsg:  msg,
	}
}

// Pagination Pagination
type Pagination struct {
	Offset uint64 `json:"offset,omitempty"`
	Page   uint64 `json:"page,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`
	Total  uint64 `json:"total,omitempty"`
	Pages  uint64 `json:"pages,omitempty"`
}

// NewPagePagination NewPagination
func NewPagePagination(page, limit, total uint64) *Pagination {
	offset := (page - 1) * limit

	p := &Pagination{
		Page:   page,
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}

	return p
}

// NewOffsetPagination NewPagination
func NewOffsetPagination(offset, limit, total uint64) *Pagination {
	page := offset/limit + 1
	pages := total/limit + 1
	p := &Pagination{
		Page:   page,
		Offset: offset,
		Limit:  limit,
		Total:  total,
		Pages:  pages,
	}

	return p
}

// BasePageResponse BasePageResponse
type BasePageResponse struct {
	Items      []interface{} `json:"data"`
	Pagination *Pagination   `json:"pagination"`
	// Total      uint64        `json:"total"`
}
