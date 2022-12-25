package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/prateek041/microservices-with-go/data"
)

type Products struct {
	l *log.Logger
}

// used for creating a new isntance of Product list.
func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handles the Read operation.
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("This is the put method", r.URL.Path)

		// Created a regex and matched the url string to find all the matches.
		regularExpression := regexp.MustCompile(`([0-9]+)`)
		matchedStrings := regularExpression.FindAllStringSubmatch(r.URL.Path, -1)

		// checking the condtion of the patterns mateched
		if len(matchedStrings) != 1 { // to make sure only one id was passed.
			http.Error(rw, "Ivalid URI", http.StatusBadRequest)
		}
		if len(matchedStrings[0]) != 2 { // there should be two objects in the interal literal
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}
		// takes out the id from the matched regex strings
		idString := matchedStrings[0][1]
		// converting it into an integer
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Error converting string to integer", http.StatusInternalServerError)
		}
		p.UpdateProduct(id, rw, r)
	}

	// catch every other method
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// used to get all the products from the database
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	productList := data.GetData()
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Error in marshalling the data", http.StatusInternalServerError)
	}
}

// used to add a new product to the database.
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("This is the add product handler")

	newProduct := &data.Product{} // creating an empty instance of Product type

	err := newProduct.FromJSON(r.Body) // putting all that is present in r.Body, into prod.
	if err != nil {                    // handling the errors
		http.Error(rw, "Ubable to unmarshal json", http.StatusInternalServerError)
	}

	data.AddData(newProduct) // Adds the product
}

func (p *Products) UpdateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	// Now, create a new Product using the data passed in the http request.

	updatedProduct := data.Product{} // creating an empty Product
	err := updatedProduct.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Error in creating a new Product object", http.StatusInternalServerError)
	}

	err = data.UpdateProduct(id, updatedProduct) // chaning the productList based on the data passed.

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not updated", http.StatusInternalServerError)
		return
	}
}
