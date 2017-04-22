package api

import "fmt"

func makeRepoReadError(err error) error {
	return fmt.Errorf("Unable to read repository: %v", err)
}
