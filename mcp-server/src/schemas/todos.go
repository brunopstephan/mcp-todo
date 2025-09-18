package schemas

type Todo struct {
	ID          string `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Date        string `json:"date" gorm:"type:date;not null"`
	Done        bool   `json:"done" gorm:"not null;default:false"`
}
