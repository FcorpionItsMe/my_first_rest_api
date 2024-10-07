package models

type TodoItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"description"`
	Done  bool   `json:"done"`
}
