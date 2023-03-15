package log

import(
	"time"
	"os"
)

func Load() {
	TimeFieldFormat = time.RFC3339Nano
	if os.Getenv("ENVIRONMENT") == "development" {
		Logger =Output(ConsoleWriter{Out: os.Stderr})
		Debug().Msg("Development Mode")

		for _, e := range os.Environ() {
			Debug().Msg(e)
		}
	}
}
