package utils

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// Set level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Set readable output for development
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}
