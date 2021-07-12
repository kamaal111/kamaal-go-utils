package files

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"
)

func AppendFileToPath(path string, file string) string {
	if len(strings.TrimSpace(path)) < 1 {
		return file
	}

	pathRune := []rune(path)
	lastCharacter := string(pathRune[len(pathRune)-1:])
	if lastCharacter == "/" {
		return path + file
	}

	return fmt.Sprintf("%s/%s", path, file)
}

func Move(fromPath string, destination string) error {
	return os.Rename(fromPath, destination)
}

func ZipFiles(filename string, files []string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err = addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = filename
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
