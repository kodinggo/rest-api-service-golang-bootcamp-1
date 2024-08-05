package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"kodinggo/internal/model"

	"github.com/redis/go-redis/v9"
)

type StoryRepository struct {
	db    *sql.DB
	redis *redis.Client
}

func NewStoryRepository(db *sql.DB, redis *redis.Client) model.IStoryRepository {
	return &StoryRepository{
		db:    db,
		redis: redis,
	}
}

func (s *StoryRepository) FindAll(ctx context.Context, filter model.StoryFilter) ([]*model.Story, error) {
	res, err := s.db.QueryContext(ctx, "SELECT id, title, content, category_id, published_at, created_at FROM stories LIMIT ? OFFSET ?", filter.Limit, filter.Offset)
	if err != nil {
		return nil, err
	}

	var stories []*model.Story
	for res.Next() {
		var story model.Story
		if err := res.Scan(&story.Id, &story.Title, &story.Content, &story.CategoryId, &story.PublishedAt, &story.CreatedAt); err != nil {
			return nil, err
		}
		stories = append(stories, &story)
	}

	return stories, nil
}

func (s *StoryRepository) FindById(ctx context.Context, id int64) (*model.Story, error) {
	storyKey := getStoryKey(id)

	var story model.Story
	st, err := s.redis.Get(ctx, storyKey).Result()
	if err == nil {
		err := json.Unmarshal([]byte(st), &story)
		if err != nil {
			return nil, err
		}

		return &story, nil
	}

	res, err := s.db.QueryContext(ctx, "SELECT id, title, content, category_id, published_at, created_at FROM stories WHERE id=?", id)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		if err := res.Scan(&story.Id, &story.Title, &story.Content, &story.CategoryId, &story.PublishedAt, &story.CreatedAt); err != nil {
			return nil, err
		}
	}

	storyJson, err := json.Marshal(story)
	if err != nil {
		return nil, err
	}

	err = s.redis.Set(ctx, storyKey, string(storyJson), 0).Err()
	if err != nil {
		return nil, err
	}

	return &story, nil
}

func (s *StoryRepository) Create(ctx context.Context, story model.Story) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO stories (title, content) VALUES (?, ?)", story.Title, story.Content)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoryRepository) Update(ctx context.Context, story model.Story) error {
	_, err := s.db.ExecContext(ctx, "UPDATE stories SET title=?, content=? WHERE id=?", story.Title, story.Content, story.Id)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoryRepository) Delete(ctx context.Context, id int64) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM stories WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func getStoryKey(id int64) string {
	return fmt.Sprintf("%s:%d", model.StoryKey, id)
}
