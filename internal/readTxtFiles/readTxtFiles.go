package readtxtfiles

import (
	"io/fs"
	"os"
	"path/filepath"
)

func ReadTxtFiles(pathToDir string) (map[string]string, error) {
	files := make(map[string]string)
	err := filepath.Walk(pathToDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".txt" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			files[filepath.Base(path)] = string(content)
		}
		return nil
	})

	return files, err
}

func GetTxtFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(d.Name()) == ".txt" {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
