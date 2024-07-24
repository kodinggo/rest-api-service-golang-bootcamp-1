package usecase

import (
	"context"
	"kodinggo/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// StoryUsecase :nodoc:
type StoryUsecase struct {
	storyRepo model.IStoryRepository
}

var v = validator.New()

// NewStoryUsecase :nodoc:
func NewStoryUsecase(
	storyRepo model.IStoryRepository,
) model.IStoryUsecase {
	return &StoryUsecase{
		storyRepo: storyRepo,
	}
}

// FindAll :nodoc:
func (s *StoryUsecase) FindAll(ctx context.Context, filter model.StoryFilter) ([]*model.Story, error) {
	log := logrus.WithFields(logrus.Fields{
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

// FindById :nodoc:
func (s *StoryUsecase) FindById(ctx context.Context, id int64) (*model.Story, error) {
	log := logrus.WithFields(logrus.Fields{
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
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"title":   in.Title,
		"content": in.Content,
	})

	err := s.validateCreateStoryInput(ctx, in)
	if err != nil {
		log.Error(err)
		return err
	}

	story := model.Story{
		Title:   in.Title,
		Content: in.Content,
	}

	err = s.storyRepo.Create(ctx, story)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *StoryUsecase) Update(ctx context.Context, in model.UpdateStoryInput) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"id":      in.Id,
		"title":   in.Title,
		"content": in.Content,
	})

	err := s.validateUpdateStoryInput(ctx, in)
	if err != nil {
		log.Error(err)
		return err
	}

	newStory := model.Story{
		Id:          in.Id,
		Title:       in.Title,
		Content:     in.Content,
		PublishedAt: in.PublishedAt,
	}

	err = s.storyRepo.Update(ctx, newStory)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *StoryUsecase) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
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

func (s *StoryUsecase) validateCreateStoryInput(ctx context.Context, in model.CreateStoryInput) error {
	err := v.StructCtx(ctx, in)
	if err != nil {
		log.Error(err)
		return model.ErrInvalidInput
	}

	return nil
}

func (s *StoryUsecase) validateUpdateStoryInput(ctx context.Context, in model.UpdateStoryInput) error {
	err := v.Struct(in)
	if err != nil {
		log.Error(err)
		return model.ErrInvalidInput
	}

	story, err := s.storyRepo.FindById(ctx, in.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	if in.PublishedAt.Before(story.CreatedAt) {
		return model.ErrPublishedAtLessThanCreatedAt
	}

	return nil
}
