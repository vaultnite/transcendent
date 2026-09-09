package storage

import (
    "os"
    "path/filepath"
)

type ProviderLocal struct {
    basePath string
}

func NewStorageProviderLocal(basePath string) *ProviderLocal {
    return &ProviderLocal{
        basePath: basePath,
    }
}

func (sp *ProviderLocal) ensureDir(path string) (string, error) {
    fullPath := filepath.Join(sp.basePath, path)
    if err := os.MkdirAll(fullPath, os.ModePerm); err != nil {
        return "", err
    }
    return fullPath, nil
}

func (sp *ProviderLocal) Open(path string) (*os.File, error) {
    fullPath := filepath.Join(sp.basePath, path)
    return os.Open(fullPath)
}

func (sp *ProviderLocal) ReadDir(path string) ([]os.DirEntry, error) {
    if _, err := sp.ensureDir(path); err != nil {
        return nil, err
    }
    fullPath := filepath.Join(sp.basePath, path)
    return os.ReadDir(fullPath)
}

func (sp *ProviderLocal) ReadFile(path string) ([]byte, error) {
    fullPath := filepath.Join(sp.basePath, path)
    return os.ReadFile(fullPath)
}

func (sp *ProviderLocal) WriteFile(path string, data []byte) error {
    if _, err := sp.ensureDir(filepath.Dir(path)); err != nil {
        return err
    }
    fullPath := filepath.Join(sp.basePath, path)
    return os.WriteFile(fullPath, data, 0644)
}
