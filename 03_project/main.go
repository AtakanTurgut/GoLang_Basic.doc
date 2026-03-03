// go run main.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          string  `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           string `json:"id"`
	CategoryName string `json:"categoryName"`
	Description  string `json:"description"`
}

func GetAllProducts() ([]Product, error) { // ([]Product, error)
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err // nil, err
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)

	//fmt.Println("Products:", products)

	return products, nil
}

func AddProduct() (Product, error) { // (Product, error)
	product := Product{
		Id:          "9", // json-server otomatik id atar
		ProductName: "Soğutucu",
		CategoryId:  2,
		UnitPrice:   500.00,
	}

	jsonProduct, _ := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		fmt.Println("Error:", err)
		return Product{}, err // nil, err
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)

	//fmt.Println("Added Product:", newProduct)

	return productResponse, nil
}

func main() {
	// ADD A NEW PRODUCT
	//AddProduct()
	//GetAllProducts()

	// GET ALL PRODUCTS
	// GetAllProducts()
	products, _ := GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Printf("Id: %s, Product Name: %s, Category Id: %d, Unit Price: %.2f\n",
			products[i].Id, products[i].ProductName, products[i].CategoryId, products[i].UnitPrice)
	}

	fmt.Println()
}
