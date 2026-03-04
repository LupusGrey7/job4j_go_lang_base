package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/tracker"
)

type UpdateItemRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type UpdateItemResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Response struct {
	Success bool        `json:"success" example:"false"`
	Message string      `json:"error" example:"Invalid request format"`
	Data    interface{} `json:"data"`
}

func OkResponse(
	c *fiber.Ctx,
	data interface{},
) error {
	return c.Status(200).JSON(&Response{
		Success: true,
		Data:    data,
	})
}

func (s *Server) UpdateItem(c *fiber.Ctx) error {
	var req UpdateItemRequest
	var resp UpdateItemResponse

	//парсим запрос на предмет параметров ("id", "name") сразу в "тело" UpdateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.ID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	item := tracker.Item{ID: req.ID, Name: req.Name}
	trackerItem, err := s.Repository.Update(c.Context(), item)
	if err != nil {
		log.Errorw("s.Repository.Delete", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	resp.ID = trackerItem.ID
	resp.Name = trackerItem.Name

	return OkResponse(c, resp)
}
