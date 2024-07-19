package usecase

import (
	"context"
	"kodinggo/internal/model"

	"github.com/sirupsen/logrus"
)

type StoryUsecase struct {
	storyRepo model.IStoryRepository
	log       *logrus.Logger
}

func NewStoryUsecase(
	storyRepo model.IStoryRepository,
	log *logrus.Logger,
) model.IStoryUsecase {
	return &StoryUsecase{
		storyRepo: storyRepo,
		log:       log,
	}
}

func (s *StoryUsecase) FindAll(ctx context.Context, filter model.StoryFilter) ([]*model.Story, error) {
	log := s.log.WithFields(logrus.Fields{
		"ctx":    ctx,
		"limit":  filter.Limit,
		"offset": filter.Offset,
	})

	stories, err := s.storyRepo.FindAll(ctx, filter)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return stories, nil
}

func (s *StoryUsecase) FindById(ctx context.Context, id int64) (*model.Story, error) {
	log := s.log.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	story, err := s.storyRepo.FindById(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return story, nil
}

func (s *StoryUsecase) Create(ctx context.Context, in model.CreateStoryInput) error {
	log := s.log.WithFields(logrus.Fields{
		"ctx":     ctx,
		"title":   in.Title,
		"content": in.Content,
	})

	story := model.Story{
		Title:   in.Title,
		Content: in.Content,
	}

	err := s.storyRepo.Create(ctx, story)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *StoryUsecase) Update(ctx context.Context, in model.UpdateStoryInput) error {
	log := s.log.WithFields(logrus.Fields{
		"ctx":     ctx,
		"id":      in.Id,
		"title":   in.Title,
		"content": in.Content,
	})

	story := model.Story{
		Id:      in.Id,
		Title:   in.Title,
		Content: in.Content,
	}

	err := s.storyRepo.Update(ctx, story)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *StoryUsecase) Delete(ctx context.Context, id int64) error {
	log := s.log.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := s.storyRepo.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
