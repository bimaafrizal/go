package repository

import (
	"balajar-golang-unit-test/entity"
	"github.com/stretchr/testify/mock"
)

// mock adalah object yang sudah kita program dengan ekspektasi tertentu
// mock digunakan untuk menggantikan object asli yang sulit untuk dites ketika melakukan unit testing
// misal kita memanggil API dari luar, atau melakukan koneksi ke database
// bisa gunakan testify
// gunakan interface agar mudah di uji

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository CategoryRepositoryMock) FindById(id string) *entity.Category {
	//memanggil mock
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		//untuk mengembalikan data dari mock
		category := arguments.Get(0).(entity.Category) //akan diconvert menjadi entity.Category
		return &category
	}
}
