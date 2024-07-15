package usecase

import (
	"context"
	"kodinggo/internal/model"
)

type IStoryUsecase interface {
	FindAll(ctx context.Context, filter model.StoryFilter) ([]*model.Story, error)
	FindById(ctx context.Context, id int64) (*model.Story, error)
	Create(ctx context.Context, in model.CreateStoryInput) error
	Update(ctx context.Context, in model.UpdateStoryInput) error
	Delete(ctx context.Context, id int64) error
}
