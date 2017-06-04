package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/rpc/v2"
)

func TestNewHandler(t *testing.T) {
	r := &testRepo{}

	s := NewHandler(r).(*rpc.Server)

	t.Run("Has all methods", func(t *testing.T) {
		methods := map[string][]string{
			"SetList": {
				"GetAll", "Get", "Create", "Update", "Delete",
			},
		}

		for obj, ms := range methods {
			for _, method := range ms {
				m := fmt.Sprintf("%s.%s", obj, method)

				if !s.HasMethod(m) {
					t.Errorf("Expected to find method %q, but not found.", m)
				}
			}
		}
	})
}

func TestNewRepoLockingHandler(t *testing.T) {
	repo := &testRepo{false}

	handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if !repo.locked {
			t.Fatal("Expected repo to be locked, but isn't")
		}
	})

	lockingHandler := NewRepoLockingHandler(repo, handler)

	rw := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	lockingHandler.ServeHTTP(rw, r)

	if repo.locked {
		t.Fatal("Expected repo to be unlocked, but isn't")
	}
}
