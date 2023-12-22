package models

type Post struct {
	Id      int
	Content string
}

type PostRequest struct {
	Content string `json:"content"`
}

type PostResponse struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}
