package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Deleted  bool   `json:"deleted"`
}

type UserDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type UserDetailedDTO struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Password string     `json:"password"`
	Deleted  bool       `json:"deleted"`
	Groups   []GroupDTO `json:"groups"`
}
