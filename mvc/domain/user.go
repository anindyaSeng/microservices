package domain

type User struct {
	UserId    uint64 `jason:"user_id"`
	FirstName string `jason:"first_name"`
	LastName  string `jason:"last_name"`
	Email     string `jason:"email"`
}
