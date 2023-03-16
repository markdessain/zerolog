package log

import(
	"time"
	"os"
	
	"github.com/markdessain/zerolog"
)

func Load() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	if os.Getenv("ENVIRONMENT") == "development" {
		Logger =Output(ConsoleWriter{Out: os.Stderr})
		Debug().Msg("Development Mode")

		for _, e := range os.Environ() {
			Debug().Msg(e)
		}
	}
}
