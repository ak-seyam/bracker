package dto

type UserDto struct {
	Name   string   `json:"name"`
	Groups []string `json:"groups"`
	ID     string   `json:"id"`
}
