package database

import (
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/yon-module/yon-framework/pagination"
	"gorm.io/gorm"
)

type Param struct {
	DB       *gorm.DB
	Page     int
	Limit    int
	OrderBy  []string
	Preloads []string
	ShowSQL  bool
}

type Paginator struct {
	TotalRecord int64       `json:"totalRecord"`
	TotalPage   int         `json:"totalPage"`
	Records     interface{} `json:"records"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prevPage"`
	NextPage    int         `json:"nextPage"`
}

func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	if len(p.Preloads) > 0 {
		for _, preload := range p.Preloads {
			db = db.Preload(preload)
		}
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int64
	var offset int

	go countRecords(db, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	db.Debug().Limit(p.Limit).Offset(offset).Find(result)

	<-done

	paginator.TotalRecord = count
	paginator.Records = result
	paginator.Page = p.Page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int64) {
	db.Model(anyType).Count(count)
	done <- true
}

func FindAllPaging[T any](pageRequest pagination.Request[T], model interface{}) *Paginator {
	query := GetDB()
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
		DB:      query,
		Page:    pageRequest.Page,
		Limit:   pageRequest.Size,
		ShowSQL: true,
	}, &model)

	return pg
}
