package models

import (
  "time"
)

type (
  // Mahasiswa
  Mahasiswa struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"matakuliah"`
	Nilai       uint   `json:"nilai"`
	IndeksNilai string `json:"indeks"`
	CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
)