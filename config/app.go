package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/naoina/kocha"
	"github.com/naoina/kocha/event/memory"
	"github.com/naoina/kocha/log"
	"github.com/naoina/vocanew/app/event"
	"github.com/naoina/vocanew/app/middleware"
)

var (
	AppName   = "vocanew"
	AppConfig = &kocha.Config{
		Addr:          kocha.Getenv("VOCANEW_ADDR", "127.0.0.1:9100"),
		AppPath:       rootPath,
		AppName:       AppName,
		DefaultLayout: "app",
		Template: &kocha.Template{
			PathInfo: kocha.TemplatePathInfo{
				Name: AppName,
				Paths: []string{
					filepath.Join(rootPath, "app", "view"),
				},
			},
			FuncMap: kocha.TemplateFuncMap{
				"isDev": func() bool {
					return kocha.Getenv("VOCANEW_ENV", "dev") != "prod"
				},
				"splitFirstNLines": func(s string, n int) []string {
					lines := strings.SplitAfterN(s, "<br />", n)
					result := []string{strings.Join(lines[:n-1], "")}
					if len(lines) == n {
						result = append(result, lines[n-1])
					}
					return result
				},
			},
		},

		// Logger settings.
		Logger: &kocha.LoggerConfig{
			Writer:    os.Stdout,
			Formatter: &log.LTSVFormatter{},
			Level:     log.INFO,
		},

		// Middlewares.
		Middlewares: []kocha.Middleware{
			&kocha.RequestLoggingMiddleware{},
			&kocha.PanicRecoverMiddleware{},
			&kocha.FormMiddleware{},
			&middleware.Vocaloid{},
			&kocha.DispatchMiddleware{},
		},

		// Event.
		Event: &kocha.Event{
			HandlerMap: kocha.EventHandlerMap{
				&memory.EventQueue{}: {
					"vocanew:send_contact_mail": event.SendContactMail,
				},
			},
		},

		MaxClientBodySize: 1024 * 1024 * 10, // 10MB
	}

	_, configFileName, _, _ = runtime.Caller(0)
	rootPath                = filepath.Dir(filepath.Join(configFileName, ".."))
)
