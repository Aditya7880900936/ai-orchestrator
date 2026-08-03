package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init() error {

	cfg := zap.NewProductionConfig()

	cfg.Encoding = "json"

	cfg.OutputPaths = []string{
		"stdout",
	}

	cfg.ErrorOutputPaths = []string{
		"stderr",
	}

	cfg.DisableCaller = false
	cfg.DisableStacktrace = false

	var err error

	Log, err = cfg.Build()
	if err != nil {
		return err
	}

	return nil
}
