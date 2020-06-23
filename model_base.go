package identity

// Pagination Pagination
type Pagination struct {
	Offset uint64 `json:"offset,omitempty"`
	Page   uint64 `json:"page,omitempty"`
	Limit  uint64 `json:"limit,omitempty"`
	Total  uint64 `json:"total,omitempty"`
	Pages  uint64 `json:"pages,omitempty"`
}
