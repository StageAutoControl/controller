package api

import (
	"strings"

	"github.com/satori/go.uuid"
)

func generateUUID() string {
	return uuid.NewV4().String()
}

func namesEqual(name1, name2 string) bool {
	name1 = strings.ToLower(strings.TrimSpace(name1))
	name2 = strings.ToLower(strings.TrimSpace(name2))

	return name1 == name2
}
