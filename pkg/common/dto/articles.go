package dto

// Create article schema
type AddArticleDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Article response schema
type ArticleDTO struct {
	ID       uint      `json:"id"`                 // Уникальный идентификатор заметки, автоматически инкрементируемый
	Title    string    `json:"title"`              // Заголовок заметки
	Content  string    `json:"content,omitempty"`  // Содержание заметки
	AuthorID uint      `json:"authorId,omitempty"` //
	Author   AuthorDTO `json:"author,omitzero"`    // Автор заметки
}

// Author response schema
type AuthorDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
