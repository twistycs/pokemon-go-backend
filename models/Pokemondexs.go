package models

type PokemonDex struct {
	ID     uint   `gorm:"primaryKey;column:POKEMON_DEX_ROW_ID"`
	UserFK uint   `gorm:"column:USER_ID_FK;" json:"-"`
	User   User   `gorm:"foreignKey:UserFK"`
	Name   string `gorm:"column:NAME;not null"`
}

func (o *PokemonDex) TableName() string {
	return "pokemondex"
}
