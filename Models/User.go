package models

type User struct {
	ID       [16]byte
	UserName string
	Password string

	Locations []Location
}
