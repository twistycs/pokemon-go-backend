package models

type Pokemon struct {
	ID            uint   `gorm:"primaryKey;column:POKEMON_ROW_ID"  json:"-"`
	Name          string `gorm:"column:NAME;not null"`
	ImageUrl      string `gorm:"column:IMAGE_URL;not null"`
	ImageUrlHiRes string `gorm:"column:IMAGE_URL_HI_RES;not null"`
	Hp            string `gorm:"column:HP"`
	SuperType     string `gorm:"column:SUPER_TYPE"`
	SubType       string `gorm:"column:SUB_TYPE"`
	Type          string `gorm:"column:TYPE"`
}
