package schema

type ListResult struct {
	Page  int64         `json:"page"`
	Size  int64         `json:"size"`
	Total int64         `json:"total"`
	Data  []interface{} `json:"data"`
}
