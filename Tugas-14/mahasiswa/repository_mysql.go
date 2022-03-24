package mahasiswa

import (
	"Tugas-14/config"
	"Tugas-14/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Mahasiswa, error) {
	var mhs []models.Mahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var temp models.Mahasiswa
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&temp.ID,
			&temp.Nama,
			&temp.MataKuliah,
			&temp.Nilai,
			&temp.IndeksNilai,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		temp.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		temp.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		mhs = append(mhs, temp)
	}
	return mhs, nil
}

// Insert Mahasiswa
func Insert(ctx context.Context, mhs models.Mahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama, matakuliah, nilai, indeks, created_at, updated_at) values('%v','%v', %v,'%v', NOW(), NOW())", table,
		mhs.Nama,
		mhs.MataKuliah,
		mhs.Nilai,
		mhs.IndeksNilai)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Mahasiswa
func Update(ctx context.Context, mhs models.Mahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama ='%s', matakuliah = '%s', nilai = %d, indeks = '%s', updated_at = NOW() where id = %s",
		table,
		mhs.Nama,
		mhs.MataKuliah,
		mhs.Nilai,
		mhs.IndeksNilai,
		id,
	)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

// Delete Movie
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
