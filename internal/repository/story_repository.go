package repository

import (
	"context"
	"database/sql"
	"kodinggo/internal/model"
)

type StoryRepository struct {
	db *sql.DB
}

func NewStoryRepository(db *sql.DB) IStoryRepository {
	return &StoryRepository{db: db}
}

func (s *StoryRepository) FindAll(ctx context.Context, filter model.StoryFilter) ([]*model.Story, error) {
	res, err := s.db.QueryContext(ctx, "SELECT id, title, content FROM stories LIMIT $1 OFFSET $2", filter.Limit, filter.Offset)
	if err != nil {
		return nil, err
	}

	var stories []*model.Story
	for res.Next() {
		var story model.Story
		if err := res.Scan(&story.Id, &story.Title, &story.Content); err != nil {
			return nil, err
		}
		stories = append(stories, &story)
	}

	return stories, nil
}

func (s *StoryRepository) FindById(ctx context.Context, id int64) (*model.Story, error) {
	res, err := s.db.QueryContext(ctx, "SELECT id, title, content FROM stories WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	var story model.Story
	for res.Next() {
		if err := res.Scan(&story.Id, &story.Title, &story.Content); err != nil {
			return nil, err
		}
	}
	return &story, nil
}

func (s *StoryRepository) Create(ctx context.Context, story model.Story) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO stories (title, content) VALUES ($1, $2)", story.Title, story.Content)
	if err != nil {
		return err
	}
	return nil
}

func (s *StoryRepository) Update(ctx context.Context, story model.Story) error {
	_, err := s.db.ExecContext(ctx, "UPDATE stories SET title=$1, content=$2 WHERE id=$3", story.Title, story.Content, story.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *StoryRepository) Delete(ctx context.Context, id int64) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM stories WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
