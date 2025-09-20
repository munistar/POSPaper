package domain

type Agent struct {
	ID        int64  `json:"id" db:"id"`
	UserID    int64  `json:"user_id" db:"user_id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email" db:"email"`
	OfficeID  int64  `json:"office_id" db:"office_id"`
}

type Customer struct {
	ID        int64  `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Phone     string `json:"phone" db:"phone"`
	// optionally link to user account if customers sign in
	UserID *int64 `json:"user_id,omitempty" db:"user_id"`
}
