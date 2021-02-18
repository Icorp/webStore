package webstore

// list of products ...
type ProductList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// List of Users ...
type UsersList struct {
	Id     int
	UserId int
	ListId int
}

// Product item ...
type ProductItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// List item ...
type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
