package internal

import (
	"time"

	"github.com/wearelumenai/clusauth/internal/conf"

	sentry "github.com/getsentry/sentry-go"
)

// ApplySentry apply sentry
func ApplySentry(conf conf.Conf) (err error) {
	if conf.Clusauth.Sentry != "" {
		err = sentry.Init(sentry.ClientOptions{
			// Either set your DSN here or set the SENTRY_DSN environment variable.
			Dsn: conf.Clusauth.Sentry,
			// Enable printing of SDK debug messages.
			// Useful when getting started or trying to figure something out.
			Debug: true,
		})
		// Flush buffered events before the program terminates.
		// Set the timeout to the maximum duration the program can afford to wait.
		defer sentry.Flush(2 * time.Second)
	}
	return
}
