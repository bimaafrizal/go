package repository

import "balajar-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
