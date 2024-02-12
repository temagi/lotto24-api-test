package tests

type SearchResultObject struct {
	Id           int    `json:"id"`
	Key          string `json:"key"`
	Title        string `json:"title"`
	Excerpt      string `json:"excerpt"`
	MatchedTitle string `json:"matched_title"`
	Description  string `json:"description"`
}

type SearchContentResponse struct {
	Pages []SearchResultObject `json:"pages"`
}

type Latest struct {
	Id        uint64 `json:"id"`
	Timestamp string `json:"timestamp"`
}

type License struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

type Page struct {
	Id           uint    `json:"id"`
	Key          string  `json:"key"`
	Title        string  `json:"title"`
	Latest       Latest  `json:"latest"`
	ContentModel string  `json:"content_model"`
	License      License `json:"license"`
	HtmlUrl      string  `json:"html_url"`
}
