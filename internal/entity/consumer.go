package entity

import "time"

type Consumer struct {
	Id           int       `json:"id"`
	Nik          string    `json:"nik"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Gaji         int       `json:"gaji"`
	FotoKtp      string    `json:"foto_ktp"`
	FotoSelfie   string    `json:"foto_selfie"`
	IsVerified   bool      `json:"is_verified"`
	UserId       int       `json:"user_id"`
}

type UpdateConsumerPayload struct {
	Nik          string `json:"nik"`
	FullName     string `json:"full_name"`
	LegalName    string `json:"legal_name"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Gaji         int    `json:"gaji"`
}

type ConsumerActionPayload struct {
	NomorKontrak string
	Action       string
}

type ConsumerId struct {
	Id int
}
