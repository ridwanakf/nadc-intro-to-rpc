package internal

// Book struct for database query
type Book struct {
	ID       int     `json:"book_id"`
	Title    string  `json:"book_title"`
	Author   string  `json:"book_author_name"`
	Category string  `json:"book_category"`
	Rate     float32 `json:"book_rate"`
}
