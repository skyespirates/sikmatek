package entity

type Contract struct {
	NomorKontrak    string `json:"nomor_kontrak"`
	Otr             int    `json:"otr"`
	AdminFee        int    `json:"admin_fee"`
	JumlahBunga     int    `json:"jumlah_bunga"`
	Tenor           int    `json:"tenor"`
	TotalPembiayaan int    `json:"total_pembiayaan"`
	Status          string `json:"status"`
	ConsumerId      int    `json:"consumer_id"`
	ProductId       int    `json:"product_id"`
	LimitId         int    `json:"limit_id"`
}

type CreateContractPayload struct {
	ConsumerId int    `json:"consumer_id"`
	ProductId  int    `json:"product_id"`
	LimitId    int    `json:"limit_id"`
	Tenor      string `json:"tenor"`
}

type UpdateContractPayload struct {
	Status string `json:"status"`
}
