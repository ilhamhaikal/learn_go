package repository

import (
	belajargolangdatabase "belajar_golang_database"
	"belajar_golang_database/entity"
	"context"
	"fmt"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repo@test.com",
		Comment: "ini adalah comment dari repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		fmt.Println("Error inserting comment:", err)
	}

	fmt.Println("Successfully inserted comment with ID:", result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comment, err := commentRepository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Comment found:", comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println("Comment found:", comment)
	}
}
