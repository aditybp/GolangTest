package models

type User struct {
	Id       uint   `json:"Id"`
	Email    string `json:"Email"`
	Nama     string `json:"Nama" `
	Password []byte `json:"Password"`
}
