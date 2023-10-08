package test

import (
	"belajar-api/app"
	"belajar-api/controller"
	"belajar-api/helper"
	"belajar-api/middleware"
	"belajar-api/model/domain"
	"belajar-api/repository"
	"belajar-api/service"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_go?parseTime=true")
	//db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_testing_golang_api?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	return middleware.NewAuthMiddleware(router)
}

func truncateCategoryTable(db *sql.DB) {
	db.Exec("TRUNCATE TABLE categories")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategoryTable(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"Handphone"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", "RAHASIA")
	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["status"])
	//dirubah menjadi map[string]interface{} karena data yang dikembalikan adalah object
	assert.Equal(t, "Handphone", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategoryTable(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", "RAHASIA")
	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategoryTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget 2"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Gadget 2", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategoryTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "Gadget",
	})
	tx.Commit()

	//assert.Equal(t, 1, category.Id)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {

}

func TestGetCategoryFailed(t *testing.T) {

}

func TestDeleteCategorySuccess(t *testing.T) {

}

func TestDeleteCategoryFailed(t *testing.T) {

}

func TestListCategoriesSuccess(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {

}
