package article_dto

type Update struct {
	ID        int64
	Author    string `json:"author"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UpdatedBy int64
}
