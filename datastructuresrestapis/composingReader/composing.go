package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(SHA1Sig("http.log.gz"))
}

// SHA1Sig() returns SHA1 signature of uncompressed file
func SHA1Sig(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("error opening file: %v", err)
		return "", err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	// gunxip

	r, err := gzip.NewReader(file)
	if err != nil {
		return "", fmt.Errorf("%q - gzip: %w", fileName, err)
	}

	w := sha1.New() // returns a hash
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - gzip: %w", fileName, err)
	}


	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
