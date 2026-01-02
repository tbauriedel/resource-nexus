package database

type User struct {
	ID           int
	Name         string
	PasswordHash string
	IsAdmin      bool
}

type Permission struct {
	ID       int
	Category string
	Resource string
	Action   string
}
