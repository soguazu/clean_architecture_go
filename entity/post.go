package entity

type Post struct {
	ID    int64  `json:"id"`
	Text  string `json:"text"`
	Title string `json:"title"`
}
