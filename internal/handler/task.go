package handler

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *Handler) GetTaskById(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	task, err := h.service.Task.GetOneById(idInt)
	if err != nil {
		return err
	}
	return c.JSON(task)
}

func (h *Handler) ApiInfo(c *fiber.Ctx) error {
	return c.JSON("hello to fiber api")
}
