package repository

import (
	Database "belajar-database"
	"belajar-database/entity"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(Database.GetConnection())

	ctx := context.Background()
	comment := entity.Comments{
		Email:   "repo@gmail.com",
		Comment: "Test Comment",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(Database.GetConnection())
	comment, err := commentRepository.FindById(context.Background(), 90)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(Database.GetConnection())
	comment, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, value := range comment {
		fmt.Println(value)
	}
}

func TestUpdate(t *testing.T) {
	commentRepository := NewCommentRepository(Database.GetConnection())
	comment := entity.Comments{
		Comment: "Test Comment Update",
	}
	result, err := commentRepository.Update(context.Background(), comment, 21)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
