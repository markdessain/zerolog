package zerolog

import(
	"time"
	"os"
	"zerolog/log"
)

func Load() {
	TimeFieldFormat = time.RFC3339Nano
	if os.Getenv("ENVIRONMENT") == "development" {
		log.Logger = log.Output(ConsoleWriter{Out: os.Stderr})
		log.Debug().Msg("Development Mode")

		for _, e := range os.Environ() {
			log.Debug().Msg(e)
		}
	}
}