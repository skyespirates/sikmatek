package entity

type Installment struct {
	ID            int    `json:"id"`
	NomorKontrak  string `json:"string"`
	BulanKe       int    `json:"bulan_ke"`
	JumlahTagihan int    `json:"jumlah_tagihan"`
	DueDate       string `json:"due_date"`
	Status        string `json:"status"`
	PaidAt        string `json:"paid_at"`
}
