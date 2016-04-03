package main

import (
	"os"
	"path/filepath"
	"fmt"
	"compress/gzip"
	"io"
)

func main() {
	gits := os.Getenv("GITS_HOME")
	home := os.Getenv("HOME")

	gzipit(gits+"/docker-client/README.md", home)
}

func gzipit(sourceDir, targetDir string) error {
	reader, err := os.Open(sourceDir)
	if err != nil {
		return err
	}

	filename := filepath.Base(sourceDir)
	fmt.Println(filename)
	targetDir = filepath.Join(targetDir, fmt.Sprintf("%s.gz", filename))

	writer, err := os.Create(targetDir)
	checkError(err)

	defer writer.Close()

	archiver := gzip.NewWriter(writer)
	archiver.Name = filename
	defer archiver.Close()

	fmt.Println("aaa")
	_, err = io.Copy(archiver, reader)
	checkError(err)

	return err
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		fmt.Println(err)
		os.Exit(1)
	}
}