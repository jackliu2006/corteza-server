package store

import (
	"fmt"
	"io"
	"path"

	"github.com/spf13/afero"
	"github.com/pkg/errors"
)

func New(namespace string) (Store, error) {
	return &store{
		namespace: namespace,
		originalFn: func(id uint64, ext string) string {
			return fmt.Sprintf("%d.%s", id, ext)
		},
		previewFn: func(id uint64, ext string) string {
			return fmt.Sprintf("%d_preview.%s", id, ext)
		},
	}, nil
}

func (s *store) Namespace() string {
	return s.namespace
}

func (s *store) check(filename string) error {
	if filename[:len(s.namespace)+1] != s.namespace + "/" {
		return errors.Errorf("Invalid namespace when trying to store file: %s (for %s)", filename, s.namespace)
	}
	return nil
}

func (s *store) Original(id uint64, ext string) string {
	return path.Join(s.namespace, s.originalFn(id, ext))
}

func (s *store) Preview(id uint64, ext string) string {
	return path.Join(s.namespace, s.previewFn(id, ext))
}

func (s *store) Save(filename string, contents io.Reader) error {
	// check filename for validity
	if err := s.check(filename); err != nil {
		return err
	}

	folder := path.Dir(filename)

	fs := afero.NewOsFs()
	fs.MkdirAll(folder, 0755)

	return afero.WriteReader(fs, filename, contents)
}

func (s *store) Remove(filename string) error {
	// check filename for validity
	if err := s.check(filename); err != nil {
		return err
	}

	fs := afero.NewOsFs()
	return fs.Remove(filename)
}

func (s *store) Open(filename string) (io.Reader, error) {
	// check filename for validity
	if err := s.check(filename); err != nil {
		return nil, err
	}

	fs := afero.NewOsFs()
	return fs.Open(filename)
}