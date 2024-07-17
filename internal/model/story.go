package model

import "context"

type IStoryRepository interface {
	FindAll(ctx context.Context, filter StoryFilter) ([]*Story, error)
	FindById(ctx context.Context, id int64) (*Story, error)
	Create(ctx context.Context, story Story) error
	Update(ctx context.Context, story Story) error
	Delete(ctx context.Context, id int64) error
}

type IStoryUsecase interface {
	FindAll(ctx context.Context, filter StoryFilter) ([]*Story, error)
	FindById(ctx context.Context, id int64) (*Story, error)
	Create(ctx context.Context, in CreateStoryInput) error
	Update(ctx context.Context, in UpdateStoryInput) error
	Delete(ctx context.Context, id int64) error
}

type Story struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
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
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
