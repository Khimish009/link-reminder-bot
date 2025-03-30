package files

import (
	"encoding/gob"
	"link-reminder-bot/lib/e"
	"link-reminder-bot/storage"
	"os"
	"path/filepath"
)

const defaultPerm = 0774

type Storage struct {
	baseUrl string
}

func New(baseUrl string) Storage {
	return Storage{
		baseUrl: baseUrl,
	}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("can't save", err) }()

	fPath := filepath.Join(s.baseUrl, page.UserName)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	fileName, err := fileName(page)

	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fileName)

	file, err := os.Create(fPath)

	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func fileName(page *storage.Page) (string, error) {
	return page.Hash()
}