package models

type Location struct {
	ID          [16]byte
	UserID      [16]byte
	Name        string
	Description string

	User  User
	Items []Item
}
