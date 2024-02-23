package database

import (
	"database/sql"

	"github.com/arraysArrais/api-products/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func newProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (pd *ProductDB) getProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, price, category_id FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close() //executa após o término da execução dos demais blocos de código

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func (pd *ProductDB) getProductById(id string) (*entity.Product, error) {
	var product entity.Product
	err := pd.db.QueryRow("SELECT id, name, price, category_id, image_url FROM products where id = ?", id).Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageURL) //queryRow para retornar um único resultado

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pd *ProductDB) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := pd.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) values(?, ?, ?, ?, ?, ?)",
		product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("select id, name, description, price, category_id, image_url from products where category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}