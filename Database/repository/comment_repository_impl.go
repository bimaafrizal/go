package repository

import (
	"belajar-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentReposirotyImpl struct {
	DB *sql.DB
}

// membuat implementasi new repository
func NewCommentRepository(db *sql.DB) CommentsRepository {
	return &commentReposirotyImpl{DB: db}
}

func (repo *commentReposirotyImpl) Insert(ctx context.Context, comment entity.Comments) (entity.Comments, error) {
	script := "INSERT INTO comments(email, comment) values (?, ?)"
	result, err := repo.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repo *commentReposirotyImpl) FindById(ctx context.Context, id int32) (entity.Comments, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, script, id)
	comment := entity.Comments{}
	if err != nil {
		return entity.Comments{}, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repo *commentReposirotyImpl) FindAll(ctx context.Context) ([]entity.Comments, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comments
	for rows.Next() {
		comment := entity.Comments{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repo *commentReposirotyImpl) Update(ctx context.Context, comment entity.Comments, id int32) (entity.Comments, error) {
	script := "UPDATE comments SET comment = ? WHERE id = ?"
	_, err := repo.DB.ExecContext(ctx, script, comment.Comment, id)
	if err != nil {
		return comment, err
	}
	return comment, nil
}
