package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/op/go-logging"
)

var (
	// Version is set at build time via ldflags
	Version  = "0.0.1"
	// BuildTime is set at build time via ldflags
	BuildTime = "unknown"
)

func main() {
	// Parse command-line flags
	var (
		showVersion = flag.Bool("version", false, "Print version information and exit")
		configFile  = flag.String("config", "", "Path to configuration file")
		// Changed default log level to debug for easier local development
		logLevel    = flag.String("log-level", "debug", "Log level: debug, info, warning, error")
	)
	flag.Parse()

	// Handle version flag
	if *showVersion {
		fmt.Printf("s-ui version %s (built %s)\n", Version, BuildTime)
		os.Exit(0)
	}

	// Initialize logger
	logger := setupLogger(*logLevel)
	logger.Infof("Starting s-ui %s", Version)

	// Load configuration
	cfg, err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize and start the application
	app, err := NewApp(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	if err := app.Start(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

// setupLogger configures the application logger with the given level.
func setupLogger(level string) *logging.Logger {
	logger := logging.MustGetLogger("s-ui")

	backend := logging.NewLogBackend(os.Stdout, "", 0)
	// Updated timestamp format to include milliseconds for easier debugging
	// Also added shortfile info to help trace log sources during development
	format := logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05.000} %{level:.4s} %{id:03x}%{color:reset} %{shortfile} %{message}`,
	)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	var lvl logging.Level
	switch level {
	case "debug":
		lvl = logging.DEBUG
	case "warning":
		lvl = logging.WARNING
	case "error":
		lvl = logging.ERROR
	default:
		lvl = logging.INFO
	}

	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(lvl, "")
	logging.SetBackend(backendLeveled)

	return logger
}
