package products

type ListFilters struct {
	Query     string
	Category  string
	MinPrice  int // cents
	MaxPrice  int // cents
	InStock   bool
	Sort      string
	Page      int
	PageSize  int
}

type CategoryFacet struct {
	Slug  string
	Name  string
	Count int64
}

type ListResult struct {
	Items      []Product
	Total      int64
	Page       int
	PageSize   int
	Categories []CategoryFacet
}
