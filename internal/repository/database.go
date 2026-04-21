package internal

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

// Database initialization
func InitDB(dataSourceName string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }

    // Test the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}

// UserRepository implementation
type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

// Add user methods here
// For example: func (repo *UserRepository) CreateUser(user User) error {...}

// RoomRepository implementation
type RoomRepository struct {
    db *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
    return &RoomRepository{db: db}
}

// Add room methods here
// For example: func (repo *RoomRepository) CreateRoom(room Room) error {...}