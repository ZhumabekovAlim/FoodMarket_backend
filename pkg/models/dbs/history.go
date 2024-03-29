package dbs

//
//import (
//	"database/sql"
//	"encoding/json"
//	"food_market/pkg/models"
//)
//
//type OrderHistoryModel struct {
//	DB *sql.DB
//}
//
//// show
//func (m *OrderHistoryModel) GetHistoryByUserId(history *models.OrderHistory) ([]byte, error) {
//	stmt := `SELECT * FROM order_history WHERE user_id = $1`
//
//	rows, err := m.DB.Query(stmt, history.UserId)
//	if err != nil {
//		return nil, err
//	}
//
//	var histories []*models.OrderHistory
//
//	for rows.Next() {
//		history := &models.OrderHistory{}
//		err = rows.Scan(
//			&history.ID,
//			&history.ProductId,
//			&history.UserId,
//			&history.Quantity,
//			&history.DateTime,
//			//&history.OrderId,
//		)
//		if err != nil {
//			return nil, err
//		}
//		// Append it to the slice of snippets.
//		histories = append(histories, history)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	convertedhistory, err := json.Marshal(histories)
//	if err != nil {
//		return nil, err
//	}
//
//	return convertedhistory, nil
//}
//
//// show
//func (m *OrderHistoryModel) GetAllHistory() ([]byte, error) {
//	stmt := `SELECT * FROM order_history`
//
//	rows, err := m.DB.Query(stmt)
//	if err != nil {
//		return nil, err
//	}
//
//	var histories []*models.OrderHistory
//
//	for rows.Next() {
//		history := &models.OrderHistory{}
//		err = rows.Scan(
//			&history.ID,
//			&history.ProductId,
//			&history.UserId,
//			&history.Quantity,
//			&history.DateTime,
//			//&history.OrderId,
//		)
//		if err != nil {
//			return nil, err
//		}
//		// Append it to the slice of snippets.
//		histories = append(histories, history)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	convertedhistory, err := json.Marshal(histories)
//	if err != nil {
//		return nil, err
//	}
//
//	return convertedhistory, nil
//}
//
//// create
//func (m *OrderHistoryModel) CreateHistory(history *models.OrderHistory) ([]byte, error) {
//	stmt := `INSERT INTO order_history ( product_id, user_id, quantity) VALUES ($1,$2,$3) RETURNING id`
//
//	err := m.DB.QueryRow(stmt, history.ProductId, history.UserId, history.Quantity).Scan(&history.ID)
//	if err != nil {
//		return nil, err
//	}
//
//	convertedhistory, err := json.Marshal(history)
//	if err != nil {
//		return nil, err
//	}
//
//	return convertedhistory, nil
//}
//
//func (m *ProductModel) GetProductById(productId string) ([]byte, error) {
//	stmt := `SELECT * FROM product WHERE id = $1`
//
//	rows, err := m.DB.Query(stmt, productId)
//	if err != nil {
//		return nil, err
//	}
//
//	p := &models.Product{}
//	err = rows.Scan(
//		&p.ID,
//		&p.ProductName,
//		&p.CategoryId,
//		&p.Price,
//		&p.Quantity,
//		&p.Type,
//		&p.PhotoUrl,
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	convertedProduct, err := json.Marshal(p)
//	if err != nil {
//		return nil, err
//	}
//
//	return convertedProduct, nil
//}
//
//func (m *ProductModel) GetHistoryById(productId string) ([]byte, error) {
//	stmt := `SELECT * FROM product WHERE id = $1`
//
//	rows, err := m.DB.Query(stmt, productId)
//	if err != nil {
//		return nil, err
//	}
//
//	p := &models.Product{}
//	err = rows.Scan(
//		&p.ID,
//		&p.ProductName,
//		&p.CategoryId,
//		&p.Price,
//		&p.Quantity,
//		&p.Type,
//		&p.PhotoUrl,
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	convertedProduct, err := json.Marshal(p)
//	if err != nil {
//		return nil, err
//	}
//
//	return convertedProduct, nil
//}
//
//// update
//func (m *OrderHistoryModel) UpdateHistory(history *models.OrderHistory) error {
//	stmt := `
//		UPDATE order_history
//		SET
//			product_id = $2,
//			user_id = $3,
//			quantity = $4
//		WHERE
//			id = $1
//	`
//
//	_, err := m.DB.Exec(stmt, history.ID, history.ProductId, history.UserId, history.Quantity, history.DateTime)
//	return err
//}
//
//// delete
//func (m *OrderHistoryModel) DeleteHistory(historyID int) error {
//	stmt := `DELETE FROM order_history WHERE id = $1`
//
//	_, err := m.DB.Exec(stmt, historyID)
//	return err
//}
