package dbs

import (
	"database/sql"
	"encoding/json"
	"food_market/pkg/models"
)

type CategoryModel struct {
	DB *sql.DB
}

func (m *CategoryModel) GetAllCategories() ([]byte, error) {
	stmt := `SELECT * FROM categories`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var categories []*models.Category

	for rows.Next() {
		category := &models.Category{}
		err = rows.Scan(&category.ID,
			&category.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	convertedCategories, err := json.Marshal(categories)
	return convertedCategories, nil
}

func (m *CategoryModel) CreateCategory(category *models.Category) ([]byte, error) {
	stmt := `INSERT INTO categories (category_name) VALUES ($1) RETURNING id`

	err := m.DB.QueryRow(stmt, category.CategoryName).Scan(&category.ID)
	if err != nil {
		return nil, err
	}

	convertedCategory, err := json.Marshal(category)
	if err != nil {
		return nil, err
	}

	return convertedCategory, nil
}

// update
func (m *CategoryModel) UpdateCategory(category *models.Category) error {
	stmt := `
		UPDATE categories
		SET
			category_name = $2	
		WHERE
			id = $1
			`

	_, err := m.DB.Exec(stmt, category.ID, category.CategoryName)
	return err
}

// delete
func (m *CategoryModel) DeleteCategory(categoryID int) error {
	stmt := `DELETE FROM categories WHERE id = $1`

	_, err := m.DB.Exec(stmt, categoryID)
	return err
}
