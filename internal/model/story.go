package model

type Story struct {
	Id      int64
	Title   string
	Content string
}

type StoryFilter struct {
	Offset int32
	Limit  int32
}

type CreateStoryInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateStoryInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
