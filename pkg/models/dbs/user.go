package dbs

import (
	"database/sql"
	"errors"
	"food_market/pkg/models"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (id, name, email, phone, password) VALUES ($1, $2, $3, $4, $5);`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword), "user")
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && strings.Contains(pgErr.Message, "users_uc_email") {
			return models.ErrDuplicateEmail
		}
		return err
	}

	return nil
}

func (m *UserModel) GetAll() ([]*models.User, error) {
	stmt := `SELECT * FROM food_market.public.users`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	users := []*models.User{}

	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippets.
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UserModel) GetRole(id int) string {
	stmt := `SELECT role FROM food_market.public.users WHERE id = $1`
	var role string
	err := m.DB.QueryRow(stmt, id).Scan(&role)
	if err != nil {
		return ""
	}

	return role
}

func (m *UserModel) Get(id int) (*models.User, error) {
	stmt := `SELECT * FROM food_market.public.users WHERE id = $1`

	userRow := m.DB.QueryRow(stmt, id)

	u := &models.User{}

	err := userRow.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Password, &u.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return u, nil
}
