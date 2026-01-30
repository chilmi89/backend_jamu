package models

type User struct {
	ID       int    `bun:"id,pk,autoincrement" json:"id"`
	Name     string `json:"name"`
	Email    string `bun:",unique" json:"email"`
	Password string `json:"-"` // Hidden in JSON
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
