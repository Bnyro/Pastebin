package entities

type Entry struct {
	UUID	string `json:"uuid" form:"uuid"`
	Content string `json:"content" form:"content"`
}