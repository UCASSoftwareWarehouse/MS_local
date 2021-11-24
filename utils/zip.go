package utils

import (
	"MS_Local/config"
	"archive/zip"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Zip srcFile could be a single file or a directory
func Zip(srcFile string, destZip string) (string, error) {
	var zipfile *os.File
	var err error
	if destZip == "" {
		zipfile, err = os.CreateTemp(config.TempFilePath, "dzip_")
	} else {
		zipfile, err = os.Create(destZip)
	}

	if err != nil {
		return "", err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		rpath, _ := filepath.Rel(srcFile, path)
		f, err := archive.Create(rpath)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}
	err = filepath.Walk(srcFile, walker)

	return zipfile.Name(), err
}

func Unzip(zipFile string, destDir string) (string, error) {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return "", err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return "", err
			}

			inFile, err := f.Open()
			if err != nil {
				return "", err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return "", err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return "", err
			}
		}
	}
	files, err := os.ReadDir(destDir)
	if err != nil {
		return "", err
	}
	if len(files) != 1 {
		log.Printf("unzip should contain only directory, not %d", len(files))
		return "", status.Errorf(codes.Internal, fmt.Sprintf("unzip should contain only directory, not %d", len(files)))
	}

	fpath := filepath.Join(destDir, files[0].Name())
	log.Printf("unzip success, file path is %v", fpath)
	return fpath, nil
}
