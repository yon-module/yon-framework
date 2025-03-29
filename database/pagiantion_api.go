package database

import (
	"fmt"
	"github.com/yon-module/yon-framework/exception"
	"github.com/yon-module/yon-framework/pagination"
	"github.com/yon-module/yon-framework/server/response"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type Pagination[T any] struct {
	db       *gorm.DB
	model    interface{}
	preloads []string
	request  *pagination.Request[T]
}

func NewPagination[T any]() *Pagination[T] {
	return &Pagination[T]{
		db: GetDB(),
	}
}

func (p *Pagination[T]) SetModal(model interface{}) *Pagination[T] {
	p.model = model
	return p
}

func (p *Pagination[T]) SetPreloads(preloads ...string) *Pagination[T] {
	p.preloads = preloads
	return p
}

func (p *Pagination[T]) SetRequest(request *pagination.Request[T]) *Pagination[T] {
	p.request = request
	return p
}

func (p *Pagination[T]) FindAllPaging() *Paginator {
	query := p.db

	pageRequest := p.request
	if pageRequest == nil {
		panic(exception.NewBadRequestExceptionStruct(response.BadRequest, "Please provide Paging Request"))
	}

	// Apply search filter (if any)
	if pageRequest.Search.Key != "" && pageRequest.Search.Value != "" {
		// Here, assume we're using dynamic column filtering based on the search key
		query = query.Where(fmt.Sprintf("%s LIKE ?", pageRequest.Search.Key), "%"+pageRequest.Search.Value+"%")
	}

	// Apply filter (if any)
	// You can extend this based on the `Filter` type if necessary, for now, let's assume it's a field to search
	if pageRequest.Filter != nil {
		// Apply the filter logic here based on your structure.
		// Example: You can add more complex filtering logic here
		// E.g., if Filter is a struct with specific fields like Name, Price, etc.
		// Use reflection to dynamically check fields in Filter
		filter := reflect.ValueOf(*pageRequest.Filter)

		// Iterate over the fields in the filter struct
		for i := 0; i < filter.NumField(); i++ {
			field := filter.Type().Field(i)
			fieldValue := filter.Field(i)

			// Skip zero values (i.e., empty or zero-value fields)
			if fieldValue.IsZero() {
				continue
			}

			// Build dynamic where clause based on field name and value
			columnName := strings.ToLower(field.Name) // Adjust to match the column name
			query = query.Where(fmt.Sprintf("%s = ?", columnName), fieldValue.Interface())
		}
	}

	pg := Paging(&Param{
		DB:       query,
		Page:     pageRequest.Page,
		Limit:    pageRequest.Size,
		OrderBy:  []string{"id desc"},
		Preloads: p.preloads,
		ShowSQL:  true,
	}, &p.model)

	return pg
}
