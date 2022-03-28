package category

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
	table          = "category"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Category
func GetAll(ctx context.Context) ([]models.Category, error) {

	var categories []models.Category

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var category models.Category
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		category.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		category.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, category)
	}

	return categories, nil
}

// Insert Category
func Insert(ctx context.Context, category models.Category) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (name, created_at, updated_at) values('%v', NOW(), NOW())", table,
		category.Name)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Category
func Update(ctx context.Context, category models.Category, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name ='%s', updated_at = NOW() where id = %s",
		table,
		category.Name,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Category
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
