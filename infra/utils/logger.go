package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

var (
	messageKey = "message"
	levelKey   = "log.level"
	timeKey    = "@timestamp"
	callerKey  = "caller"
)

var logger *zap.Logger

// initLogger initializes the global logger based on environment specific configuration.
func initLogger() *zap.Logger {
	// Check if we are running locally
	env, envPresent := os.LookupEnv("SERV_HUNT_ENV")
	if !envPresent {
		env = "not-local"
	}
	// Check if we are running in a CI server environment like Gitlab.
	_, inCiEnv := os.LookupEnv("CI_SERVER")
	if inCiEnv {
		env = "test"
	}

	var cfg zap.Config

	serviceFields := map[string]interface{}{
		"name":        "servhunt",
		"version":     "1.0.0",
		"environment": env,
		"tier":        "backend",
	}

	initialFields := map[string]interface{}{
		"service": serviceFields,
	}

	stdout := []string{"stdout"}
	stderr := []string{"stderr"}
	level := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	dev := false
	encoding := "json"

	if env == "local" || env == "test" {
		dev = true
		encoding = "console"
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	cfg = zap.Config{
		Level:             level,
		Development:       dev,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          encoding,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: messageKey,

			LevelKey:    levelKey,
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    timeKey,
			EncodeTime: zapcore.RFC3339NanoTimeEncoder,

			CallerKey:    callerKey,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      stdout,
		ErrorOutputPaths: stderr,
		InitialFields:    initialFields,
	}

	newLogger, err := cfg.Build()

	if err != nil {
		panic(err)
	}

	defer newLogger.Sync()

	newLogger.Info("Logger successfully initialized")

	return newLogger
}

// LogHTTPRequestInfo it logs all http requests data.
func LogHTTPRequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := GetRootLogger()
		r := c.Request
		var requestFields []zapcore.Field
		headers, err := json.Marshal(r.Header)
		if err == nil {
			requestFields = append(requestFields, zap.ByteString("http.request.headers", headers))
		} else {
			logger.Error("Failed to marshal headers", zap.NamedError("error.message", err))
		}
		net, _ := c.RemoteIP()
		urlPath := c.Request.URL.Path
		buf, _ := ioutil.ReadAll(c.Request.Body)
		readerForLogging := ioutil.NopCloser(bytes.NewBuffer(buf))
		readerForProcessing := ioutil.NopCloser(bytes.NewBuffer(buf))

		body, err := readBody(readerForLogging)
		if err == nil {
			requestFields = append(requestFields,
				zap.String("http.request.body.content", body),
				zap.String("http.request.source.ip.address", net.String()))
		} else {
			logger.Error("Failed to read request body", zap.NamedError("error.message", err))
		}

		message := fmt.Sprintf("Decoded HTTP Request for endpoint: %s", urlPath)
		logger.With(requestFields...).Info(message)

		c.Request.Body = readerForProcessing
		c.Next()
		var responseFields []zapcore.Field

		responseFields = append(responseFields, zap.String("http.request.uri", urlPath),
			zap.Int("http.response.status.code", c.Writer.Status()))
		logger.With(responseFields...).Info(fmt.Sprintf("HTTP response info for endpoint: %s", urlPath))
	}

}

func readBody(reader io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		GetRootLogger().Error("An error occurred while reading request body", zap.Error(err))
		return "", nil
	}
	return buf.String(), err
}

// once ensures that the logger is initialized only a single time by
// the first call to the GetRootLogger method.
var once sync.Once

// GetRootLogger is a singleton pattern method to get and or initialize the global logger.
func GetRootLogger() *zap.Logger {
	once.Do(func() {
		logger = initLogger()
	})
	return logger
}
