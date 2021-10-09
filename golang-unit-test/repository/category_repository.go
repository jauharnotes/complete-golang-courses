package repository

import entity "golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}