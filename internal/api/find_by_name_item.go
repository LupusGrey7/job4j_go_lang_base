package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type FindByItemResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *Server) FindByNameItem(c *fiber.Ctx) error {
	var resp FindByItemResponse

	//парсим url запроса на предмет параметров ("id"/ "name") сразу в переменную
	requestParam := c.Query("name", "")
	if requestParam == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	trackerItem, err := s.Repository.FindByName(c.Context(), requestParam)
	if err != nil {
		log.Errorw("s.Repository.FindByName", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	resp.ID = trackerItem.ID
	resp.Name = trackerItem.Name

	return OkResponse(c, resp)
}
