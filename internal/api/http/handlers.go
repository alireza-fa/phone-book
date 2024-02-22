package http

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
)

func (handler *Server) register(c *fiber.Ctx) error {
	request := struct{ Email, Password string }{}
	if err := c.BodyParser(&request); err != nil {
		errString := "Error parsing request body"
		handler.logger.Error(errString, zap.Any("request", request), zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	// TODO: if user with this email exists ot not

	// TODO: create user

	token, err := handler.token.CreateTokenString(1)
	if err != nil {
		errString := "Error creating JWT token for user"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	}

	response := map[string]string{"Token": token}
	return c.Status(http.StatusCreated).JSON(&response)
}

func (handler *Server) login(c *fiber.Ctx) error {
	request := struct{ Email, Password string }{}
	if err := c.BodyParser(&request); err != nil {
		errString := "Error parsing request body"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	// TODO: if user with combination of email and password exists or not

	token, err := handler.token.CreateTokenString(1)
	if err != nil {
		errString := "Error creating JWT token for user"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	}

	response := map[string]string{"Token": token}
	return c.Status(http.StatusOK).JSON(&response)
}

func (handler *Server) needsAuthentication(c *fiber.Ctx) error {
	userId, ok := c.Locals("user-id").(uint64)
	if !ok || userId == 0 {
		handler.logger.Error("Invalid user-id local")
		return c.SendStatus(http.StatusInternalServerError)
	}

	return nil
}
