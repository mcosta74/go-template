package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"syscall"

	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/exp/slog"
)

const AppName = "go-template"

// build info
var (
	version = "development"
	commit  = "N.A."
	date    = "N.A."
)

// flags
var (
	fs *flag.FlagSet

	showVersion   bool
	showBuildInfo bool

	logLevel  string
	logFormat string

	debugAddr   string
	healthPath  string
	metricsPath string
)

func init() {
	fs = flag.NewFlagSet(AppName, flag.ExitOnError)

	fs.BoolVar(&showVersion, "v", false, "Print version and exit")
	fs.BoolVar(&showBuildInfo, "V", false, "Print build information and exit")

	fs.StringVar(&logFormat, "log.format", getEnv("APP_LOG_FORMAT", "text"), "Log format (text, json)")
	fs.StringVar(&logLevel, "log.level", getEnv("APP_LOG_LEVEL", slog.LevelInfo.String()), "Log level (debug, info, warn, error)")

	fs.StringVar(&debugAddr, "debug.listen", getEnv("APP_DEBUG_LISTEN_ADDR", ":8081"), "Address of the debug HTTP server")
	fs.StringVar(&healthPath, "health.path", getEnv("APP_HEALTH_PATH", "/health"), "Path of the health endpoint")
	fs.StringVar(&metricsPath, "metrics.path", getEnv("APP_METRICS_PATH", "/metrics"), "Path of the metrics endpoint")
}

func main() {
	fs.Parse(os.Args[1:])

	if showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if showBuildInfo {
		fmt.Printf("Version:%s, GitCommit:%s, BuildDate:%s\n", version, commit, date)
		os.Exit(0)
	}

	logger := initLogger()

	logger.Info("started")
	defer logger.Info("stopped")

	var (
		debugHandler = makeDebugHandler()
	)

	var g run.Group
	{
		// Signal Handler
		g.Add(run.SignalHandler(context.Background(), syscall.SIGTERM, syscall.SIGINT))
	}
	{
		// Debug Handler
		l, err := net.Listen("tcp", debugAddr)
		if err != nil {
			logger.Error("fail to listen", "addr", debugAddr, "err", err)
		}
		logger.Info("debug server running",
			"metrics", fmt.Sprintf("http://%s%s", l.Addr(), metricsPath),
			"health", fmt.Sprintf("http://%s%s", l.Addr(), healthPath),
		)

		g.Add(func() error {
			return http.Serve(l, debugHandler)
		}, func(err error) {
			l.Close()
		})
	}
	logger.Info("shutdown", "err", g.Run())
}

func getEnv(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

func initLogger() *slog.Logger {
	var lvl slog.Level
	if err := lvl.UnmarshalText([]byte(logLevel)); err != nil {
		lvl = slog.LevelInfo
	}
	opts := &slog.HandlerOptions{
		Level:     lvl,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				// use UTC time
				a.Value = slog.TimeValue(a.Value.Time().UTC())

			case slog.SourceKey:
				// remove directories from File
				source := a.Value.Any().(*slog.Source)
				source.File = filepath.Base(source.File)
			}
			return a
		},
	}

	var h slog.Handler = slog.NewTextHandler(os.Stdout, opts)
	if logFormat == "json" {
		h = slog.NewJSONHandler(os.Stdout, opts)
	}
	return slog.New(h)
}

func makeDebugHandler() http.Handler {
	h := http.NewServeMux()

	h.HandleFunc(healthPath, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	h.Handle(metricsPath, promhttp.Handler())

	return h
}
