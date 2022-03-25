package nilai

import (
	"Tugas-15/config"
	"Tugas-15/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "nilai"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Nilai
func GetAll(ctx context.Context) ([]models.Nilai, error) {

	var sumnilai []models.Nilai

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
		var nilai models.Nilai
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Indeks,
			&nilai.Skor,
			&createdAt,
			&updatedAt,
			&nilai.MahasiswaId,
			&nilai.MataKuliahId); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		nilai.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		nilai.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		sumnilai = append(sumnilai, nilai)
	}

	return sumnilai, nil
}

// Insert Nilai
func Insert(ctx context.Context, nilai models.Nilai) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (indeks, skor, created_at, updated_at) values('%v', %d , NOW(), NOW())", table,
		nilai.Indeks,
		nilai.Skor)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Nilai
func Update(ctx context.Context, nilai models.Nilai, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set indeks ='%s', skor = %d, updated_at = NOW() where id = %s",
		table,
		nilai.Indeks,
		nilai.Skor,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Nilai
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
