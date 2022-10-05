package store

import (
	"os"
	"encoding/json"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
    FileType Type = "file"
	MonoType Type = "mongo"
)

func New(store Type, fileName string) Store {
	switch store {
    case FileType:
        return &fileStore{fileName}
	}
	return nil
}

type fileStore struct {
	FilePath string
}

func (f *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "\t")
	if err!= nil {
        return err
    }
	return os.WriteFile(f.FilePath, fileData, 0644)
}

func (f *fileStore) Read(data interface{}) error {
	fileData, err := os.ReadFile(f.FilePath)
    if err!= nil {
        return err
    }
	return json.Unmarshal(fileData, &data)
}