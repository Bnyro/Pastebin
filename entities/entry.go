package entities

type Entry struct {
	ID		uint32 `gorm:"primaryKey;autoIncrement:true"`
	UUID	string `json:"uuid"`
	Title 	string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Expires uint   `json:"expires" form:"expires"`
}