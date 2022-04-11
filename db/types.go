package db

type AuthEntry struct {
	ID        *int64 `db:"id"`
	Name      string `db:"name"`
	Password  string `db:"password"`
	LastLogin int    `db:"last_login"`
}

type AuthRepository interface {
	Migrate() error
	GetByUsername(username string) (*AuthEntry, error)
	Create(entry *AuthEntry) error
	Update(entry *AuthEntry) error
	Delete(id int64) error
	Close() error
}
