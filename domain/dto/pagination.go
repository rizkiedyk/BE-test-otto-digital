package dto

type Pagination struct {
	Page          int    `json:"page"`
	Limit         int    `json:"limit"`
	SortBy        string `json:"sort_by"`
	SortOrder     string `json:"sort_order"`
	FilterByKey   string `json:"filter_by_key"`
	FilterByValue string `json:"filter_by_value"`
	Total         int64  `json:"total"`
	TotalPage     int    `json:"total_page"`
}

type ReqPagination struct {
	Page          int    `json:"page"`
	Limit         int    `json:"limit"`
	SortBy        string `json:"sort_by"`
	SortOrder     string `json:"sort_order"`
	FilterByKey   string `json:"filter_by_key"`
	FilterByValue string `json:"filter_by_value"`
	Total         int64  `json:"total"`
	TotalPage     int    `json:"total_page"`
}
