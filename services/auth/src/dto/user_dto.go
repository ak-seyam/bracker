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

type UserLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponseDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
