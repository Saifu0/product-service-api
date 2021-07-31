package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	Createdon   string  `json:"-"`
	Updatedon   string  `json:"-"`
	Deletedon   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return ProductList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, idx, err := findProduct(id)

	if err != nil {
		return err
	}

	p.ID = id
	ProductList[idx] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for idx, p := range ProductList {
		if p.ID == id {
			return p, idx, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lastId := ProductList[len(ProductList)-1].ID
	return lastId + 1
}

var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		Createdon:   time.Now().UTC().String(),
		Updatedon:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "def123",
		Createdon:   time.Now().UTC().String(),
		Updatedon:   time.Now().UTC().String(),
	},
}
