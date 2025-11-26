package v1

import (
	"github.com/fakelozic/tasker/internal/handler"
	"github.com/fakelozic/tasker/internal/middleware"
	"github.com/labstack/echo/v4"
)

func registerCommentRoutes(r *echo.Group, h *handler.CommentHandler, auth *middleware.AuthMiddleware) {
	// Comment operations
	comments := r.Group("/comments")
	comments.Use(auth.RequireAuth)

	// Individual comment operations
	dynamicComment := comments.Group("/:id")
	dynamicComment.PATCH("", h.UpdatedComment)
	dynamicComment.DELETE("", h.DeleteComment)
}
