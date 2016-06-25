package file

import (
	"os"
	"path/filepath"
	"strings"
)

func WalkDir(dir string) ([]string, error) {
	var names []string
	e := filepath.Walk(dir, func(path string, finfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".go") && !finfo.IsDir() {
			names = append(names, path)
		}
		return nil
	})

	return names, e
}
