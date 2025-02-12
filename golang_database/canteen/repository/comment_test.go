package repository

import (
	"context"
	"fmt"
	"kantin2/canteen/entity"
	"kantin2/canteen/internal/db/driver"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

type CommentRepository interface {
	InsertComment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindByTime(ctx context.Context, date time.Time) ([]entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
	FindByRating(ctx context.Context, rating int) ([]entity.Comment, error)
}

type CommentRepositoryImpl struct {
	//Prepare statement
	db *pgxpool.Pool
}

// FindByTime implements CommentRepository.

func NewCommentRepo(db *pgxpool.Pool) CommentRepository {
	return &CommentRepositoryImpl{db:db}
}

// FindAll implements CommentRepository.
func (c *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	findAllQuery := `SELECT * FROM komentar_toko`
	rows, err := c.db.Query(ctx, findAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		var comment entity.Comment
		if err := rows.Scan(&comment.IDKomentar, &comment.IDToko, &comment.IDPembeli, &comment.Komentar, &comment.Rating, &comment.TglKomentar); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

// FindByRating implements CommentRepository.
func (c *CommentRepositoryImpl) FindByRating(ctx context.Context, rating int) ([]entity.Comment, error) {
	findByRatingQuery := `SELECT * FROM komentar_toko WHERE rating = $1`
	rows, err := c.db.Query(ctx, findByRatingQuery, rating)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment

	for rows.Next() {
		var comment entity.Comment
		if err := rows.Scan(&comment.IDKomentar, &comment.IDToko, &comment.IDPembeli, &comment.Komentar, &comment.Rating, &comment.TglKomentar); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

// FindByTime implements CommentRepository.
func (c *CommentRepositoryImpl) FindByTime(ctx context.Context, date time.Time) ([]entity.Comment, error) {

	ByTimeQuery := `SELECT * FROM komentar_toko WHERE tgl_komentar::date = $1`
	rows, err := c.db.Query(ctx, ByTimeQuery, date.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entity.Comment

	for rows.Next() {
		var comment entity.Comment
		if err := rows.Scan(&comment.IDKomentar, &comment.IDToko, &comment.IDPembeli, &comment.Komentar, &comment.Rating, &comment.TglKomentar); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("rows error")
	}
	return comments, nil
}

// InsertComment implements CommentRepository.
func (c *CommentRepositoryImpl) InsertComment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {

	insertQuery := `INSERT INTO komentar_toko (id_toko, id_pembeli, komentar, rating, tgl_komentar) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id_komentar`
	row := c.db.QueryRow(ctx, insertQuery, comment.IDToko, comment.IDPembeli, comment.Komentar, comment.Rating, comment.TglKomentar)
	err:=row.Scan(&comment.IDKomentar)
	if err != nil {
		return entity.Comment{}, err
	}
	return comment,nil
}

func TestInsertComment(t *testing.T) {
	//setup mock pool
	driver.InitDBPool()
	defer driver.CloseDBPool()
	db:=driver.GetDbpool()

	repo:=NewCommentRepo(db)

	//Prepare expectations
	comment := entity.Comment{
		IDToko:      "qwety123!@#",
		IDPembeli:   "9c8e980c-4310-48ea-b5a9-c6aa02e9d712",
		Komentar:    "Komentar test",
		Rating:      5,
		TglKomentar: time.Now(),
	}
	result,err:=repo.InsertComment(context.Background(), comment)
	
	assert.NoError(t, err)
	assert.NotZero(t, result.IDKomentar)

}
