package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"syscall"

	"github.com/mcosta74/change-me/endpoints"
	"github.com/mcosta74/change-me/service"
	"github.com/mcosta74/change-me/transport"
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

	httpAddr string
)

func init() {
	fs = flag.NewFlagSet(AppName, flag.ExitOnError)

	fs.BoolVar(&showVersion, "v", false, "Print version and exit")
	fs.BoolVar(&showBuildInfo, "V", false, "Print build information and exit")

	fs.StringVar(&logFormat, "log.format", getStringEnv("APP_LOG_FORMAT", "text"), "Log format (text, json)")
	fs.StringVar(&logLevel, "log.level", getStringEnv("APP_LOG_LEVEL", slog.LevelInfo.String()), "Log level (debug, info, warn, error)")

	fs.StringVar(&debugAddr, "debug.listen-addr", getStringEnv("APP_DEBUG_LISTEN_ADDR", ":8081"), "Address of the debug HTTP server")
	fs.StringVar(&healthPath, "health.path", getStringEnv("APP_HEALTH_PATH", "/health"), "Path of the health endpoint")
	fs.StringVar(&metricsPath, "metrics.path", getStringEnv("APP_METRICS_PATH", "/metrics"), "Path of the metrics endpoint")

	fs.StringVar(&httpAddr, "http.listen-addr", getStringEnv("APP_HTTP_LISTEN_ADDR", ":8080"), "Address of the HTTP server")
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

	logger := initLogger(logLevel, logFormat)

	logger.Info("started")
	defer logger.Info("stopped")

	var (
		debugHandler = makeDebugHandler()

		svc         = service.NewItemService()
		eps         = endpoints.MakeEndpoints(svc)
		httpHandler = transport.MakeHTTPHandler(eps)
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
		logger.Info("metrics available", "addr", fmt.Sprintf("http://%s%s", l.Addr(), metricsPath))
		logger.Info("health status available", "addr", fmt.Sprintf("http://%s%s", l.Addr(), healthPath))

		g.Add(func() error {
			return http.Serve(l, debugHandler)
		}, func(err error) {
			l.Close()
		})
	}
	{
		// HTTP Handler
		l, err := net.Listen("tcp", httpAddr)
		if err != nil {
			logger.Error("fail to listen", "addr", debugAddr, "err", err)
		}
		g.Add(func() error {
			logger.Info("HTTP server running", "addr", fmt.Sprintf("http://%s", l.Addr()))
			return http.Serve(l, httpHandler)
		}, func(err error) {
			l.Close()
		})
	}
	logger.Info("shutdown", "err", g.Run())
}

func makeDebugHandler() http.Handler {
	h := http.NewServeMux()

	h.HandleFunc(healthPath, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	h.Handle(metricsPath, promhttp.Handler())

	return h
}
