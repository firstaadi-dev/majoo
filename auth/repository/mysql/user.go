package mysql

import (
	"database/sql"

	"github.com/firstaadi-dev/majoo-backend-test/auth"
	"github.com/firstaadi-dev/majoo-backend-test/domain"
)

type mysqlUserRepo struct {
	DB *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) auth.UserRepository {
	return &mysqlUserRepo{
		DB: db,
	}
}

// GetUser implements auth.UserRepository
func (u *mysqlUserRepo) GetUser(username string, password string) (*domain.User, error) {
	var user domain.User
	err := u.DB.
		QueryRow("SELECT id, name, user_name, created_at, created_by, updated_at, updated_by FROM `Users` WHERE `user_name` = ? AND `password` = MD5(?)", username, password).
		Scan(
			&user.ID,
			&user.Name,
			&user.UserName,
			&user.CreatedAt,
			&user.CreatedBy,
			&user.UpdatedAt,
			&user.UpdatedBy,
		)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
