package database

import (
	"fmt"
	"github.com/yon-module/yon-framework/exception"
	"github.com/yon-module/yon-framework/logger"
	"github.com/yon-module/yon-framework/pagination"
	"github.com/yon-module/yon-framework/server/response"
	"gorm.io/gorm"
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
		if filterMap, ok := any(*pageRequest.Filter).(map[string]interface{}); ok {
			for key, value := range filterMap {
				if value == nil {
					continue
				}

				switch v := value.(type) {
				case string:
					if v == "" {
						continue
					}
					query = query.Where(fmt.Sprintf("%s = ?", key), v)

				case float64:
					query = query.Where(fmt.Sprintf("%s = ?", key), v)

				case bool:
					query = query.Where(fmt.Sprintf("%s = ?", key), v)

				case []interface{}:
					query = query.Where(fmt.Sprintf("%s IN ?", key), v)

				default:
					logger.Log.Warn().Str("key", key).Interface("value", value).Msg("Unhandled filter type")
				}
			}
		} else {
			logger.Log.Warn().Msg("Filter is not a map[string]interface{}")
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
