package db

type AuthEntry struct {
	ID        *int64 `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	LastLogin int    `json:"last_login"`
}

type PrivilegeEntry struct {
	ID        int64  `json:"id"`
	Privilege string `json:"privilege"`
}

type AuthRepository interface {
	GetByUsername(username string) (*AuthEntry, error)
	Create(entry *AuthEntry) error
	Update(entry *AuthEntry) error
	Delete(id int64) error
}

type PrivilegeRepository interface {
	GetByID(id int64) ([]*PrivilegeEntry, error)
	Create(entry *PrivilegeEntry) error
	Delete(id int64, privilege string) error
}
