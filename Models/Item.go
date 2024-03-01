package models

type Item struct {
	ID          [16]byte
	LocationID  [16]byte
	Name        string
	Description string

	Location Location
}
