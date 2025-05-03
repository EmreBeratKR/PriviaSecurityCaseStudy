package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(context *fiber.Ctx) error {
	log.Printf("Method: %s | Path: %s", context.Method(), context.Path())
	return context.Next()
}
