package services

import (
	db "netl/pkg/database"

	"github.com/rs/zerolog"
)

// BoardService type embeds the database type to create sessions and database queries
type BoardService struct {
	dtb    db.DB
	logger zerolog.Logger
}
