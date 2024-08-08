package utils

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yuanzix/userAuth/internal/database"
	"github.com/yuanzix/userAuth/models"
)

type Storage interface {
	CreateUser(*models.User) (*database.User, error)
	DeleteUser(string) error
	UpdateUser(*models.User) (*database.User, error)
	GetUserByEmail(string) (*database.User, error)
	GetAllUsers() (*[]database.User, error)
}

type PostgresStore struct {
	queries *database.Queries
}

func NewPostgresStore() (*PostgresStore, error) {
	host, port, username, dbName, password, err := ReadPostgresDetails()

	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=require", host, port, username, dbName, password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	queries := database.New(db)

	return &PostgresStore{
		queries: queries,
	}, nil
}

func (s *PostgresStore) CreateUser(u *models.User) (*database.User, error) {

	user, err := s.queries.CreateUser(context.Background(), database.CreateUserParams{
		Email:          u.Email,
		Username:       u.Username,
		HashedPassword: u.HashedPassword,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		DateOfBirth:    u.DateOfBirth,
	})
	return &user, err
}

func (s *PostgresStore) DeleteUser(email string) error {
	err := s.queries.DeleteUser(context.Background(), email)
	return err
}

func (s *PostgresStore) UpdateUser(u *models.User) (*database.User, error) {
	return &database.User{}, nil
}

func (s *PostgresStore) GetUserByEmail(email string) (*database.User, error) {
	user, err := s.queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return &database.User{}, err
	}
	return &user, nil
}

func (s *PostgresStore) GetAllUsers() (*[]database.User, error) {
	users, err := s.queries.GetAllUsers(context.Background())
	if err != nil {
		return nil, err
	}
	return &users, nil
}
