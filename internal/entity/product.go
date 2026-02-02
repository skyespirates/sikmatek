package entity

type Product struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama_produk"`
	Kategori string `json:"kategori"`
	Harga    string `json:"harga"`
}

type CreateProductPayload struct {
	Nama     string `json:"nama_produk"`
	Kategori string `json:"kategori"`
	Harga    string `json:"harga"`
}
