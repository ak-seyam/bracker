package dto

type CreateUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserDto struct {
	Username string   `json:"username"`
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Groups   []string `json:"groups"`
}
