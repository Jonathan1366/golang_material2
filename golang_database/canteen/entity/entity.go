package entity

import "time"

type Comment struct{
	IDKomentar   int64     `json:"id_komentar"`
	IDToko       string    `json:"id_toko"`
	IDPembeli    string    `json:"id_pembeli"`
	Komentar     string    `json:"komentar"`
	Rating       int       `json:"rating"`
	TglKomentar  time.Time `json:"tgl_komentar"`
}