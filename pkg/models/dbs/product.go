package dbs

import (
	"database/sql"
	"encoding/json"
	"food_market/pkg/models"
)

type ProductModel struct {
	DB *sql.DB
}

// create
func (m *ProductModel) CreateProduct(Product *models.Product) ([]byte, error) {
	stmt := `INSERT INTO product (product_name, category_id, price, quantity, type, photo_url) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	err := m.DB.QueryRow(stmt, Product.ProductName, Product.CategoryId, Product.Price, Product.Quantity, Product.Type, Product.PhotoUrl).Scan(&Product.ID)
	if err != nil {
		return nil, err
	}

	convertedProduct, err := json.Marshal(Product)
	if err != nil {
		return nil, err
	}

	return convertedProduct, nil
}

// show
func (m *ProductModel) GetAllProducts() ([]byte, error) {
	stmt := `SELECT * FROM product`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var products []*models.Product

	for rows.Next() {
		product := &models.Product{}
		err = rows.Scan(&product.ID,
			&product.ProductName,
			&product.CategoryId,
			&product.Price,
			&product.Quantity,
			&product.Type,
			&product.PhotoUrl)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippets.
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	convertedProduct, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	return convertedProduct, nil
}

func (m *ProductModel) GetProductById(productId string) ([]byte, error) {
	stmt := `SELECT * FROM product WHERE id = $1`

	row := m.DB.QueryRow(stmt, productId)

	p := &models.Product{}
	err := row.Scan(
		&p.ID,
		&p.ProductName,
		&p.CategoryId,
		&p.Price,
		&p.Quantity,
		&p.Type,
		&p.PhotoUrl,
	)
	if err != nil {
		return nil, err
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	convertedProduct, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return convertedProduct, nil
}

// update
func (m *ProductModel) UpdateProduct(product *models.Product) error {
	stmt := `
		UPDATE product
		SET
			product_name = $2,
			category_id = $3,
			price = $4,
			quantity = $5,
			type = $6,
			photo_url = $7
		WHERE
			id = $1`

	_, err := m.DB.Exec(stmt, product.ID, product.ProductName, product.CategoryId, product.Price, product.Quantity, product.Type, product.PhotoUrl)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

// delete
func (m *ProductModel) DeleteProduct(productID int) error {
	stmt := `DELETE FROM product WHERE id = $1`

	_, err := m.DB.Exec(stmt, productID)
	return err
}
