package webstore

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"name"`
	Password string `json:"name"`
}
