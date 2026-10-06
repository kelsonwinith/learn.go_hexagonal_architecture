package sharedDomain

// ============================================================================
// Constants
// ============================================================================

const (
	DefaultPageSize = 10
	MaxPageSize     = 100
)

// ============================================================================
// Types
// ============================================================================

type Pagination struct {
	Page     int
	PageSize int
}

type PaginationLimits struct {
	DefaultPageSize int
	MaxPageSize     int
}

type Page[T any] struct {
	Items []T
	Pagination
	Total int64
}

// ============================================================================
// Constructors
// ============================================================================

func NewPagination(page, pageSize int, limits ...PaginationLimits) Pagination {
	defaultPageSize := DefaultPageSize
	maxPageSize := MaxPageSize

	if len(limits) > 0 {
		if limits[0].DefaultPageSize > 0 {
			defaultPageSize = limits[0].DefaultPageSize
		}
		if limits[0].MaxPageSize > 0 {
			maxPageSize = limits[0].MaxPageSize
		}
	}

	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = defaultPageSize
	}

	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	return Pagination{Page: page, PageSize: pageSize}
}

// ============================================================================
// Methods
// ============================================================================

func (p Pagination) Limit() int {
	return p.PageSize
}

func (p Pagination) Offset() int {
	return (p.Page - 1) * p.PageSize
}
