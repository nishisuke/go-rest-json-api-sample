package db

// CHECK

import (
	"your_module/internal/app/api/err"
	wrap "your_module/internal/pkg/db"

	"gorm.io/gorm"
)

type (
	Page[T any] struct {
		Pagination int
		TotalPage  int
		HasNext    bool
		TotalCount int
		Items      []T
	}

	SequenceItems[T any] struct {
		HasNext bool
		Items   []T
	}
)

func Paging[T any](db *gorm.DB, limit int, page int, findScope func(*gorm.DB) *gorm.DB) (Page[T], error) {
	if page <= 0 {
		return Page[T]{}, err.ErrZeroPage
	}

	items, err := wrap.All[T](db, func(d *gorm.DB) *gorm.DB {
		return d.Scopes(findScope).Offset((page - 1) * limit).Limit(limit)
	})
	if err != nil {
		return Page[T]{}, err
	}

	total, err := wrap.Count[T](db, findScope)
	if err != nil {
		return Page[T]{}, err
	}

	totalPage := total / limit
	if total%limit > 0 {
		totalPage++
	}

	return Page[T]{
		Pagination: page,
		TotalCount: total,
		TotalPage:  totalPage,
		HasNext:    page != totalPage,
		Items:      items,
	}, nil
}

func SequencePaging[T any](db *gorm.DB, limit int, page int, findScope func(*gorm.DB) *gorm.DB) (SequenceItems[T], error) {
	if page <= 0 {
		return SequenceItems[T]{}, err.ErrZeroPage
	}
	plusOne := limit + 1
	items, err := wrap.All[T](db, func(d *gorm.DB) *gorm.DB {
		return d.Scopes(findScope).Offset((page - 1) * limit).Limit(plusOne)
	})

	if err != nil {
		return SequenceItems[T]{}, err
	}

	hasNext := len(items) == plusOne

	return SequenceItems[T]{
		HasNext: hasNext,
		Items:   items[:min(limit, len(items))],
	}, nil
}

func min(items, b int) int {
	if items < b {
		return items
	}
	return b
}
