package store

import "time"

type User struct {
	ID        int       `db:"id"`
	UserType  string    `db:"type"`
	Name      string    `db:"name"`
	Password  []byte    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type userStore struct {
	db *dbContext
}

func newUserStore(db *dbContext) *userStore {
	return &userStore{db: db}
}
