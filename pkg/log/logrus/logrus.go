package logrus

import (
	"errors"
	"github.com/alecthomas/units"
	rotators "github.com/lestrrat-go/file-rotatelogs"
	"github.com/seed95/shortening/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/xhit/go-str2duration/v2"
	"io"
	"path/filepath"
)

var (
	errorNilOption    = errors.New("option can not be nil")
	errorEmptyPath    = errors.New("empty path for log")
	errorEmptyPattern = errors.New("empty pattern for log")
)

type (
	logBundle struct {
		logger *logrus.Logger
	}

	Option struct {
		Path, Pattern, RotationSize, RotationTime, MaxAge string
	}
)

//New is a constructor for logger from 'logrus' package
func New(opt *Option) (log.Logger, error) {

	if opt == nil {
		return nil, errorNilOption
	}

	l := &logBundle{logger: logrus.New()}

	writer, err := getLoggerWriter(opt)
	if err != nil {
		return nil, err
	}

	l.logger.SetOutput(writer)
	l.logger.SetFormatter(&logrus.JSONFormatter{})

	return l, nil
}

func getLoggerWriter(opt *Option) (io.Writer, error) {

	if len(opt.Path) == 0 {
		return nil, errorEmptyPath
	}

	if len(opt.Pattern) == 0 {
		return nil, errorEmptyPattern
	}

	rotationSize, err := units.ParseBase2Bytes(opt.RotationSize)
	if err != nil {
		return nil, err
	}

	maxAge, err := str2duration.ParseDuration(opt.MaxAge)
	if err != nil {
		return nil, err
	}

	rotationTime, err := str2duration.ParseDuration(opt.RotationTime)
	if err != nil {
		return nil, err
	}

	return rotators.New(
		filepath.Join(opt.Path, opt.Pattern),
		rotators.WithRotationSize(int64(rotationSize)),
		rotators.WithMaxAge(maxAge),
		rotators.WithRotationTime(rotationTime),
	)
}

func (l *logBundle) Info(field *log.Field) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Info(field.Message)
}

func (l *logBundle) Warning(field *log.Field) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Warning(field.Message)
}

func (l *logBundle) Error(field *log.Field) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Error(field.Message)
}
