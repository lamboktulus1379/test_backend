package repository

import (
	"test_backend_2/Model"

	"gorm.io/gorm"
)

type Strategy interface {
	Width() float64
}

type RectangleStrategy struct {
	length float64
	width  float64
}

func (rectangle RectangleStrategy) Width() float64 {
	return rectangle.length * rectangle.width
}

type SquareStrategy struct {
	width float64
}

func (square SquareStrategy) Width() float64 {
	return square.width * square.width
}

type TriangleStrategy struct {
	base   float64
	height float64
}

func (triangle TriangleStrategy) Width() float64 {
	return (triangle.base * triangle.height) / 2
}

type RealContext struct {
	Strategy Strategy
}

func (r *RealContext) SetStrategy(strategy Strategy) {
	r.Strategy = strategy
}

func (r *RealContext) ExecuteStrategy() float64 {
	return r.Strategy.Width()
}

type AreaRepository struct {
	DB *gorm.DB
}

func NewAreaRepository(DB *gorm.DB) *AreaRepository {
	return &AreaRepository{DB: DB}
}

func (_r *AreaRepository) InsertArea(param1 float64, param2 float64, t string) (err error) {
	var ar Model.Area
	_r.DB.Model(ar)
	var r RealContext
	switch t {
	case "persegi panjang":
		r = RealContext{RectangleStrategy{param1, param2}}
	case "persegi":
		r = RealContext{SquareStrategy{param1}}
	case "segitiga":
		r = RealContext{TriangleStrategy{param1, param2}}
	default:
		r = RealContext{RectangleStrategy{4, 3}}
	}
	ar.AreaValue = r.ExecuteStrategy()
	ar.AreaType = t
	err = _r.DB.Create(&ar).Error
	if err != nil {
		return err
	}
	return err
}
