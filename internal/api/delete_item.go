package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type DeleteItemRequest struct {
	ID string `json:"id"`
}

func (s *Server) DeleteItem(c *fiber.Ctx) error {

	//парсим url запроса на предмет параметров ("id"/ "name") сразу в переменную
	requestParam := c.Query("id", "")
	if requestParam == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}

	err := s.Repository.Delete(c.Context(), requestParam)
	if err != nil {
		log.Errorw("s.Repository.DeleteByID: ", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
