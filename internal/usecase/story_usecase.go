package usecase

import (
	"context"
	"kodinggo/internal/model"
	"kodinggo/internal/repository"
)

type StoryUsecase struct {
	storyRepo repository.IStoryRepository
}

func NewStoryUsecase(storyRepo repository.IStoryRepository) IStoryUsecase {
	return &StoryUsecase{
		storyRepo: storyRepo,
	}
}

func (s *StoryUsecase) FindAll(ctx context.Context, filter model.StoryFilter) ([]*model.Story, error) {
	stories, err := s.storyRepo.FindAll(ctx, filter)
	if err != nil {
		return nil, err
	}
	return stories, nil
}

func (s *StoryUsecase) FindById(ctx context.Context, id int64) (*model.Story, error) {
	story, err := s.storyRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return story, nil
}

func (s *StoryUsecase) Create(ctx context.Context, in model.CreateStoryInput) error {
	story := model.Story{
		Title:   in.Title,
		Content: in.Content,
	}

	return s.storyRepo.Create(ctx, story)
}

func (s *StoryUsecase) Update(ctx context.Context, in model.UpdateStoryInput) error {
	story := model.Story{
		Title:   in.Title,
		Content: in.Content,
	}

	return s.storyRepo.Update(ctx, story)
}

func (s *StoryUsecase) Delete(ctx context.Context, id int64) error {
	return s.storyRepo.Delete(ctx, id)
}
