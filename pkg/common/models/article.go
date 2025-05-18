package models

// Article -
type Article struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`    // Уникальный идентификатор заметки, автоматически инкрементируемый
	Title    string `json:"title" gorm:"size:150;not null;unique"` // Заголовок заметки
	Content  string `json:"content" gorm:"type:text;not null"`     // Содержание заметки
	AuthorId uint   `json:"authorId" gorm:"not null"`              // Автор заметки
}
