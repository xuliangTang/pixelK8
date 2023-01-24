package rbac

import (
	"io"
	"os"
)

func CertData(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	return b, err
}
