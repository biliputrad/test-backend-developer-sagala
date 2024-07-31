package article_dto

type Create struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedBy int64
}
