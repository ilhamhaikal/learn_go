package repository

import (
	"belajar_golang_database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comment(email, comment) VALUES ($1, $2) RETURNING id"
	var id int
	err := repository.DB.QueryRowContext(ctx, query, comment.Email, comment.Comment).Scan(&id)
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comment WHERE id = $1 LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comment"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		err = rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
