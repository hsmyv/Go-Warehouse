package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/hsmyv/go-warehouse/pkg/utils"
	"github.com/hsmyv/go-warehouse/pkg/models"
	
)

var NewProduct models.Product

func GetProduct(w http.ResponseWriter, r *http.Request){
	newProducts := models.GetAllProducts()
	res, _ :=json.Marshal(newProducts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetproductById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	productDetails, _ := models.GetproductById(ID)
	res, _ :=json.Marshal(productDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateProduct(w http.ResponseWriter, r *http.Request){
	CreateProduct := &models.Product{}
	utils.ParseBody(r, CreateProduct)
	b := CreateProduct.CreateProduct()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	product := models.DeleteProduct(ID)
	res, _ := json.Marshal(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	var UpdateProduct = &models.Product{}
	utils.ParseBody(r, UpdateProduct)
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	productDetails, db:=models.GetproductById(ID)

	if UpdateProduct.Name != ""{
		productDetails.Name = UpdateProduct.Name
	}
	if UpdateProduct.Price_buy != ""{
		productDetails.Price_buy = UpdateProduct.Price_buy
	}
	if UpdateProduct.Name != ""{
		productDetails.Price_buy = UpdateProduct.Price_buy
	}

	db.Save(&productDetails)

	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}