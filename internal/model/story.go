package model

import (
	"context"
	"errors"
	"time"
)

var (
	ErrInvalidInput                 = errors.New("invalid input")
	ErrPublishedAtLessThanCreatedAt = errors.New("published_at must be greater than created_at")
)

// IStoryRepository :nodoc:
type IStoryRepository interface {
	FindAll(ctx context.Context, filter StoryFilter) ([]*Story, error)
	FindById(ctx context.Context, id int64) (*Story, error)
	Create(ctx context.Context, story Story) error
	Update(ctx context.Context, story Story) error
	Delete(ctx context.Context, id int64) error
}

// IStoryUsecase :nodoc:
type IStoryUsecase interface {
	FindAll(ctx context.Context, filter StoryFilter) ([]*Story, error)
	FindById(ctx context.Context, id int64) (*Story, error)
	Create(ctx context.Context, in CreateStoryInput) error
	Update(ctx context.Context, in UpdateStoryInput) error
	Delete(ctx context.Context, id int64) error
}

// Story represents a story model
type Story struct {
	Id          int64      `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
}

// StoryFilter represent struct for story filter
type StoryFilter struct {
	Offset int32
	Limit  int32
}

type CreateStoryInput struct {
	Title   string `json:"title" validate:"required,min=3,max=50"`
	Content string `json:"content" validate:"required"`
}

type UpdateStoryInput struct {
	Id              int64      `json:"id"`
	Title           string     `json:"title" validate:"required"`
	Content         string     `json:"content"`
	PublishedAt     *time.Time `json:"published_at"`
	MediaExternalId string
}
