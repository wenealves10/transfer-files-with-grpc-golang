package storage

import "os"

type Manager interface {
	Store(file *File) error
}

var _ Manager = &Storage{}

type Storage struct {
	dir string
}

func New(dir string) Storage {
	return Storage{
		dir: dir,
	}
}

func (s Storage) Store(file *File) error {
	if err := os.WriteFile(s.dir+file.name, file.buffer.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}
