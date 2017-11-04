package ministore

import (
	"os"
	"testing"

	"github.com/SetyaK/BL-Onboarding3-Go-package/database"
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
)

func dbSessionSetup() (*dbr.Session, error) {
	os.Setenv("DATABASE_ADAPTER", "sqlite3")

	// Initialize database session
	sess, err := database.NewSession()
	if err != nil {
		return nil, err
	}

	// Migrate database schema
	m := database.Migration{Session: sess}
	_, err = m.Migrate()
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func TestAdd(t *testing.T) {
	// Init session
	sess, err := dbSessionSetup()
	if err != nil {
		t.Fatal(err)
	}
	// Test add product
	pr := ProductRepository{Session: sess}
	productName := "Test Product"
	productDescription := "The description of test product"
	productStock := 2
	newProductID, err := pr.Add(productName, productDescription, productStock)
	if err != nil {
		t.Fatal(err)
	}
	if newProductID == 0 {
		t.Fatal("Inserted product id should not zero")
	}

	// Check the product in database
	insertedProduct, err := pr.GetByID(newProductID)
	if err != nil {
		t.Fatal(err)
	}
	if insertedProduct.Name != productName {
		t.Fatal("Inserted product name doesn't match")
	}
	if insertedProduct.Description != productDescription {
		t.Fatal("Inserted product description doesn't match")
	}
	if insertedProduct.Stock != productStock {
		t.Fatal("Inserted product stock doesn't match")
	}

	// Error test for add
	pr.Session.DB.Close()
	result, err := pr.Add("", "", 0)
	if err == nil || result > 0 {
		t.Fatal("Product add should be failed if database connection was closed")
	}
}
func TestUpdate(t *testing.T) {
	// Init session
	sess, err := dbSessionSetup()
	if err != nil {
		t.Fatal(err)
	}

	// Insert product
	pr := ProductRepository{Session: sess}
	pID, err := pr.Add("p1", "d1", 1)
	if err != nil {
		t.Fatal(err)
	}

	// Get the product in database
	product, err := pr.GetByID(pID)
	if err != nil {
		t.Fatal(err)
	}

	// Update product
	productNewName := "Test Product"
	productNewDescription := "The description of test product"
	product.Name = productNewName
	product.Description = productNewDescription
	result, err := pr.Update(&product)
	if err != nil {
		t.Fatal(err)
	}
	if !result {
		t.Fatal("Product update failed")
	}

	// Read the product again from database
	updatedProduct, err := pr.GetByID(product.ProductID)
	if err != nil {
		t.Fatal(err)
	}
	if updatedProduct.Name != productNewName {
		t.Fatal("Inserted product name doesn't match")
	}
	if updatedProduct.Description != productNewDescription {
		t.Fatal("Inserted product description doesn't match")
	}

	// Error test for update
	pr.Session.DB.Close()
	result, err = pr.Update(&product)
	if err == nil || result {
		t.Fatal("Product update should be failed if database connection was closed")
	}
}

func TestDelete(t *testing.T) {
	// Init session
	sess, err := dbSessionSetup()
	if err != nil {
		t.Fatal(err)
	}

	// Insert product
	pr := ProductRepository{Session: sess}
	pID, err := pr.Add("p1", "d1", 1)
	if err != nil {
		t.Fatal(err)
	}

	// Get the product in database
	product, err := pr.GetByID(pID)
	if err != nil {
		t.Fatal(err)
	}

	// Delete product
	result, err := pr.Delete(product.ProductID)
	if err != nil {
		t.Fatal(err)
	}
	if !result {
		t.Fatal("Delete product failed")
	}

	// Check the product in database
	checkProduct, err := pr.GetByID(product.ProductID)
	if err != nil {
		t.Fatal(err)
	}
	if checkProduct.ProductID > 0 {
		t.Fatal("Product still exists in database")
	}

	// Error test for delete
	pr.Session.DB.Close()
	result, err = pr.Delete(product.ProductID)
	if err == nil || result {
		t.Fatal("Product delete should be failed if database connection was closed")
	}

}

func TestGetAll(t *testing.T) {
	// Init session
	sess, err := dbSessionSetup()
	if err != nil {
		t.Fatal(err)
	}

	// Insert product
	pr := ProductRepository{Session: sess}
	pr.Add("p1", "d1", 1)
	pr.Add("p2", "d2", 1)
	pr.Add("p3", "d3", 1)
	pr.Add("p4", "d4", 1)
	pr.Add("p5", "d5", 1)

	// Get All
	products, count, err := pr.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	if count != 5 {
		t.Fatal("Get all count should return 5")
	}
	if products[0].Name != "p1" {
		t.Fatal("Fist product name should be 'p1'")
	}
	if products[4].Description != "d5" {
		t.Fatal("Fifth product description should be 'd5'")
	}
}
