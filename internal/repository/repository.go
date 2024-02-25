package repository

import (
	"embed"
	"fmt"
	"github.com/alireza-fa/phone-book/internal/models"
	"github.com/alireza-fa/phone-book/pkg/rdbms"
	"github.com/alireza-fa/phone-book/pkg/utils"
	"go.uber.org/zap"
	"io/fs"
	"strings"
)

type Repository interface {
	Migrate(migrate models.Migrate) error

	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	GetUserByEmailAndPassword(email, password string) (*models.User, error)
}

type repository struct {
	logger *zap.Logger
	config *Config
	rdbms  rdbms.RDBMS
}

func New(lg *zap.Logger, cfg *Config, rdbms rdbms.RDBMS) Repository {
	return &repository{
		logger: lg,
		config: cfg,
		rdbms:  rdbms,
	}
}

//go:embed migrations
var migrations embed.FS

func (r *repository) Migrate(direction models.Migrate) error {
	files, err := fs.ReadDir(migrations, "migrations")
	if err != nil {
		return fmt.Errorf("error reading migrations directory:\n%v", err)
	}

	result := make([]string, 0, len(files)/2)

	for _, file := range files {
		splits := strings.Split(file.Name(), ".")
		if splits[1] == string(direction) {
			result = append(result, file.Name())
		}
	}

	result = utils.Sort(result)

	for index := 0; index < len(result); index++ {
		file := "migrations/"

		if direction == models.Up {
			file += result[index]
		} else {
			file += result[len(result)-index-1]
		}

		data, err := fs.ReadFile(migrations, file)
		if err != nil {
			return fmt.Errorf("error reading migrations file: %s\n%v", file, err)
		}

		if err := r.rdbms.Execute(string(data), []any{}); err != nil {
			return fmt.Errorf("error migrating the file: %s\n%v", file, err)
		}
	}

	return nil
}
