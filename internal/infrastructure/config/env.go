package config

import "github.com/vinnedev/http-server-go-boilerplate/pkg/dotenv"

// Service
var ENV_MODE = dotenv.GetEnv("ENV_MODE", "development")
var PORT = dotenv.GetEnv("PORT", "8080")
