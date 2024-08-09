package types

import "database/sql"

// I do not feel like rolling out my own auth.
// Not enough time
type User struct {
	Username string
	Email    string
	Public   bool
}

func (u User) Store(db *sql.DB) error {

	return nil
}

func (u User) Update(db *sql.DB) error {

	return nil
}

func (u User) Delete(db *sql.DB) error {

	return nil
}
