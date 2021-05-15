package logger

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestSetLoggerLogLevel(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		logLevel   string
		expected   string
		loggerType string
	}{
		{
			logLevel:   "info",
			expected:   "info",
			loggerType: "",
		},
		{
			logLevel:   "warn",
			expected:   "warn",
			loggerType: "",
		},
		{
			logLevel:   "debug",
			expected:   "debug",
			loggerType: "",
		},
		{
			logLevel:   "error",
			expected:   "error",
			loggerType: "",
		},
		{
			logLevel:   "fatal",
			expected:   "fatal",
			loggerType: "",
		},
		{
			logLevel:   "trace",
			expected:   "trace",
			loggerType: "",
		},
		{
			logLevel:   "panic",
			expected:   "panic",
			loggerType: "",
		},
		{
			logLevel:   "plop",
			expected:   "info",
			loggerType: "",
		},
		{
			logLevel:   "plop",
			expected:   "info",
			loggerType: "shell",
		},
	}

	for _, tc := range tests {
		if tc.loggerType != "" {
			os.Setenv("LOGGER_TYPE", tc.loggerType)
			defer os.Unsetenv("LOGGER_TYPE")
		}
		os.Setenv("LOG_LEVEL", tc.logLevel)
		defer os.Unsetenv("LOG_LEVEL")
		SetLoggerLogLevel()
		z := zerolog.GlobalLevel().String()

		assert.Equal(tc.expected, z)
	}
}

func TestLogger_info(t *testing.T) {
	SetLoggerLogLevel()
	log.Info().Msgf("Testing logger")
}
