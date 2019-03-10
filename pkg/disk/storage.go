package disk

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/StageAutoControl/controller/pkg/internal/stringslice"

	"github.com/peterbourgon/diskv"
)

// Storage is a filesystem abstraction storage for controller data
type Storage struct {
	disk *diskv.Diskv
}

func transform(s string) []string {
	parts := strings.Split(s, "_")
	if len(parts) == 1 {
		return parts
	}

	return parts[:len(parts)-1]
}

// New returns a new Storage instance with the given storage fileName set
func New(storageDir string) *Storage {
	return &Storage{
		disk: diskv.New(diskv.Options{
			Transform: transform,
			BasePath:  storageDir,
			TempDir:   os.TempDir(),
		}),
	}
}

func (s *Storage) buildFileName(key string, value interface{}) string {
	return fmt.Sprintf("%s_%s.json", s.getType(value), key)
}

// Has returns weather the storage has the given entity or not
func (s *Storage) Has(key string, kind interface{}) bool {
	return stringslice.Contains(key, s.List(kind))
}

// Write a given value with the given fileName to disk
func (s *Storage) Write(key string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value of type %s: %v", s.getType(value), err)
	}

	fileName := s.buildFileName(key, value)
	if err := s.disk.Write(fileName, b); err != nil {
		return fmt.Errorf("failed to write value of type %s to disk: %v", s.getType(value), err)
	}

	return nil
}

// Read a given id into the given value
func (s *Storage) Read(key string, value interface{}) error {
	fileName := s.buildFileName(key, value)

	b, err := s.disk.Read(fileName)
	if err != nil {
		return fmt.Errorf("failed to read value of type %s from disk: %v", s.getType(value), err)
	}

	if err := json.Unmarshal(b, value); err != nil {
		return fmt.Errorf("failed to unmarshal value of type %s: %v", s.getType(value), err)
	}

	return nil
}

// List the keys of a given kind
func (s *Storage) List(kind interface{}) []string {
	var keys []string
	kindType := s.getType(kind)

	for key := range s.disk.KeysPrefix(kindType, nil) {
		// Remove the custom file schema from the name, which should only return the pure key
		key = strings.TrimSuffix(key, ".json")
		key = strings.TrimPrefix(key, fmt.Sprintf("%s_", kindType))
		keys = append(keys, key)
	}

	return keys
}

// Delete a given key of a given kind
func (s *Storage) Delete(key string, kind interface{}) error {
	fileName := s.buildFileName(key, kind)
	if err := s.disk.Erase(fileName); err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}

func (s *Storage) getType(kind interface{}) string {
	t := reflect.TypeOf(kind)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}

	return t.Name()
}
