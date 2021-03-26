package main

type Pages struct {
	Volumes []struct {
		Volume   Volume    `json:"volume"`
		Chapters []Chapter `json:"chapters"`
	} `json:"volumes"`
	SingleChapters []Chapter `json:"singleChapters"`
}
