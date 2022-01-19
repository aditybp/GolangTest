package models

type Produk struct {
	Id           uint   `json:"IdBarang"`
	NamaBarang   string `json:"NamaBarang"`
	JumlahBarang string `json:"JumlahBarang"`
	KodeBarang   string `json:"KodeBarang"`
}
