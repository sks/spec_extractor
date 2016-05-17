package extractor

import (
	"os"
	"path/filepath"
	"strings"
)

func FindSpecFiles(specfiles *[]string) filepath.WalkFunc {
	return filepath.WalkFunc(func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Base(path) != "spec" {
			return nil
		}

		if !strings.Contains(path, "jobs") {
			return nil
		}

		*specfiles = append(*specfiles, path)
		return nil
	})
}
