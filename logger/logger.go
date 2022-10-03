package logger

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	infoLevel  = "INFO"
	errorLevel = "ERROR"
)

// NewLogger fund
func NewLogger() Logger {
	var log Logger
	return log
}

// Parameters struct
type Parameters struct {
	Level                string `json:"level,omitempty"`
	Error                string `json:"error,omitempty"`
	Strategy             string `json:"strategy,omitempty"`
	Input_queue_version  string `json:"input_queue_version,omitempty"`
	Output_queue_version string `json:"output_queue_version,omitempty"`
	Input_topic          string `json:"input_topic,omitempty"`
}

type parametersKeyValue map[string]string

// Logger struct
type Logger struct {
	log *zap.Logger
}

func (logger *Logger) restore() {
	logger.log = initLoggerZap()
}

func (logger *Logger) withFields(parameters parametersKeyValue) {
	for key, value := range parameters {
		logger.log = logger.log.With(zap.String(key, string(value)))
	}
}

// Info log
func (logger *Logger) Info(message string, parameters Parameters) {
	logger.restore()
	parameters.Level = infoLevel
	parameters.Strategy = os.Getenv("STRATEGY")
	logger.withFields(convertParametersToKeyValue(parameters))
	logger.log.Info(message)
}

// Error log
func (logger *Logger) Error(message string, parameters Parameters) {
	logger.restore()
	parameters.Level = errorLevel
	parameters.Strategy = os.Getenv("STRATEGY")
	logger.withFields(convertParametersToKeyValue(parameters))
	logger.log.Error(message)
}

func initLoggerZap() *zap.Logger {
	config := zap.Config{
		Encoding:         "json",
		DisableCaller:    true,
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ := config.Build()
	return log
}

func convertParametersToKeyValue(parameters Parameters) parametersKeyValue {
	var parametersKeyValue map[string]string

	parametersJSON, _ := json.Marshal(parameters)
	json.Unmarshal(parametersJSON, &parametersKeyValue)
	return parametersKeyValue
}
