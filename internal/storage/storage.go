package storage

import "os"

// abstracted inferface for safely reading files; primarily hotfixes, from either local disk or remote (s3)

type Provider interface {
	Open(path string) (*os.File, error)
	ReadDir(path string) ([]os.DirEntry, error)
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte) error
}
