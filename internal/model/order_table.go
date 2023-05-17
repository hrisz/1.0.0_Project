package orderTable

type Order_Table struct {
	OrderId    int    `gorm:"primaryKey;column:order_id;autoIncrement" json:"order_id"`
	Nama       string `gorm:"column:nama" json:"nama"`
	Tanggal    string `gorm:"column:tanggal" json:"tanggal"`
	Barang     string `gorm:"column:barang" json:"barang"`
	Lokasi     string `gorm:"column:lokasi" json:"lokasi"`
	Pembayaran string `gorm:"column:pembayaran" json:"pembayaran"`
}
