package repository

import (
	"context"
	"kodinggo/internal/model"
)

type IStoryRepository interface {
	FindAll(ctx context.Context, filter model.StoryFilter) ([]*model.Story, error)
	FindById(ctx context.Context, id int64) (*model.Story, error)
	Create(ctx context.Context, story model.Story) error
	Update(ctx context.Context, story model.Story) error
	Delete(ctx context.Context, id int64) error
}
