package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func Init(logLevel ...int) (err error) {
	level := 0
	if len(logLevel) > 0 {
		level = logLevel[0]
	}
	Logger, err = NewLogger(level)
	return
}

func NewLogger(level int) (sugar *zap.SugaredLogger, err error) {
	config := zap.Config{
		Level:    zap.NewAtomicLevelAt(zapcore.Level(level)),
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:          "T",
			LevelKey:         "L",
			NameKey:          "N",
			CallerKey:        "C",
			FunctionKey:      zapcore.OmitKey,
			MessageKey:       "M",
			StacktraceKey:    "S",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      definedCapitalColorLevelEncoder,
			EncodeTime:       zapcore.ISO8601TimeEncoder,
			EncodeDuration:   zapcore.StringDurationEncoder,
			EncodeCaller:     definedShortCallerEncoder,
			ConsoleSeparator: " ",
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := config.Build(zap.AddStacktrace(zap.DPanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Color represents a text color.
type Color uint8

// Add adds the coloring to the given string.
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}

var _levelToColor = map[zapcore.Level]Color{
	zapcore.DebugLevel:  Magenta,
	zapcore.InfoLevel:   Blue,
	zapcore.WarnLevel:   Yellow,
	zapcore.ErrorLevel:  Red,
	zapcore.DPanicLevel: Red,
	zapcore.PanicLevel:  Red,
	zapcore.FatalLevel:  Red,
}

func definedCapitalColorLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	color, ok := _levelToColor[l]
	if !ok {
		color = Yellow
	}
	s := color.Add("[" + CapitalString(l) + "]")
	enc.AppendString(s)
}

func CapitalString(l zapcore.Level) string {
	// Printing levels in all-caps is common enough that we should export this
	// functionality.
	switch l {
	case zapcore.DebugLevel:
		return "D"
	case zapcore.InfoLevel:
		return "I"
	case zapcore.WarnLevel:
		return "W"
	case zapcore.ErrorLevel:
		return "E"
	case zapcore.DPanicLevel:
		return "P"
	case zapcore.PanicLevel:
		return "P"
	case zapcore.FatalLevel:
		return "F"
	default:
		return fmt.Sprintf("LEVEL(%d)", l)
	}
}

func definedShortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

func Debug(args ...any) {
	Logger.Debug(args...)
}

func Debugf(template string, args ...any) {
	Logger.Debugf(template, args...)
}

func Info(args ...any) {
	Logger.Info(args...)
}

func Infof(template string, args ...any) {
	Logger.Infof(template, args...)
}

func Error(args ...any) {
	Logger.Error(args...)
}

func Errorf(template string, args ...any) {
	Logger.Errorf(template, args...)
}

func Warn(args ...any) {
	Logger.Warn(args...)
}

func Warnf(template string, args ...any) {
	Logger.Warnf(template, args...)
}

func Panic(args ...any) {
	Logger.Panic(args...)
}

func Panicf(template string, args ...any) {
	Logger.Panicf(template, args...)
}

func DPanic(args ...any) {
	Logger.DPanic(args...)
}

func DPanicf(template string, args ...any) {
	Logger.DPanicf(template, args...)
}

func Fatal(args ...any) {
	Logger.Fatal(args...)
}

func Fatalf(template string, args ...any) {
	Logger.Fatalf(template, args...)
}
