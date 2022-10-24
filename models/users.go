package models

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Id        string `db:"Id" json:"id" swaggerignore:"true"`
	FirstName string `db:"FirstName" json:"firstName"`
	LastName  string `db:"LastName" json:"lastName"`
	Bio       string `db:"Bio" json:"bio"`
	Username  string `db:"Username" json:"username"`
}

type GetUserResponse struct {
	Id        string `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Username  string `db:"username" json:"username"`
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	Bio       string `db:"bio" json:"bio"`
}

type SignUpRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
	HashedPassword       string `json:"hashedPassword" swaggerignore:"true"`
	Username             string `json:"username"`
}

type SignUpResponse struct {
	Id    string `db:"id" json:"id"`
	Email string `db:"email" json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserCoordinates struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
