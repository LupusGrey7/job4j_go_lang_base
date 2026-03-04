package api

import "github.com/gofiber/fiber/v2"

const (
	ItemPath  = "/item"
	ItemsPath = "/items"
)

func (s *Server) Route(route fiber.Router) {
	route.Post(ItemPath, s.CreateItem)
	route.Put(ItemPath, s.UpdateItem)
	route.Delete(ItemPath, s.DeleteItem)
	route.Get(ItemsPath, s.GetItems)
	route.Get(ItemPath, s.FindByNameItem)
}
