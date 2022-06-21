package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

// Inside the tgz archive, there is another a tar file. Decompress both into the selected directory.
func DecompressTgz(tgzPath, destPath string) error {
	// Open the tgz file
	tarGzFile, err := os.Open(tgzPath)
	if err != nil {
		return err
	}
	defer tarGzFile.Close()

	// Create a gzip reader from the tgz file
	gzipReader, err := gzip.NewReader(tarGzFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	// Create a tar reader from the gzip reader
	tarReader := tar.NewReader(gzipReader)

	// Iterate through the files in the tgz archive
	for {
		// Get the next file in the archive
		header, err := tarReader.Next()
		if err == io.EOF {
			// End of tar archive
			break
		}
		if err != nil {
			return err
		}

		// Create the file or directory in the destination path
		path := filepath.Join(destPath, header.Name)
		if header.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			_, err := os.Stat(path)
			if err != nil {
				// Folder does not exist, create it
				err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
				if err != nil {
					return err
				}
			}

			// Create the file
			file, err := os.Create(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// Write the file contents to the new file
			_, err = io.Copy(file, tarReader)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
