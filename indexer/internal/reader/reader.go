package reader

import (
	"io"
	"os"
	"path"
	"path/filepath"

	"archive/zip"
)

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func UnzipEPub(filepath string) (string, error) {
	filename := getFileNameWithoutExt(filepath)
	tempDir, err := os.MkdirTemp("", filename)
	if err != nil {
		return "", err
	}
	reader, err := zip.OpenReader(filepath)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}
		fileReader, err := file.Open()
		if err != nil {
			return "", err
		}
		defer fileReader.Close()

		targetPath := path.Join(tempDir, file.Name)
		if err := os.MkdirAll(path.Dir(targetPath), os.ModePerm); err != nil {
			return "", err
		}
		targetFile, err := os.Create(targetPath)
		if err != nil {
			return "", err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, targetFile); err != nil {
			return "", err
		}
	}
	return tempDir, nil
}
