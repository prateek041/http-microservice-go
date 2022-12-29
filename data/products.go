package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// defined a type Product
type Product struct {
	// Adding struct field tags.
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// defining a type, to attach the methods onto it. We can call it the model of out database.
type Products []*Product

// This is used to have an instance anywhere.
func GetData() Products {
	return productList
}

func AddData(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, updatedProduct *Product) error {
	// find the product to be updated
	prod, index, err := findProduct(id)

	if err != nil {
		return err
	}
	updatedProduct.ID = prod.ID
	productList[index] = updatedProduct
	return nil
}

// this function Encodes the data into the response writer
func (p *Products) ToJSON(rw http.ResponseWriter) error {
	jsonEncoder := json.NewEncoder(rw)
	return jsonEncoder.Encode(p)
}

// This function Decodes the data from request object
func (p *Product) FromJSON(r io.Reader) error {
	jsonDecoder := json.NewDecoder(r)
	return jsonDecoder.Decode(p)
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for index, prod := range productList {
		if prod.ID == id {
			return prod, index, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// This is the database
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "latte",
		Description: "frothy milky coffee",
		Price:       3,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "espresso",
		Description: "coffee",
		Price:       2,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
