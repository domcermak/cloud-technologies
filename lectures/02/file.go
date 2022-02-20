package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// file1()
	// file2()
}

func file1() {
	f, err := os.Open("README.md")
	if err != nil {
		exit(fmt.Errorf("read README.MD: %v", err))
	}
	defer f.Close()

	buff := make([]byte, 4096)

	for {
		n, err := f.Read(buff)
		if n > 0 {
			//chunk := buff[:n]
			fmt.Println(fmt.Sprintf("read %v bytes", n))
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				exit(fmt.Errorf("read: %w", err))
			}
		}
	}

	if err := f.Close(); err != nil {
		exit(fmt.Errorf("close: %w", err))
	}
}

func file2() {
	f, err := os.CreateTemp(os.TempDir(), "ctc")
	if err != nil {
		exit(fmt.Errorf("read README.MD: %v", err))
	}
	defer f.Close()

	if _, err := f.Write([]byte("hello world")); err != nil {
		exit(fmt.Errorf("write: %w", err))
	}

	if err := f.Close(); err != nil {
		exit(fmt.Errorf("close: %w", err))
	}
}

func file3() {
	if err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fmt.Println("path", path, "dir", d.IsDir())
		return nil
	}); err != nil {
		exit(err)
	}
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
