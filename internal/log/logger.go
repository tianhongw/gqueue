package log

import (
	"errors"
	"strings"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Warningln(args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	Level() Level
	SetLevel(l Level)

	V(l int) bool
	Flush() error
}

type Level int

const (
	LevelDebug = iota - 1
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

type Format int

const (
	FormatUniversal = iota
	FormatJSON
	FormatPlain
)

type LoggerType string

const (
	LoggerTypeStd = "std"
	LoggerTypeZap = "zap"
)

type options struct {
	Development  bool
	Level        Level
	Format       Format
	Outputs      []string
	ErrorOutputs []string
	Prefix       string
	MaxSize      int
	MaxAge       int
	MaxBackups   int
	Compress     bool
}

type Option func(*options)

func WithLevel(levelText string) Option {
	return func(o *options) {
		var level Level
		switch strings.ToLower(levelText) {
		case "debug":
			level = LevelDebug
		case "info":
			level = LevelInfo
		case "warning":
			level = LevelInfo
		case "error":
			level = LevelError
		case "fatal":
			level = LevelFatal
		default:
			level = LevelInfo
		}
		o.Level = level
	}
}

func WithFormat(formatText string) Option {
	return func(o *options) {
		var format Format
		switch strings.ToLower(formatText) {
		case "json":
			format = FormatJSON
		case "plain", "":
			format = FormatPlain
		default:
			format = FormatPlain
		}
		o.Format = format
	}
}

func WithMaxSize(maxSize int) Option {
	return func(o *options) {
		if maxSize > 0 {
			o.MaxSize = maxSize
		}
	}
}

func WithMaxAge(maxAge int) Option {
	return func(o *options) {
		if maxAge > 0 {
			o.MaxAge = maxAge
		}
	}
}

func WithMaxBackup(maxBackups int) Option {
	return func(o *options) {
		if maxBackups > 0 {
			o.MaxBackups = maxBackups
		}
	}
}

func NewLogger(loggerType LoggerType, opts ...Option) (Logger, error) {
	o := &options{
		Development:  false,
		Level:        LevelInfo,
		Format:       FormatPlain,
		ErrorOutputs: []string{"stderr"},
		MaxSize:      500,
		MaxAge:       7,
		MaxBackups:   10,
		Compress:     true,
	}

	for _, opt := range opts {
		opt(o)
	}

	switch loggerType {
	case LoggerTypeStd:
		return nil, errors.New("unknown logger type")
	case LoggerTypeZap:
		return nil, errors.New("unknown logger type")
	default:
		return nil, errors.New("unknown logger type")
	}
}
