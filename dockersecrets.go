package dockersecretsenv

import (
	"io/ioutil"
	"os"
	"strings"

	"io/fs"
)

func LoadSecrets(_path string) error {
	// load files in /run/secrets/ and expose values as environment variables

	root := "/run/secrets"

	if _path != "" {
		root = _path
	}

	if _, err := os.Stat(root); os.IsNotExist(err) {
		return nil
	}

	fileSystem := os.DirFS(root)

	err := fs.WalkDir(fileSystem, ".", func(file string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		envName := strings.Replace(file, "/", "_", -1)
		path := root + "/" + file
		envValue, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		os.Setenv(envName, string(envValue))
		return nil
	})

	return err
}
