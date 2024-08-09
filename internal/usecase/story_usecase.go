package usecase

import (
	"context"
	"kodinggo/internal/model"
	"kodinggo/internal/worker"
	"strconv"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"

	pbCategory "github.com/kodinggo/category-service-gb1/pb/category"
	pbComment "github.com/kodinggo/comment-service-gb1/pb/comment"
)

// StoryUsecase :nodoc:
type StoryUsecase struct {
	storyRepo       model.IStoryRepository
	commentService  pbComment.CommentServiceClient
	categoryService pbCategory.CategoryServiceClient
	workerClient    *worker.AsynqClient
}

var v = validator.New()

// NewStoryUsecase :nodoc:
func NewStoryUsecase(
	storyRepo model.IStoryRepository,
	commentService pbComment.CommentServiceClient,
	categoryService pbCategory.CategoryServiceClient,
	workerClient *worker.AsynqClient,
) model.IStoryUsecase {
	return &StoryUsecase{
		storyRepo:       storyRepo,
		commentService:  commentService,
		categoryService: categoryService,
		workerClient:    workerClient,
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

	var wg sync.WaitGroup
	// getting comments from comment service
	for idx, story := range stories {
		wg.Add(1)
		go func(idx int, story *model.Story) {
			defer wg.Done()

			idReq := strconv.Itoa(int(story.Id))

			commentFilter := pbComment.CommentRequest{
				StoryId: idReq,
			}

			comments, err := s.commentService.FindComments(ctx, &commentFilter)
			if err != nil {
				log.Error(err)
				return
			}

			for _, comment := range comments.Comments {
				story.Comments = append(story.Comments, model.Comment{
					Id:        comment.Id,
					StoryId:   comment.StoryId,
					Content:   comment.Content,
					CreatedAt: comment.CreatedAt,
					UpdatedAt: comment.UpdatedAt,
				})
			}
		}(idx, story)

		wg.Add(1)
		go func(idx int, story *model.Story) {
			defer wg.Done()
			categoryId := story.CategoryId

			category, err := s.categoryService.FindCategoryById(ctx, &pbCategory.CategoryRequest{Id: categoryId})
			if err != nil {
				log.Error(err)
				return
			}

			catId, err := strconv.Atoi(category.Id)
			if err != nil {
				log.Error(err)
			}

			stories[idx].Category = &model.Category{
				Id:   int64(catId),
				Name: category.Name,
			}
		}(idx, story)

		wg.Wait()
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

	// TODO: Run this!
	// Send task to queue
	_, err = s.workerClient.SendEmail(worker.SendEmailPayload{
		From:    "john@gmail.com",
		To:      "mark@yahoo.com",
		Subject: "Test Queue",
		Message: "Mencoba queue messaging",
	})
	if err != nil {
		log.Errorf("failed enqueue send send email task")
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
