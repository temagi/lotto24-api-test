package tests

type Page struct {
	Id int `json:"id"`
	Key string `json:"key"`
	Title string `json:"title"`
	Excerpt string `json:"excerpt"`
	MatchedTitle string `json:"matched_title"`
	Description string `json:"description"`
}

type PagesResponse struct {
	Pages []Page `json:"pages"`
}