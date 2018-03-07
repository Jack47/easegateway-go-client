package pdu

type paginationInfo struct {
	Total int    `json:"total"`
	Page  uint32 `json:"page"`
	Limit uint32 `json:"limit"`
}
