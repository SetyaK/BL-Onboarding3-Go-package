package ministore

import (
	"github.com/gocraft/dbr"
)

// Product object
type Product struct {
	ProductID   int64
	Name        string
	Description string
	Stock       int
}

// ProductRepository handler
type ProductRepository struct {
	Session *dbr.Session
}

// GetAll return []Product and count
func (pr ProductRepository) GetAll() ([]Product, int, error) {
	var result []Product
	count, err := pr.Session.Select("*").From("products").LoadStructs(&result)
	return result, count, err
}

// GetByID return product by id
func (pr ProductRepository) GetByID(ID int64) (Product, error) {
	var result Product
	_, err := pr.Session.Select("*").From("products").Where("product_id = ?", ID).Load(&result)
	return result, err
}

// Add insert new product return inserted ID
func (pr ProductRepository) Add(name string, description string, initialStock int) (int64, error) {
	result, err := pr.Session.InsertInto("products").
		Columns("name", "description", "stock").
		Values(name, description, initialStock).Exec()
	if err != nil {
		return 0, err
	}

	lastID, _ := result.LastInsertId()
	return lastID, nil
}

// Update the product
func (pr ProductRepository) Update(p *Product) (bool, error) {
	result, err := pr.Session.Update("products").Set("name", p.Name).
		Set("description", p.Description).
		Where("product_id = ?", p.ProductID).Exec()
	if err != nil {
		return false, err
	}
	row, _ := result.RowsAffected()
	return row > 0, nil
}

// Delete the product by ID
func (pr ProductRepository) Delete(ID int64) (bool, error) {
	result, err := pr.Session.DeleteFrom("products").Where("product_id = ?", ID).Exec()
	if err != nil {
		return false, err
	}
	row, _ := result.RowsAffected()
	return row > 0, nil
}
