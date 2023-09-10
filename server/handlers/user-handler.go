package handlers

import (
    "github.com/gofiber/fiber/v2"
)

type UserHandler struct {
    app fiber.App

    routeGroup fiber.Router

    di *DependencyInjector
}

func NewUserHandler(app fiber.App, di *DependencyInjector) {
    handler := &UserHandler{app, di}
    handler.RegisterRoutes()

    return handler
}

func (h *UserHandler) registerRoutes() {
    routeGroup := h.app.Group("/users")

    routeGroup.Get("/", h.getAll)

    h.routeGroup = routeGroup
}

func (h *UserHandler) getAll(c *fiber.Ctx) error {
    return c.JSON(h.di.UserService.FindAll())
}
