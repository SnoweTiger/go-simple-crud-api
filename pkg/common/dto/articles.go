package dto

// Create article schema
type AddArticleDTO struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// Article response schema
type ArticleDTO struct {
	ID       uint      `json:"id" binding:"required"`                 // Уникальный идентификатор заметки, автоматически инкрементируемый
	Title    string    `json:"title" binding:"required"`              // Заголовок заметки
	Content  string    `json:"content,omitempty" binding:"required"`  // Содержание заметки
	AuthorID uint      `json:"authorId,omitempty" binding:"required"` //
	Author   AuthorDTO `json:"author,omitzero" binding:"required"`    // Автор заметки
}

// Author response schema
type AuthorDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
