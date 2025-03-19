package pagination

type SearchKey struct {
	Key   string `json:"key" form:"key"`
	Value string `json:"value" form:"value"`
}

type Request[T any] struct {
	Page   int       `json:"page" form:"page"`
	Size   int       `json:"size" form:"size"`
	Filter *T        `json:"filter" form:"filter"`
	Search SearchKey `json:"search" form:"search"`
}
