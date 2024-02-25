package http

import (
	"github.com/alireza-fa/phone-book/internal/models"
	"github.com/alireza-fa/phone-book/pkg/rdbms"
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

	user, err := handler.repository.GetUserByEmail(request.Email)
	if err != nil && err.Error() != rdbms.ErrNotFound {
		errString := "Error while retrieving data from database"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	} else if err == nil || user != nil && user.Id != 0 {
		errString := "User with given email already exists"
		handler.logger.Error(errString, zap.String("email", request.Email))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	user = &models.User{Email: request.Email, Password: request.Password}
	if err := handler.repository.CreateUser(user); err != nil {
		errString := "Error happened while creating the user"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	} else if user.Id == 0 {
		errString := "Error invalid user id created"
		handler.logger.Error(errString, zap.Any("user", user))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	}

	token, err := handler.token.CreateTokenString(user.Id)
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

	user, err := handler.repository.GetUserByEmailAndPassword(request.Email, request.Password)
	if err != nil {
		errString := "wrong email or password has been given"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	} else if user == nil {
		errString := "Error invalid user retured"
		handler.logger.Error(errString, zap.Any("request", request))
		return c.Status(http.StatusInternalServerError).SendString(errString)
	}

	token, err := handler.token.CreateTokenString(user.Id)
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
