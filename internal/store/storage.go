package store

import ("context"
	"database/sql"
	"errors"
	"time"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}

	Users interface {
		Create(context.Context *User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage {
		Posts: &PostsStore{db},
		Users: &UsersStore{db}
	}
}