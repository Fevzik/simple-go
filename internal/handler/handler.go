package handler

import (
	"github.com/Fevzik/simple-go/internal/constants"
	"github.com/Fevzik/simple-go/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRoutes(app *fiber.App, mode string) {
	apiGroup := app.Group("/api")
	moduleGroup := apiGroup.Group("/" + constants.ModuleCode)
	moduleGroup.Get("/", h.ApiInfo)
	moduleGroup.Get("/:id", h.GetTaskById)
}
