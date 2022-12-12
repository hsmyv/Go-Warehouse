package models

import(
	"github.com/jinzhu/gorm"
	"github.com/hsmyv/go-warehouse/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:"" json:"name"`
	Price_buy string `json:"price_buy"`
	Price_sold string `json:"price_sold"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}