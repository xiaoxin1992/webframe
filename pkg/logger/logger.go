package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
	"time"
	"webframe/conf"
)

var Logger *LogZap

func NewLogger(rootPath string) (err error) {
	logConfig := conf.GetConfig().Log
	app := conf.GetConfig().App.Name
	encoderConfig := zap.NewProductionEncoderConfig()
	if logConfig.LocalTime {
		encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format("2006-01-02 15:04:05"))
		}
	} else {
		encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.UTC().Format("2006-01-02 15:04:05"))
		}
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if logConfig.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}
	level, err := zapcore.ParseLevel(logConfig.Level)
	if err != nil {
		return
	}
	Logger = &LogZap{
		name:    app,
		level:   level,
		encoder: encoder,
	}
	logPath := logConfig.Path
	if !path.IsAbs(logConfig.Path) {
		logPath = path.Join(rootPath, logConfig.Path)
	}
	file := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    logConfig.MaxSize,
		MaxBackups: logConfig.MaxBackup,
		MaxAge:     logConfig.MaxBackup,
		Compress:   logConfig.Compress,
		LocalTime:  logConfig.LocalTime,
	}
	if logConfig.Console {
		Logger.sugarLog(os.Stdout, file)
	} else {
		Logger.sugarLog(file)
	}
	return
}

type LogZap struct {
	f           *os.File
	name        string
	level       zapcore.LevelEnabler
	encoder     zapcore.Encoder
	sugarLogger *zap.SugaredLogger
}

func (l *LogZap) sugarLog(writes ...io.Writer) {
	core := l.core(writes...)
	l.sugarLogger = zap.New(core, zap.AddCaller()).Named(l.name).Sugar()
}

func (l *LogZap) core(writes ...io.Writer) zapcore.Core {
	cores := make([]zapcore.Core, 0)
	for _, w := range writes {
		cores = append(cores, zapcore.NewCore(l.encoder, zapcore.AddSync(w), l.level))
	}
	return zapcore.NewTee(cores...)
}
func (l *LogZap) S(namespace string) *zap.SugaredLogger {
	return l.sugarLogger.Named(namespace)
}

func (l *LogZap) Sync() (err error) {
	err = l.sugarLogger.Sync()
	return
}
