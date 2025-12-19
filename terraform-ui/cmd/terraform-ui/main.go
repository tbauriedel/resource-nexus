package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/tbauriedel/terraform-ui/internal/config"
	"github.com/tbauriedel/terraform-ui/internal/logging"
	"github.com/tbauriedel/terraform-ui/internal/utils/fileutils"
)

var (
	configPath string
	conf       config.Config
	logger     *logging.Logger
	err        error
)

func init() {
	// Set locale to C to avoid translations in command output
	_ = os.Setenv("LANG", "C")

	// Set slog defaults
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	slog.SetDefault(slog.New(handler))

	// Add and parse flags
	flag.StringVar(&configPath, "config", "config_test.json", "Config file")

	flag.Parse()

	// Load server config
	if !fileutils.FileExists(configPath) {
		slog.Info("no config file found. loading defaults")
		conf = config.LoadDefaults()
	} else {
		slog.Info("loading config file", "file", configPath)
		conf, err = config.LoadFromJSONFile(configPath)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}
}

func main() {
	// init logger
	if conf.Logging.Type == "file" {
		// open logfile
		f, err := fileutils.OpenFile(conf.Logging.File)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		defer f.Close()

		// create logger for type 'file'
		logger = logging.NewLoggerFile(conf.Logging, f)
	} else {
		// create logger for type 'stdout'
		logger = logging.NewLoggerStdout(conf.Logging)
	}

	// Get redacted config for initial log output
	redacted, err := json.Marshal(conf.GetConfigRedacted())
	if err != nil {
		logger.Error(fmt.Sprintf("failed to marshal config: %s", err.Error()))
		os.Exit(1)
	}

	logger.Info("starting terraform-ui", "config", json.RawMessage(redacted))
}
