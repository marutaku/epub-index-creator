package expand

import (
	"encoding/xml"
	"errors"
	"fmt"
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
		contents := make([]byte, file.FileInfo().Size())
		if _, err := io.ReadFull(fileReader, contents); err != nil {
			return "", err
		}
		if _, err = targetFile.Write(contents); err != nil {
			return "", err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, targetFile); err != nil {
			return "", err
		}
	}
	return tempDir, nil
}

func FindOPFFilePath(tempDir string) (string, error) {
	containerFilePath := path.Join(tempDir, "META-INF/container.xml")
	fileContents, err := os.ReadFile(containerFilePath)
	if err != nil {
		return "", err
	}
	container := &Container{}
	if err := xml.Unmarshal(fileContents, container); err != nil {
		return "", fmt.Errorf("error unmarshalling container.xml: %v", err)
	}
	if len(container.RootFiles) == 0 {
		return "", errors.New("rootfiles not found")
	}
	rootfile := container.RootFiles[0]
	return path.Join(tempDir, rootfile.FullPath), nil
}
