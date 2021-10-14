package ez_log

import (
	"io"
	"time"

	"github.com/hxsf/ez-kit/buffer-pool"
)

type rawFormatter struct{}

func NewRawFormatter() LogFormatter {
	return &rawFormatter{}
}

const (
	layout     = "2006-01-02 15:04:05.000"
	timeLength = len(layout)
	sep        = ' '
)

func (r rawFormatter) FormatWrite(w io.Writer, time time.Time, level Level, category *Category, fields []Field, ext ...Field) {
	buf := buffer_pool.Get()
	defer buf.Free()
	totalLen := timeLength + len(EOL) + LevelMaxLength
	if category == nil {
		totalLen += len(Emptycategory)
	} else {
		totalLen += category.ByteLength()
	}
	for _, field := range fields {
		totalLen += field.ByteLength()
	}
	for _, field := range ext {
		totalLen += field.ByteLength()
	}
	totalLen += len(fields) + len(ext) - 1
	buf.Prealloc(totalLen)
	buf.AppendTime(time, layout)
	buf.AppendByte(sep)
	buf.AppendString(level.String())
	buf.AppendByte(sep)
	if category == nil {
		buf.AppendString(Emptycategory)
	} else {
		category.serializeToBuffer(buf)
	}
	for _, field := range fields {
		buf.AppendByte(sep)
		if keyField, ok := (field).(KeyField); ok {
			buf.AppendString(keyField.Key())
			buf.AppendString("=")
		}
		_, _ = field.WriteTo(buf)
	}
	for _, field := range ext {
		buf.AppendByte(sep)
		_, _ = field.WriteTo(buf)
	}
	buf.AppendString(EOL)
	_, _ = w.Write(buf.Bytes())
}

var _ LogFormatter = (*rawFormatter)(nil)
