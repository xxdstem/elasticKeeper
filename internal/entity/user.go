package entity

type User struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"username" db:"username"`
	NameAKA    string `json:"username_aka" db:"username_aka"`
	Privileges int    `json:"privileges" db:"privileges"`
	BetaKey    string `json:"beta_key" db:"beta_key"`
	Email      string `json:"email" db:"email"`
	Country    string `json:"country" db:"country"`
}
