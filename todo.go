package todo

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"id" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Descripition string `json:"description"`
	Done         bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
