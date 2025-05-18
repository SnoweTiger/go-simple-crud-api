package models

// User object
type User struct {
	ID       uint      `json:"id" gorm:"primaryKey;autoIncrement"` // Уникальный идентификатор заметки, автоматически инкрементируемый
	Name     string    `json:"name" gorm:"size:255;not null"`      // Заголовок заметки
	Login    string    `json:"login" gorm:"size:255;not null;unique"`
	Password string    `json:"password" gorm:"size:255;not null"` // Содержание заметки
	Articles []Article `json:"articles" gorm:"foreignKey:AuthorId"`
}
