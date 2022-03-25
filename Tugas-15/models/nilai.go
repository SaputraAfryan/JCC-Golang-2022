package models

import (
	"time"
)

type (
	// Nilai
	Nilai struct {
		ID           uint        `json:"id"`
		Indeks       string      `json:"indeks"`
		Skor         uint        `json:"skor"`
		CreatedAt    time.Time   `json:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at"`
		MahasiswaId  interface{} `json:"mahasiswa_id"`
		MataKuliahId interface{} `json:"mata_kuliah_id"`
	}
)
