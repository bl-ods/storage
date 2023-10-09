package storage

import "github.com/Gear19/STORAGE/internal/file"

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) (file.File, error) {
	// newFile, err :=
	return file.NewFile(filename, blob)
	// if err != nil {
	// 	return nil,err
	// }
	// return newFile, nil
	// poshel nahui
}
