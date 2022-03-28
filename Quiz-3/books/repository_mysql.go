package books

import (
	"Quiz-3/config"
	"Quiz-3/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "books"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Books
func GetAll(ctx context.Context, params ...string) ([]models.Books, error) {
	sql := "SELECT * FROM " + table

	for i := 0; i < len(params); i++ {
		if params[i] != "" {
			sql += " WHERE"
			break
		}

		if i == len(params)-1 && params[i] == "" {
			sql += " Order By id DESC"
		}
	}

	if params[0] != "" {
		sql += " title = " + params[0]
	}
	if params[1] != "" {
		sql += ", release_year >= " + params[1]
	}
	if params[2] != "" {
		sql += ", release_year <= " + params[2]
	}
	if params[3] != "" {
		sql += ", total_page >= " + params[3]
	}
	if params[4] != "" {
		sql += ", total_page <= " + params[4]
	}
	if params[5] != "" && params[5] == "ASC" {
		sql += ", ORDER BY title " + params[5]
	}
	if params[5] != "" && params[5] == "DESC" {
		sql += ", ORDER BY title " + params[5]
	}

	var books []models.Books

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf(sql)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Books
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&book.ID,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_Year,
			&book.Price,
			&book.Total_Page,
			&book.Thickness,
			&createdAt,
			&updatedAt,
			&book.Category_ID); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	return books, nil
}

// Insert Books
func Insert(ctx context.Context, book models.Books) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) values('%v', '%v', '%v', %d,'%v', %d, '%v', NOW(), NOW(), '%v')", table,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_Year,
		book.Price,
		book.Total_Page,
		book.Thickness,
		book.Category_ID)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Books
func Update(ctx context.Context, book models.Books, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set title ='%s', description = %s, image_url = %s, release_year = %d, price = %s, total_page = %d, thickness = %s, updated_at = NOW(), category_id = %d where id = %s",
		table,
		book.Title,
		book.Description,
		book.Image_url,
		book.Release_Year,
		book.Price,
		book.Total_Page,
		book.Thickness,
		book.Category_ID,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Books
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

// Get Books by Category
func GetBooks(ctx context.Context, id string, params ...string) ([]models.Books, error) {
	sql := "SELECT * FROM " + table + " WHERE category_id = " + id
	if params[0] != "" {
		sql += ", title = " + params[0]
	}
	if params[1] != "" {
		sql += ", release_year >= " + params[1]
	}
	if params[2] != "" {
		sql += ", release_year <= " + params[2]
	}
	if params[3] != "" {
		sql += ", total_page >= " + params[3]
	}
	if params[4] != "" {
		sql += ", total_page <= " + params[4]
	}
	if params[5] != "" && params[5] == "ASC" {
		sql += ", ORDER BY title " + params[5]
	}
	if params[5] != "" && params[5] == "DESC" {
		sql += ", ORDER BY title " + params[5]
	}
	var books []models.Books

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}
	queryText := fmt.Sprintf(sql)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Books
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&book.ID,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_Year,
			&book.Price,
			&book.Total_Page,
			&book.Thickness,
			&createdAt,
			&updatedAt,
			&book.Category_ID); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		book.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		book.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	return books, nil
}
