package models

import(
	"github.com/jinzhu/gorm"
	"github.com/hsmyv/go-warehouse/pkg/config"
)

var db *gorm.DB

type Product struct{
	gorm.Model
	Name string `gorm:"" json:"name"`
	Price_buy string `json:"price_buy"`
	Price_sold string `json:"price_sold"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func (b *Product) CreateProduct() *Product{
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllProducts() []Product{
	var Products []Product
	db.Find(&Products)
	return Products
}
func GetproductById(Id int64) (*Product, *gorm.DB){
	var getProduct Product
	db:=db.Where("ID=?", Id).Find(&getProduct)
	return &getProduct, db
}
func DeleteProduct(ID int64) Product{
	var product Product
	db.Where("ID=?", ID).Delete(product)
	return product
}
