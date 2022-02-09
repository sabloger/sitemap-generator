package smg

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

// checkAndMakeDir makes the path in case of absence of the OutputPath
func checkAndMakeDir(path string) error {
	if _, err := os.Stat(path); path != "" && os.IsNotExist(err) {
		err := os.MkdirAll(path, 0666)
		if err != nil {
			return err
		}
	}
	return nil
}

// writeToFile uses an io.WriterTo interface param to write the Sitemap file.
// writer param cab be a Sitemap or SitemapIndex instance.
// filename param is a full filename with extension and path is the dir path.
// compress defines whether the file must be gzip compressed or not.
// returns n for number of written bytes and error in case of any problem.
func writeToFile(writer io.WriterTo, filename, path string, compress bool) (int64, error) {
	file, err := os.OpenFile(filepath.Join(path, filename), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	if compress {
		w := gzip.NewWriter(file)
		defer w.Close()

		return writer.WriteTo(w)
	}
	return writer.WriteTo(file)
}
