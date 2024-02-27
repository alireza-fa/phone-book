package config

import (
	"github.com/alireza-fa/phone-book/internal/repository"
	"github.com/alireza-fa/phone-book/pkg/logger"
	"github.com/alireza-fa/phone-book/pkg/rdbms"
	"github.com/alireza-fa/phone-book/pkg/token"
	"time"
)

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
			Encoding:    "console",
		},
		RDBMS: &rdbms.Config{
			Host:     "localhost",
			Port:     5433,
			Username: "PHONEBOOK_USER",
			Password: "PHONEBOOK_PASSWORD",
			Database: "PHONEBOOK_DB",
		},
		Repository: &repository.Config{
			CursorSecret: "A?D(G-KapdSgVVVG",
			Limit: struct {
				Min int "koanf:\"min\""
				Max int "koanf:\"max\""
			}{12, 48},
		},
		Token: &token.Config{
			PublicPem:  "-----BEGIN PUBLIC KEY-----\nMCowBQYDK2VwAyEAJiIlevPkjU0KhKAc2TO78tQ42kjUocxpgjEI3wp+WTY=\n-----END PUBLIC KEY-----",
			PrivatePem: "-----BEGIN PRIVATE KEY-----\nMC4CAQAwBQYDK2VwBCIEIAndFawSGPx2G5nnyLCXhF1jlaK7PCOL2gekpjU3dFUu\n-----END PRIVATE KEY-----",
			Expiration: time.Minute * 15,
		},
	}
}
