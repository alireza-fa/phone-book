package http

import (
	"encoding/json"
	"fmt"
	"github.com/alireza-fa/phone-book/internal/repository"
	"github.com/alireza-fa/phone-book/pkg/token"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	logger     *zap.Logger
	token      token.Token
	repository repository.Repository
	app        *fiber.App
}

func New(lg *zap.Logger, token token.Token, repo repository.Repository) *Server {
	s := &Server{logger: lg, token: token, repository: repo}

	s.app = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	v1 := s.app.Group("api/v1")

	auth := v1.Group("auth")
	auth.Post("/register", s.register)
	auth.Post("/login", s.login)

	contacts := v1.Group("contacts", s.fetchUserId)
	contacts.Get("/", s.needsAuthentication)

	return s
}

func (server *Server) Serve(port int) error {
	addr := fmt.Sprintf(":%d", port)
	if err := server.app.Listen(addr); err != nil {
		server.logger.Error("error resolving server", zap.Error(err))
		return err
	}
	return nil
}
