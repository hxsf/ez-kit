package ez_log

import (
	"encoding/hex"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"time"
)

type Field interface {
	// just write value
	io.WriterTo
	// total key/value in bytes length
	ByteLength() int
}

type KeyField interface {
	Field
	Key() string
}

type keyField struct {
	Field
	key string
}

func (f *keyField) Key() string {
	return f.key
}

func (f *keyField) ByteLength() int {
	return len(f.key) + f.Field.ByteLength()
}

type stringField struct {
	value string
}

func String(value string) Field {
	return &stringField{value: value}
}

func StringKey(key string, value string) Field {
	return &keyField{String(value), key}
}

func (f *stringField) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(f.value))
	return int64(n), err
}

func (f *stringField) ByteLength() int {
	return len(f.value)
}

type stringsField struct {
	value []string
}

func Strings(value []string) Field {
	return &stringsField{value: value}
}

func StringsKey(key string, value []string) Field {
	return &keyField{Strings(value), key}
}

func (f *stringsField) WriteTo(w io.Writer) (total int64, err error) {
	n, err := w.Write([]byte("["))
	total += int64(n)
	if err != nil {
		return
	}
	for i, value := range f.value {
		n, err = w.Write([]byte(value))
		total += int64(n)
		if err != nil {
			return
		}
		if i != len(f.value)-1 {
			n, err = w.Write([]byte(", "))
		} else {
			n, err = w.Write([]byte("]"))
		}
		total += int64(n)
		if err != nil {
			return
		}
	}
	return
}

func (f *stringsField) ByteLength() int {
	n := 0
	for _, value := range f.value {
		n += len(value)
	}
	return len(f.value)*2 + 1 + n
}

type stringerField struct {
	value fmt.Stringer

	_bytes []byte
	_len   int
}

func Stringer(value fmt.Stringer) Field {
	return &stringerField{value: value, _bytes: nil, _len: -1}
}

func StringerKey(key string, value fmt.Stringer) Field {
	return &keyField{Stringer(value), key}
}

func (f *stringerField) cache() {
	f._bytes = []byte(f.value.String())
	f._len = len(f._bytes)
}

func (f *stringerField) ByteLength() int {
	if f._len == -1 {
		f.cache()
	}
	return f._len
}
func (f *stringerField) WriteTo(w io.Writer) (int64, error) {
	if f._len == -1 {
		f.cache()
	}
	n, err := w.Write(f._bytes)
	return int64(n), err
}

type IByteser interface {
	Bytes() []byte
}

type byteserField struct {
	value IByteser

	_bytes []byte
	_len   int
}

func Byteser(value IByteser) Field {
	return &byteserField{value: value, _bytes: nil, _len: -1}
}
func ByteserKey(key string, value IByteser) Field {
	return &keyField{Byteser(value), key}
}

func (f *byteserField) cache() {
	f._bytes = f.value.Bytes()
	f._len = len(f._bytes)
}
func (f *byteserField) WriteTo(w io.Writer) (int64, error) {
	if f._len == -1 {
		f.cache()
	}
	n, err := w.Write(f._bytes)
	return int64(n), err
}

func (f *byteserField) ByteLength() int {
	if f._len == -1 {
		f.cache()
	}
	return f._len
}

type uintField struct {
	value uint
	len   int
}

func Uint(value uint) Field {
	return &uintField{value: value, len: uintLength(value)}
}

func UintKey(key string, value uint) Field {
	return &keyField{Uint(value), key}
}

func (f *uintField) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(strconv.FormatUint(uint64(f.value), 10)))
	return int64(n), err
}

func (f *uintField) ByteLength() int {
	return f.len
}

type uint8Field struct {
	value uint8
	len   int
}

func Uint8(value uint8) Field {
	return &uint8Field{value: value, len: uint8Length(value)}
}

func (f *uint8Field) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(strconv.FormatUint(uint64(f.value), 10)))
	return int64(n), err
}

func (f *uint8Field) ByteLength() int {
	return f.len
}

type uint16Field struct {
	value uint16
	len   int
}

func Uint16(value uint16) Field {
	return &uint16Field{value: value, len: uint16Length(value)}
}

func (f *uint16Field) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(strconv.FormatUint(uint64(f.value), 10)))
	return int64(n), err
}

func (f *uint16Field) ByteLength() int {
	return f.len
}

type uint32Field struct {
	value uint32
	len   int
}

func Uint32(value uint32) Field {
	return &uint32Field{value: value, len: uint32Length(value)}
}

func (f *uint32Field) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(strconv.FormatUint(uint64(f.value), 10)))
	return int64(n), err
}

func (f *uint32Field) ByteLength() int {
	return f.len
}

type uint64Field struct {
	value uint64
	len   int
}

func Uint64(value uint64) Field {
	return &uint64Field{value: value, len: uint64Length(value)}
}

func (f *uint64Field) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(strconv.FormatUint(uint64(f.value), 10)))
	return int64(n), err
}

