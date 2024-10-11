package logging

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger

// TODO: put it in another package (probably, name it "middleware")
func NewLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	f, err := os.Open("./log/logs.txt")
	if os.IsNotExist(err) {
		f, err = os.Create("./log/logs.txt")
	}
	defer f.Close()

	*Logger = zerolog.New(f)
}
