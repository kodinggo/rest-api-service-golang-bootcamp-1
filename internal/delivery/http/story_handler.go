package handler

import (
	"kodinggo/internal/model"
	"kodinggo/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StoryHandler struct {
	storyUsecase usecase.IStoryUsecase
}

func NewStoryHandler(e *echo.Group, us usecase.IStoryUsecase) {
	handlers := &StoryHandler{
		storyUsecase: us,
	}

	e.GET("/stories", handlers.GetStories)
	e.GET("/stories/:id", handlers.GetStory)
	e.POST("/stories", handlers.CreateStory)
	e.PUT("/stories/:id", handlers.UpdateStory)
	e.DELETE("/stories/:id", handlers.DeleteStory)
}

func (s *StoryHandler) GetStories(c echo.Context) error {
	// Get query string limit and offset
	reqLimit := c.QueryParam("limit")
	reqOffset := c.QueryParam("offset")

	var limit, offset int32
	if reqLimit == "" {
		limit = 10 // default limit
	}
	if reqOffset == "" {
		offset = 0 // default offset
	}

	stories, err := s.storyUsecase.FindAll(c.Request().Context(), model.StoryFilter{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(200, response{
		Message: "success",
		Data:    stories,
	})
}

func (s *StoryHandler) GetStory(c echo.Context) error {
	id := c.Param("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	story, err := s.storyUsecase.FindById(c.Request().Context(), int64(parsedId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Message: "success",
		Data:    story,
	})
}

func (s *StoryHandler) CreateStory(c echo.Context) error {
	var in model.CreateStoryInput
	if err := c.Bind(&in); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	if err := s.storyUsecase.Create(c.Request().Context(), in); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	return c.JSON(http.StatusCreated, response{
		Message: "success",
	})
}

func (s *StoryHandler) UpdateStory(c echo.Context) error {
	storyId := c.Param("id")
	parsedId, err := strconv.Atoi(storyId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var in model.UpdateStoryInput
	if err := c.Bind(&in); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	if parsedId > 0 {
		in.Id = int64(parsedId)
	}

	if err := s.storyUsecase.Update(c.Request().Context(), in); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	return c.JSON(200, response{
		Message: "success",
	})
}

func (s *StoryHandler) DeleteStory(c echo.Context) error {
	id := c.Param("id")
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	if err := s.storyUsecase.Delete(c.Request().Context(), int64(parsedId)); err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	return c.JSON(200, response{
		Message: "success",
	})
}