func (f *uint64Field) ByteLength() int {
	return f.len
}

type intField struct {
	value int
	len   int
}

func Int(value int) Field {
	return &intField{value: value, len: intLength(value)}
}

func (f *intField) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(strconv.FormatInt(int64(f.value), 10)))
	return int64(n), err
}

func (f *intField) ByteLength() int {
	return f.len
}

type int8Field struct {
	value int8
	len   int
}

func Int8(value int8) Field {
	return &int8Field{value: value, len: int8Length(value)}
}

func (f *int8Field) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(strconv.FormatInt(int64(f.value), 10)))
	return int64(n), err
}

func (f *int8Field) ByteLength() int {
	return f.len
}

type int16Field struct {
	value int16
	len   int
}

func Int16(value int16) Field {
	return &int16Field{value: value, len: int16Length(value)}
}

func (f *int16Field) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(strconv.FormatInt(int64(f.value), 10)))
	return int64(n), err
}

func (f *int16Field) ByteLength() int {
	return f.len
}

type int32Field struct {
	value int32
	len   int
}

func Int32(value int32) Field {
	return &int32Field{value: value, len: int32Length(value)}
}

func (f *int32Field) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(strconv.FormatInt(int64(f.value), 10)))
	return int64(n), err
}

func (f *int32Field) ByteLength() int {
	return f.len
}

type int64Field struct {
	value int64
	len   int
}

func Int64(value int64) Field {
	return &int64Field{value: value, len: int64Length(value)}
}

func (f *int64Field) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(strconv.FormatInt(int64(f.value), 10)))
	return int64(n), err
}

func (f *int64Field) ByteLength() int {
	return f.len
}

type boolField struct {
	value bool
	len   int
}

func Bool(value bool) Field {
	l := 5
	if value {
		l = 4
	}
	return &boolField{value: value, len: l}
}

func (f *boolField) WriteTo(w io.Writer) (int64, error) {
	if f.value {
		n, err := w.Write([]byte("true"))
		return int64(n), err
	}
	n, err := w.Write([]byte("false"))
	return int64(n), err
}

func (f *boolField) ByteLength() int {
	return f.len
}

type byteField struct {
	value byte
}

func Byte(value byte) Field {
	return &byteField{value: value}
}

func (f *byteField) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(strconv.FormatUint(uint64(f.value), 10)))
	return int64(n), err
}

func (f *byteField) ByteLength() int {
	return 1
}

type durationField struct {
	value time.Duration

	_bytes []byte
	_len   int
}

func Duration(value time.Duration) Field {
	return &durationField{value: value, _len: -1}
}

func (f *durationField) WriteTo(w io.Writer) (int64, error) {
	if f._len < 0 {
		f.cache()
	}
	n, err := w.Write(f._bytes)
	return int64(n), err
}

func (f *durationField) cache() {
	f._bytes = []byte(f.value.String())
	f._len = len(f._bytes)
}

func (f *durationField) ByteLength() int {
	if f._len < 0 {
		f.cache()
	}
	return f._len
}

type bytesHexField struct {
	value []byte
}

func BytesHex(value []byte) Field {
	return &bytesHexField{value: value}
}

func (f *bytesHexField) WriteTo(w io.Writer) (int64, error) {
	dst := make([]byte, hex.EncodedLen(len(f.value)))
	hex.Encode(dst, f.value)
	n, err := w.Write(dst)
	return int64(n), err
}

func (f *bytesHexField) ByteLength() int {
	return len(f.value)
}

type bytesStringField struct {
	value []byte
}

func BytesString(value []byte) Field {
	return &bytesStringField{value: value}
}

func (f *bytesStringField) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(f.value)
	return int64(n), err
}

func (f *bytesStringField) ByteLength() int {
	return len(f.value)
}

type errorStringField struct {
	value error
}

type nilErrorField struct{}

const nilString = "<nil>"

func (f *nilErrorField) WriteTo(w io.Writer) (n int64, err error) {
	nn, err := w.Write([]byte(nilString))
	return int64(nn), err
}

func (f *nilErrorField) ByteLength() int {
	return len(nilString)
}

func Error(value error) Field {
	if value == nil {
		return &nilErrorField{}
	}
	return &errorStringField{value: value}
}

func (f *errorStringField) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(f.value.Error()))
	return int64(n), err
}

func (f *errorStringField) ByteLength() int {
	return len(f.value.Error())
}

type timeField struct {
	value time.Time
}

func Time(value time.Time) Field {
	return &timeField{value: value}
}

func (f *timeField) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte(f.value.Format(time.RFC3339Nano)))
	return int64(n), err
}

func (f *timeField) ByteLength() int {
	return len(time.RFC3339Nano)
}

type callerField struct {
	value runtime.Frame
	len   int
}

func Caller(value runtime.Frame) Field {
	return &callerField{
		value: value,
		len:   len(value.File) + len(value.File) + intLength(value.Line) + baseCallerLen,
	}
}

