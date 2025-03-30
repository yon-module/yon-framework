package database

import "gorm.io/gorm"

type GpaRepository[T interface{}] struct {
	model T
	db    *gorm.DB
}

func NewGpaRepository[T any](model T) *GpaRepository[T] {
	return &GpaRepository[T]{
		model: model,
		db:    GetDB(),
	}
}

func (g *GpaRepository[T]) FindById(id uint) *T {
	var result T
	g.db.Where("id = ?", id).First(&result)
	return &result
}

func (g *GpaRepository[T]) FindByIds(ids []uint) []T {
	var results []T
	g.db.Where("id IN (?)", ids).Find(&results)
	return results
}

func (g *GpaRepository[T]) DeleteById(id uint) error {
	return g.db.Where("id = ?", id).Delete(&g.model).Error
}
