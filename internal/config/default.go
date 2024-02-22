package config

import (
	"github.com/alireza-fa/phone-book/pkg/logger"
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
		Token: &token.Config{
			PublicPem:  "-----BEGIN PUBLIC KEY-----\nMCowBQYDK2VwAyEAJiIlevPkjU0KhKAc2TO78tQ42kjUocxpgjEI3wp+WTY=\n-----END PUBLIC KEY-----",
			PrivatePem: "-----BEGIN PRIVATE KEY-----\nMC4CAQAwBQYDK2VwBCIEIAndFawSGPx2G5nnyLCXhF1jlaK7PCOL2gekpjU3dFUu\n-----END PRIVATE KEY-----",
			Expiration: time.Minute * 15,
		},
	}
}
