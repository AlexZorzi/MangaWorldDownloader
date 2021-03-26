package main

type Chapter struct {
	ID    string   `json:"_id"`
	Pages []string `json:"pages"`
	Manga string   `json:"manga"`
	Name  string   `json:"name"`
	Slug  string   `json:"slugFolder"`
}

