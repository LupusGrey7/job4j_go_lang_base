package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type FindByNameItemsResponse struct {
	Items []ItemRequest `json:"items"`
}

func (s *Server) FindByNameItem(c *fiber.Ctx) error {
	//парсим url запроса на предмет параметров ("id"/ "name") сразу в переменную
	requestParam := c.Query("name")
	if requestParam == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	items, err := s.Repository.FindByName(c.Context(), requestParam)
	if err != nil {
		log.Errorw("s.Repository.FindByName", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	res := make([]ItemRequest, 0, len(items))
	for _, item := range items {
		res = append(res, ItemRequest{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return c.Status(fiber.StatusOK).JSON(GetItemsResponse{Items: res})
}
