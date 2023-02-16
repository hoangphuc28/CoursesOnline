package middleware

import "github.com/hoangphuc28/CoursesOnline/API-Gateway/config"

type MiddleareManager struct {
	cfg *config.Config
}

func NewMiddlewareManager(cfg *config.Config) *MiddleareManager {
	return &MiddleareManager{cfg}
}
