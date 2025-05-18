package models

// Article -
type Article struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"` // Уникальный идентификатор заметки, автоматически инкрементируемый
	Title    string `gorm:"size:150;not null;unique"` // Заголовок заметки
	Content  string `gorm:"type:text;not null"`       // Содержание заметки
	AuthorID uint   `gorm:"not null"`                 //
	Author   User   `gorm:"not null"`                 // Автор заметки
}
