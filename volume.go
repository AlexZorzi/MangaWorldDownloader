package main

type Volume struct {
	ID    string `json:"_id"`
	Manga string `json:"manga"`
	Name  string `json:"name"`
	Slug  string `json:"slugFolder"`
	Image string `json:"image"`
}
