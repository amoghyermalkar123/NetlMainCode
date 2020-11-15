package handler

import (
	"netl/pkg/services"

	"github.com/rs/zerolog"
)

// Handler type embeds the services of this project and a local logger
type Handler struct {
	userSvc  services.UserService
	boardSvc services.BoardService
	logger   zerolog.Logger
}
