package pricerepository

import (
	"os"
)

func (r *repository) Insert(filename, content string) error {

	f, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)

	if err != nil {
		return err
	}

	return nil
}
