package tableDB

//Table orders from database teknidb
type Order_Table struct {
	OrderId   int    `gorm:"primaryKey;column:order_id;autoIncrement" json:"order_id"`
	Tanggal   string `gorm:"column:tanggal" json:"tanggal"`
	Nama      string `gorm:"column:nama" json:"nama"`
	NomorHp   string `gorm:"column:nomor_hp" json:"nomor_hp"`
	Barang    string `gorm:"column:barang" json:"barang"`
	Deskripsi string `gorm:"column:desc_masalah" json:"desc_masalah"`
	Lokasi    string `gorm:"column:lokasi" json:"lokasi"`
}

//Table history from database teknidb
type Order_History struct {
	OrderId   int    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Tanggal   string `gorm:"column:tanggal" json:"tanggal"`
	Nama      string `gorm:"column:nama" json:"nama"`
	NomorHp   string `gorm:"column:nomor_hp" json:"nomor_hp"`
	Barang    string `gorm:"column:barang" json:"barang"`
	Deskripsi string `gorm:"column:desc_masalah" json:"desc_masalah"`
	Lokasi    string `gorm:"column:lokasi" json:"lokasi"`
}