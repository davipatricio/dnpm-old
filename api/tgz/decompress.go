package tgz

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func LoadTgzAndExtractTo(compressedFile string, destination string) error {
	// Open the outer tar.gz file
	f, err := os.Open(compressedFile)
	if err != nil {
		return err
	}

	// Create a new gzip reader on top of the file
	gz, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	// Create a new tar reader on top of the gzip reader
	tr := tar.NewReader(gz)

	// Ensure the destination directory exists
	os.MkdirAll(destination, 0755)

	// Iterate through the files in the outer tar archive
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of outer tar archive
			break
		}
		if err != nil {
			return err
		}

		// Extract the file to the destination directory
		err = extractFile(hdr, tr, destination)
		if err != nil {
			return err
		}
	}

	if err = f.Close(); err != nil {
		return err
	}

	if err = gz.Close(); err != nil {
		return err
	}

	return err
}

func extractFile(hdr *tar.Header, r io.Reader, destination string) error {
	// create paths if they don't exist abc/xyz/123.js -> abc/xyz
	os.MkdirAll(filepath.Join(destination, filepath.Dir(hdr.Name)), 0755)

	// Read the file data and write it to bytes var
	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(destination, hdr.Name), bytes, 0644)

	return err
}