const (
	callerStart   = "File: "
	callerSep     = ":"
	baseCallerLen = len(callerStart) + len(callerSep)
)

func (f *callerField) ByteLength() int {
	return f.len
}

func (f *callerField) WriteTo(w io.Writer) (total int64, err error) {
	n, err := w.Write([]byte(callerStart))
	total += int64(n)
	n, err = w.Write([]byte(f.value.File))
	total += int64(n)
	n, err = w.Write([]byte(callerSep))
	total += int64(n)
	n, err = w.Write([]byte(strconv.Itoa(f.value.Line)))
	total += int64(n)
	return
}

type reflectField struct {
	value interface{}

	_bytes []byte
	_len   int
}

func Reflect(value interface{}) Field {
	return &reflectField{value: value, _len: -1}
}

func (f *reflectField) cache() {
	f._bytes = []byte(fmt.Sprintf("%+v", f.value))
	f._len = len(f._bytes)
}

func (f *reflectField) WriteTo(w io.Writer) (int64, error) {
	if f._len < 0 {
		f.cache()
	}
	n, err := w.Write(f._bytes)
	return int64(n), err
}

func (f *reflectField) ByteLength() int {
	if f._len < 0 {
		f.cache()
	}
	return f._len
}

func Any(value interface{}) Field {
	switch val := value.(type) {
	case bool:
		return Bool(val)
	case int:
		return Int(val)
	case int64:
		return Int64(val)
	case int32:
		return Int32(val)
	case int16:
		return Int16(val)
	case int8:
		return Int8(val)
	case string:
		return String(val)
	case uint:
		return Uint(val)
	case uint64:
		return Uint64(val)
	case uint32:
		return Uint32(val)
	case uint16:
		return Uint16(val)
	case uint8:
		return Uint8(val)
	case []byte:
		return BytesHex(val)
	case time.Time:
		return Time(val)
	case time.Duration:
		return Duration(val)
	case error:
		return Error(val)
	case fmt.Stringer:
		return Stringer(val)
	default:
		return Reflect(val)
	}
}

func AnyKey(key string, value interface{}) Field {
	switch val := value.(type) {
	case bool:
		return BoolKey(key, val)
	case int:
		return IntKey(key, val)
	case int64:
		return Int64Key(key, val)
	case int32:
		return Int32Key(key, val)
	case int16:
		return Int16Key(key, val)
	case int8:
		return Int8Key(key, val)
	case string:
		return StringKey(key, val)
	case uint:
		return UintKey(key, val)
	case uint64:
		return Uint64Key(key, val)
	case uint32:
		return Uint32Key(key, val)
	case uint16:
		return Uint16Key(key, val)
	case uint8:
		return Uint8Key(key, val)
	case []byte:
		return BytesHexKey(key, val)
	case time.Time:
		return TimeKey(key, val)
	case time.Duration:
		return DurationKey(key, val)
	case error:
		return ErrorKey(key, val)
	case fmt.Stringer:
		return StringerKey(key, val)
	default:
		return ReflectKey(key, val)
	}
}

func Uint8Key(key string, value uint8) Field {
	return &keyField{key: key, Field: Uint8(value)}
}

func Uint16Key(key string, value uint16) Field {
	return &keyField{key: key, Field: Uint16(value)}
}

func Uint32Key(key string, value uint32) Field {
	return &keyField{key: key, Field: Uint32(value)}
}

func Uint64Key(key string, value uint64) Field {
	return &keyField{key: key, Field: Uint64(value)}
}

func IntKey(key string, value int) Field {
	return &keyField{key: key, Field: Int(value)}
}

func Int8Key(key string, value int8) Field {
	return &keyField{key: key, Field: Int8(value)}
}

func Int16Key(key string, value int16) Field {
	return &keyField{key: key, Field: Int16(value)}
}

func Int32Key(key string, value int32) Field {
	return &keyField{key: key, Field: Int32(value)}
}

func Int64Key(key string, value int64) Field {
	return &keyField{key: key, Field: Int64(value)}
}

func ByteKey(key string, value byte) Field {
	return &keyField{key: key, Field: Byte(value)}
}

func BytesHexKey(key string, value []byte) Field {
	return &keyField{key: key, Field: BytesHex(value)}
}

func BytesStringKey(key string, value []byte) Field {
	return &keyField{key: key, Field: BytesString(value)}
}

func TimeKey(key string, value time.Time) Field {
	return &keyField{key: key, Field: Time(value)}
}

func ErrorKey(key string, value error) Field {
	return &keyField{key: key, Field: Error(value)}
}

func CallerKey(key string, value runtime.Frame) Field {
	return &keyField{key: key, Field: Caller(value)}
}

func BoolKey(key string, value bool) Field {
	return &keyField{key: key, Field: Bool(value)}
}

func DurationKey(key string, value time.Duration) Field {
	return &keyField{key: key, Field: Duration(value)}
}
func ReflectKey(key string, value interface{}) Field {
	return &keyField{key: key, Field: Reflect(value)}
}
