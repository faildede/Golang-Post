package blog

import "time"

type Post struct {
	Title       string
	Author      string
	Time        time.Time
	Path        string
	Description string
	ImagePath   string
	Markdown    string
}

type RenderedPost struct { 
	Post post
	HTML string
}