package models

type TodoList struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"description"`
}
