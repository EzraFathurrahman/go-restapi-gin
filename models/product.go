package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NameProduct string `gorm:"type:varchar(300)"  json:"nama_product"`
	Deskripsi   string `gorm:"type:text"  json:"deskripsi"`
}
