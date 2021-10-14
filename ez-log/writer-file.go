package ez_log

import (
	"os"
	"path/filepath"
)

type fileWriter struct {
	out *os.File
}

func prepareFile(file string) (*os.File, error) {
	dir := filepath.Base(file)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
}

func NewFileWriter(logFilepath string) (WriteSyncer, error) {
	outFile, err := prepareFile(logFilepath)
	if err != nil {
		return nil, err
	}
	return &fileWriter{out: outFile}, nil
}

func (f *fileWriter) Write(buf []byte) (int, error) {
	return f.out.Write(buf)
}


func (f *fileWriter) Sync() error {
	return f.out.Sync()
}
