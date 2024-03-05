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
	stmt := `SELECT * FROM category`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var categories []*models.Category

	for rows.Next() {
		category := &models.Category{}
		err = rows.Scan(
			&category.ID,
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

func (m *CategoryModel) GetCategoryById(categoryId string) ([]byte, error) {
	stmt := `SELECT * FROM category WHERE id = $1`

	row := m.DB.QueryRow(stmt, categoryId)

	c := &models.Category{}
	err := row.Scan(
		&c.ID,
		&c.CategoryName,
	)
	if err != nil {
		return nil, err
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	categoryResponse, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return categoryResponse, nil
}

func (m *CategoryModel) CreateCategory(category *models.Category) ([]byte, error) {
	stmt := `INSERT INTO category (category_name) VALUES ($1) RETURNING id`

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
		UPDATE category
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
	stmt := `DELETE FROM category WHERE id = $1`

	_, err := m.DB.Exec(stmt, categoryID)
	return err
}
